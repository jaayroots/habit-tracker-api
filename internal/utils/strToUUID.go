package utils

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func StrToUUID(s string) (uuid.UUID, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return uuid.Nil, fmt.Errorf("uuid is empty")
	}
	id, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid uuid %q: %w", s, err)
	}
	return id, nil
}
