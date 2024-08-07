package gocachemanager

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/singleflight"
	"google.golang.org/protobuf/proto"
)

var ErrInvalidSingleFlightType = errors.New(
	"invalid single flight type returned from refreshing the cache",
)

type CacheManager[TInput proto.Message, TOutput proto.Message] struct {
	prefix            string
	wrapper           *GoCacheWrapper
	factory           func() TOutput
	updateFn          func(context.Context, TInput, map[string]any) (TOutput, error)
	singleFlightGroup *singleflight.Group
}

func NewCacheManager[TInput proto.Message, TOutput proto.Message](
	prefix string,
	factory func() TOutput,
	updateFn func(context.Context, TInput, map[string]any) (TOutput, error),
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
		prefix:            prefix,
		wrapper:           wrapper,
		factory:           factory,
		updateFn:          updateFn,
		singleFlightGroup: new(singleflight.Group),
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
	dependencies ...map[string]any,
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

	output, err := cm.Refresh(ctx, input, dependencies...)
	if err != nil {
		return val, fmt.Errorf("updating data: %w", err)
	}

	return output, nil
}

func (cm *CacheManager[TInput, TOutput]) Refresh(
	ctx context.Context,
	input TInput,
	dependencies ...map[string]any,
) (TOutput, error) {
	var empty TOutput

	key, err := cm.getKey(input)
	if err != nil {
		return empty, fmt.Errorf("getting key: %w", err)
	}

	val, err, _ := cm.singleFlightGroup.Do(string(key), func() (interface{}, error) {
		deps := map[string]any{}
		if len(dependencies) > 0 {
			deps = dependencies[0]
		}
		val, err := cm.updateFn(ctx, input, deps)
		if err != nil {
			return nil, fmt.Errorf("updating data: %w", err)
		}

		return val, nil
	})
	if err != nil {
		return empty, err
	}

	converted, ok := val.(TOutput)
	if !ok {
		return empty, fmt.Errorf("converting data: %w", ErrInvalidSingleFlightType)
	}

	data, err := proto.MarshalOptions{
		Deterministic: true,
	}.Marshal(converted)
	if err != nil {
		return empty, fmt.Errorf("marshalling data: %w", err)
	}

	if err := cm.wrapper.Set(ctx, key, data); err != nil {
		return empty, fmt.Errorf("setting data: %w", err)
	}

	return converted, nil
}
