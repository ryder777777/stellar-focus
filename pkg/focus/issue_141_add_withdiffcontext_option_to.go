package focus

import (
	"errors"
	"strings"
)

// ErrAddWithdiffcontextOptionTo141EmptyTarget is returned when the target key is blank.
var ErrAddWithdiffcontextOptionTo141EmptyTarget = errors.New("focus: target key cannot be empty")

// AddWithdiffcontextOptionTo141Report summarizes deterministic diagnostics for issue #141.
type AddWithdiffcontextOptionTo141Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateAddWithdiffcontextOptionTo141 computes a deterministic AddWithdiffcontextOptionTo141Report for the supplied inputs.
func EvaluateAddWithdiffcontextOptionTo141(targetKey string, countBase, deltaUsed int) (AddWithdiffcontextOptionTo141Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return AddWithdiffcontextOptionTo141Report{}, ErrAddWithdiffcontextOptionTo141EmptyTarget
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
	return AddWithdiffcontextOptionTo141Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
