package utils

import "github.com/google/uuid"

type HasCreatedUpdatedBy interface {
	GetCreatedBy() uuid.UUID
	GetUpdatedBy() uuid.UUID
	GetDeletedBy() uuid.UUID
}

func ExtractAuditUserID[T HasCreatedUpdatedBy](item T) []uuid.UUID {
	return ExtractAuditUserIDs([]T{item})
}

func ExtractAuditUserIDs[T HasCreatedUpdatedBy](items []T) []uuid.UUID {
	uniqueUserIDs := make(map[uuid.UUID]struct{})

	for _, item := range items {
		if item.GetCreatedBy() != uuid.Nil {
			uniqueUserIDs[item.GetCreatedBy()] = struct{}{}
		}
		if item.GetUpdatedBy() != uuid.Nil {
			uniqueUserIDs[item.GetUpdatedBy()] = struct{}{}
		}
		if item.GetDeletedBy() != uuid.Nil {
			uniqueUserIDs[item.GetUpdatedBy()] = struct{}{}
		}
	}

	result := make([]uuid.UUID, 0, len(uniqueUserIDs))
	for id := range uniqueUserIDs {
		result = append(result, id)
	}

	return result
}
