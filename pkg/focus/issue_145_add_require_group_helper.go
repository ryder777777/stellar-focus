package focus

import (
	"errors"
	"strings"
)

// ErrAddRequireGroupHelper145EmptyTarget is returned when the target key is blank.
var ErrAddRequireGroupHelper145EmptyTarget = errors.New("focus: target key cannot be empty")

// AddRequireGroupHelper145Report summarizes deterministic diagnostics for issue #145.
type AddRequireGroupHelper145Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateAddRequireGroupHelper145 computes a deterministic AddRequireGroupHelper145Report for the supplied inputs.
func EvaluateAddRequireGroupHelper145(targetKey string, countBase, deltaUsed int) (AddRequireGroupHelper145Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return AddRequireGroupHelper145Report{}, ErrAddRequireGroupHelper145EmptyTarget
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
	return AddRequireGroupHelper145Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
