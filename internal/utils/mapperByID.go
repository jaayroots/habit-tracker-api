package utils

import "github.com/google/uuid"

type WithID interface {
	GetID() uuid.UUID
}

func MapperByID[T WithID](items []T) map[uuid.UUID]T {
	result := make(map[uuid.UUID]T)
	for _, item := range items {
		result[item.GetID()] = item
	}
	return result
}
