package utils

type HasCreatedUpdatedBy interface {
	GetCreatedBy() uint
	GetUpdatedBy() uint
	GetDeletedBy() uint
}

func ExtractAuditUserID[T HasCreatedUpdatedBy](item T) []uint {
	return ExtractAuditUserIDs([]T{item})
}

func ExtractAuditUserIDs[T HasCreatedUpdatedBy](items []T) []uint {
	uniqueUserIDs := make(map[uint]struct{})

	for _, item := range items {
		if item.GetCreatedBy() > 0 {
			uniqueUserIDs[item.GetCreatedBy()] = struct{}{}
		}
		if item.GetUpdatedBy() > 0 {
			uniqueUserIDs[item.GetUpdatedBy()] = struct{}{}
		}
		if item.GetDeletedBy() > 0 {
			uniqueUserIDs[item.GetUpdatedBy()] = struct{}{}
		}
	}

	result := make([]uint, 0, len(uniqueUserIDs))
	for id := range uniqueUserIDs {
		result = append(result, id)
	}

	return result
}
