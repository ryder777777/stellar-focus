package focus

import (
	"errors"
	"strings"
)

// ErrAddRapidPropertyBased153EmptyTarget is returned when the target key is blank.
var ErrAddRapidPropertyBased153EmptyTarget = errors.New("focus: target key cannot be empty")

// AddRapidPropertyBased153Report summarizes deterministic diagnostics for issue #153.
type AddRapidPropertyBased153Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateAddRapidPropertyBased153 computes a deterministic AddRapidPropertyBased153Report for the supplied inputs.
func EvaluateAddRapidPropertyBased153(targetKey string, countBase, deltaUsed int) (AddRapidPropertyBased153Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return AddRapidPropertyBased153Report{}, ErrAddRapidPropertyBased153EmptyTarget
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
	return AddRapidPropertyBased153Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
