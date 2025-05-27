package utils

type WithID interface {
	GetID() uint
}

func MapperByID[T WithID](items []T) map[uint]T {
	result := make(map[uint]T)
	for _, item := range items {
		result[item.GetID()] = item
	}
	return result
}
