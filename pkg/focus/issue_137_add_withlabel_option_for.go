package focus

import (
	"errors"
	"strings"
)

// ErrAddWithlabelOptionFor137EmptyTarget is returned when the target key is blank.
var ErrAddWithlabelOptionFor137EmptyTarget = errors.New("focus: target key cannot be empty")

// AddWithlabelOptionFor137Report summarizes deterministic diagnostics for issue #137.
type AddWithlabelOptionFor137Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateAddWithlabelOptionFor137 computes a deterministic AddWithlabelOptionFor137Report for the supplied inputs.
func EvaluateAddWithlabelOptionFor137(targetKey string, countBase, deltaUsed int) (AddWithlabelOptionFor137Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return AddWithlabelOptionFor137Report{}, ErrAddWithlabelOptionFor137EmptyTarget
	}
	if countBase < 0 {
		countBase = 0
	}
	if deltaUsed < 0 {
		deltaUsed = 0
	}
	statusTag := "nominal"
	if countBase+deltaUsed > 100 {
		statusTag = "elevated"
	}
	return AddWithlabelOptionFor137Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
