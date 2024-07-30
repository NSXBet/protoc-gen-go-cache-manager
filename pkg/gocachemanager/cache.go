package gocachemanager

import (
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/protobuf/proto"
)

type CacheManager[TInput proto.Message, TOutput proto.Message] struct {
	prefix   string
	wrapper  *GoCacheWrapper
	factory  func() TOutput
	updateFn func(context.Context, TInput) (TOutput, error)
}

func NewCacheManager[TInput proto.Message, TOutput proto.Message](
	prefix string,
	factory func() TOutput,
	updateFn func(context.Context, TInput) (TOutput, error),
	options ...CacheOption,
) (*CacheManager[TInput, TOutput], error) {
	settings := DefaultCacheSettings()
	for _, option := range options {
		option(settings)
	}

	wrapper, err := NewGoCacheWrapper(prefix, 10*time.Second, settings)
	if err != nil {
		return nil, fmt.Errorf("creating cache manager: %w", err)
	}

	return &CacheManager[TInput, TOutput]{
		prefix:   prefix,
		wrapper:  wrapper,
		factory:  factory,
		updateFn: updateFn,
	}, nil
}

func (cm *CacheManager[TInput, TOutput]) getKey(
	input TInput,
) ([]byte, error) {
	data, err := proto.MarshalOptions{
		Deterministic: true,
	}.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("marshalling input: %w", err)
	}

	return data, nil
}

func (cm *CacheManager[TInput, TOutput]) Get(
	ctx context.Context,
	input TInput,
) (TOutput, error) {
	var val TOutput

	key, err := cm.getKey(input)
	if err != nil {
		return val, fmt.Errorf("getting key: %w", err)
	}

	data, err := cm.wrapper.Get(ctx, key)
	if err != nil {
		if !errors.Is(err, ErrCacheMiss) {
			return val, fmt.Errorf("getting data: %w", err)
		}
	}

	if data != nil {
		result := cm.factory()
		if err := proto.Unmarshal(data, result); err != nil {
			return val, fmt.Errorf("unmarshalling data: %w", err)
		}

		return result, nil
	}

	output, err := cm.Refresh(ctx, input)
	if err != nil {
		return val, fmt.Errorf("updating data: %w", err)
	}

	return output, nil
}

func (cm *CacheManager[TInput, TOutput]) Refresh(
	ctx context.Context,
	input TInput,
) (TOutput, error) {
	// TODO: singleflight
	val, err := cm.updateFn(ctx, input)
	if err != nil {
		return val, fmt.Errorf("updating data: %w", err)
	}

	data, err := proto.MarshalOptions{
		Deterministic: true,
	}.Marshal(val)
	if err != nil {
		return val, fmt.Errorf("marshalling data: %w", err)
	}

	key, err := cm.getKey(input)
	if err != nil {
		return val, fmt.Errorf("getting key: %w", err)
	}

	if err := cm.wrapper.Set(ctx, key, data); err != nil {
		return val, fmt.Errorf("setting data: %w", err)
	}

	return val, nil
}
