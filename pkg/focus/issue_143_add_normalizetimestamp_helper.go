package focus

import (
	"errors"
	"strings"
)

// ErrAddNormalizetimestampHelperFor143EmptyTarget is returned when the target key is blank.
var ErrAddNormalizetimestampHelperFor143EmptyTarget = errors.New("focus: target key cannot be empty")

// AddNormalizetimestampHelperFor143Report summarizes deterministic diagnostics for issue #143.
type AddNormalizetimestampHelperFor143Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateAddNormalizetimestampHelperFor143 computes a deterministic AddNormalizetimestampHelperFor143Report for the supplied inputs.
func EvaluateAddNormalizetimestampHelperFor143(targetKey string, countBase, deltaUsed int) (AddNormalizetimestampHelperFor143Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return AddNormalizetimestampHelperFor143Report{}, ErrAddNormalizetimestampHelperFor143EmptyTarget
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
	return AddNormalizetimestampHelperFor143Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
