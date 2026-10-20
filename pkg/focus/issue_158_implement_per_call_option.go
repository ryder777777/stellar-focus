package focus

import (
	"errors"
	"strings"
)

// ErrImplementPerCallOption158EmptyTarget is returned when the target key is blank.
var ErrImplementPerCallOption158EmptyTarget = errors.New("focus: target key cannot be empty")

// ImplementPerCallOption158Report summarizes deterministic diagnostics for issue #158.
type ImplementPerCallOption158Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateImplementPerCallOption158 computes a deterministic ImplementPerCallOption158Report for the supplied inputs.
func EvaluateImplementPerCallOption158(targetKey string, countBase, deltaUsed int) (ImplementPerCallOption158Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return ImplementPerCallOption158Report{}, ErrImplementPerCallOption158EmptyTarget
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
	return ImplementPerCallOption158Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
