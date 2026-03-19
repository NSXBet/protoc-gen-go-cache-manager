package v2

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/singleflight"
	"google.golang.org/protobuf/proto"
)

// ErrInvalidSingleFlightType is returned when the singleflight result is not the expected type.
var ErrInvalidSingleFlightType = errors.New(
	"invalid single flight type returned from refreshing the cache",
)

// ErrTooManyDependencyMaps is returned when more than one dependency map is passed to Get, Refresh, Replace, or Delete.
var ErrTooManyDependencyMaps = errors.New(
	"at most one dependency map is allowed",
)

// CacheManagerV2 is a cache manager that stores protobuf values as raw []byte (no base64, no gzip).
type CacheManagerV2[TInput proto.Message, TOutput proto.Message] struct {
	prefix            string
	wrapper           *GoCacheWrapperV2
	factory           func() TOutput
	updateFn          func(context.Context, TInput, map[string]any) (TOutput, error)
	singleFlightGroup *singleflight.Group
}

// NewCacheManager creates a v2 cache manager with the same API as v1 CacheManager.
func NewCacheManager[TInput proto.Message, TOutput proto.Message](
	prefix string,
	factory func() TOutput,
	updateFn func(context.Context, TInput, map[string]any) (TOutput, error),
	options ...CacheOption,
) (*CacheManagerV2[TInput, TOutput], error) {
	if factory == nil {
		return nil, errors.New("factory is required")
	}

	if updateFn == nil {
		return nil, errors.New("updateFn is required")
	}

	settings := DefaultCacheSettings()
	for _, option := range options {
		option(settings)
	}

	wrapper, err := NewGoCacheWrapper(prefix, 10*time.Second, settings)
	if err != nil {
		return nil, fmt.Errorf("creating cache manager: %w", err)
	}

	return &CacheManagerV2[TInput, TOutput]{
		prefix:            prefix,
		wrapper:           wrapper,
		factory:           factory,
		updateFn:          updateFn,
		singleFlightGroup: new(singleflight.Group),
	}, nil
}

func (cm *CacheManagerV2[TInput, TOutput]) getKey(input TInput) ([]byte, error) {
	data, err := proto.MarshalOptions{
		Deterministic: true,
	}.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("marshalling input: %w", err)
	}
	return data, nil
}

// Get returns the cached value or refreshes and caches it on miss.
func (cm *CacheManagerV2[TInput, TOutput]) Get(
	ctx context.Context,
	input TInput,
	dependencies ...map[string]any,
) (TOutput, error) {
	var val TOutput

	if len(dependencies) > 1 {
		return val, ErrTooManyDependencyMaps
	}

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

// Refresh loads fresh data via updateFn and stores it in the cache.
func (cm *CacheManagerV2[TInput, TOutput]) Refresh(
	ctx context.Context,
	input TInput,
	dependencies ...map[string]any,
) (TOutput, error) {
	var empty TOutput

	if len(dependencies) > 1 {
		return empty, ErrTooManyDependencyMaps
	}

	key, err := cm.getKey(input)
	if err != nil {
		return empty, fmt.Errorf("getting key: %w", err)
	}

	val, err, _ := cm.singleFlightGroup.Do(string(key), func() (interface{}, error) {
		deps := map[string]any{}
		if len(dependencies) > 0 && dependencies[0] != nil {
			deps = dependencies[0]
		}
		return cm.updateFn(ctx, input, deps)
	})
	if err != nil {
		return empty, err
	}

	converted, ok := val.(TOutput)
	if !ok {
		return empty, fmt.Errorf("converting data: %w", ErrInvalidSingleFlightType)
	}

	return cm.Replace(ctx, input, converted, dependencies...)
}

// Replace writes the given value to the cache.
func (cm *CacheManagerV2[TInput, TOutput]) Replace(
	ctx context.Context,
	input TInput,
	newValue TOutput,
	dependencies ...map[string]any,
) (TOutput, error) {
	var empty TOutput

	if len(dependencies) > 1 {
		return empty, ErrTooManyDependencyMaps
	}

	key, err := cm.getKey(input)
	if err != nil {
		return empty, fmt.Errorf("getting key: %w", err)
	}

	data, err := proto.MarshalOptions{
		Deterministic: true,
	}.Marshal(newValue)
	if err != nil {
		return empty, fmt.Errorf("marshalling data: %w", err)
	}

	if err := cm.wrapper.Set(ctx, key, data); err != nil {
		return empty, fmt.Errorf("setting data: %w", err)
	}
	return newValue, nil
}

// Delete removes the entry for the given input from the cache.
func (cm *CacheManagerV2[TInput, TOutput]) Delete(
	ctx context.Context,
	input TInput,
	dependencies ...map[string]any,
) error {
	if len(dependencies) > 1 {
		return ErrTooManyDependencyMaps
	}

	key, err := cm.getKey(input)
	if err != nil {
		return fmt.Errorf("getting key: %w", err)
	}
	if err := cm.wrapper.Delete(ctx, key); err != nil {
		return fmt.Errorf("deleting data: %w", err)
	}
	return nil
}
