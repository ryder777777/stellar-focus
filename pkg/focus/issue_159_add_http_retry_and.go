package focus

import (
	"errors"
	"strings"
)

// ErrAddHttpRetryAnd159EmptyTarget is returned when the target key is blank.
var ErrAddHttpRetryAnd159EmptyTarget = errors.New("focus: target key cannot be empty")

// AddHttpRetryAnd159Report summarizes deterministic diagnostics for issue #159.
type AddHttpRetryAnd159Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateAddHttpRetryAnd159 computes a deterministic AddHttpRetryAnd159Report for the supplied inputs.
func EvaluateAddHttpRetryAnd159(targetKey string, countBase, deltaUsed int) (AddHttpRetryAnd159Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return AddHttpRetryAnd159Report{}, ErrAddHttpRetryAnd159EmptyTarget
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
	return AddHttpRetryAnd159Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
