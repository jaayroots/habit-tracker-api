package model

type (
	Habit struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Frequency   int    `json:"frequency"`
		TargetCount int    `json:"target_count"`
	}

	HabitReq struct {
		Title       string `json:"title" validate:"required"`
		Description string `json:"description" validate:"required"`
		Frequency   int    `json:"frequency" validate:"required"`
		TargetCount int    `json:"target_count" validate:"required"`
	}

	HabitRes struct {
		ID          int     `json:"id"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Frequency   int     `json:"frequency"`
		TargetCount int     `json:"target_count"`
		Checkin     int     `json:"checkin"`
		CreatedAt   int64   `json:"created_at"`
		UpdatedAt   int64   `json:"updated_at"`
		CreatedBy   *string `json:"created_by"`
		UpdatedBy   *string `json:"updated_by"`
	}

	HabitSearchReq struct {
		Page   int            `json:"page" validate:"required"`
		Limit  int            `json:"limit" validate:"required"`
		Filter HabitFilterReq `json:"filter" validate:"required"`
	}

	HabitFilterReq struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
		Frequency   *int    `json:"frequency"`
	}

	HabitSearchRes struct {
		Item     []*HabitRes    `json:"item"`
		Paginate PaginateResult `json:"paginate"`
	}

	PaginateResult struct {
		Page      int64 `json:"page"`
		TotalPage int64 `json:"total_page"`
		Total     int64 `json:"total"`
	}
)
