package model

type (
	Checkin struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Frequency   int    `json:"frequency"`
		TargetCount int    `json:"target_count"`
	}

	CheckinReq struct {
		HabitID int `json:"habit_id" validate:"required"`
	}

	CheckinRes struct {
		ID        int     `json:"id"`
		HabitID   int     `json:"habit_id"`
		CreatedAt int64   `json:"created_at"`
		UpdatedAt int64   `json:"updated_at"`
		CreatedBy *string `json:"created_by"`
		UpdatedBy *string `json:"updated_by"`
	}

	CheckinSearchReq struct {
		Page   int              `json:"page" validate:"required"`
		Limit  int              `json:"limit" validate:"required"`
		Filter CheckinFilterReq `json:"filter" validate:"required"`
	}

	CheckinFilterReq struct {
		HabitID   *int `json:"habit_id"`
		Frequency *int `json:"frequency"`
	}

	CheckinSearchRes struct {
		Item     []*CheckinRes  `json:"item"`
		Paginate PaginateResult `json:"paginate"`
	}

	PaginateResult struct {
		Page      int64 `json:"page"`
		TotalPage int64 `json:"total_page"`
		Total     int64 `json:"total"`
	}
)
