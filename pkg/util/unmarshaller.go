package util

import (
	"encoding/json"
	"fmt"
)

type Registry[T any] struct {
	constructors map[string]func() T
}

func NewRegistry[T any]() *Registry[T] {
	return &Registry[T]{
		constructors: make(map[string]func() T),
	}
}

func (r *Registry[T]) Register(typeName string, constructor func() T) {
	r.constructors[typeName] = constructor
}

func (r *Registry[T]) Create(typeName string) (T, bool) {
	constructor, ok := r.constructors[typeName]
	if !ok {
		var zero T
		return zero, false
	}
	return constructor(), true
}

func MapField(raw map[string]json.RawMessage, field string, target interface{}, errHandler *ErrorHandler) {
	if _, ok := raw[field]; !ok {
		return
	}
	errHandler.Push(field, json.Unmarshal(raw[field], target))
}

func UnmarshalWithRegistry[T any](data []byte, registry *Registry[T], typeField string) (T, error) {
	var zero T

	// First, unmarshal just to get the type field
	var typeDiscriminator map[string]interface{}
	if err := json.Unmarshal(data, &typeDiscriminator); err != nil {
		return zero, fmt.Errorf("failed to unmarshal for type discrimination: %w", err)
	}

	// Get the type value
	typeValue, ok := typeDiscriminator[typeField].(string)
	if !ok {
		return zero, fmt.Errorf("type field '%s' not found or not a string", typeField)
	}

	// Create the appropriate concrete type
	instance, ok := registry.Create(typeValue)
	if !ok {
		return zero, fmt.Errorf("unknown type: %s", typeValue)
	}

	// Unmarshal into the concrete type
	if err := json.Unmarshal(data, &instance); err != nil {
		return zero, fmt.Errorf("failed to unmarshal into concrete type: %w", err)
	}

	return instance, nil
}

// Step 5: Helper for unmarshaling slices
func UnmarshalSliceWithRegistry[T any](data []byte, registry *Registry[T], typeField string) ([]T, error) {
	// First unmarshal into []json.RawMessage
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(data, &rawMessages); err != nil {
		return nil, fmt.Errorf("failed to unmarshal slice: %w", err)
	}

	// Process each element
	result := make([]T, 0, len(rawMessages))
	for i, raw := range rawMessages {
		item, err := UnmarshalWithRegistry(raw, registry, typeField)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal element %d: %w", i, err)
		}
		result = append(result, item)
	}

	return result, nil
}
