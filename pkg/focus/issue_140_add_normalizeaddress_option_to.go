package focus

import (
	"errors"
	"strings"
)

// ErrAddNormalizeaddressOptionTo140EmptyTarget is returned when the target key is blank.
var ErrAddNormalizeaddressOptionTo140EmptyTarget = errors.New("focus: target key cannot be empty")

// AddNormalizeaddressOptionTo140Report summarizes deterministic diagnostics for issue #140.
type AddNormalizeaddressOptionTo140Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateAddNormalizeaddressOptionTo140 computes a deterministic AddNormalizeaddressOptionTo140Report for the supplied inputs.
func EvaluateAddNormalizeaddressOptionTo140(targetKey string, countBase, deltaUsed int) (AddNormalizeaddressOptionTo140Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return AddNormalizeaddressOptionTo140Report{}, ErrAddNormalizeaddressOptionTo140EmptyTarget
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
	return AddNormalizeaddressOptionTo140Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
