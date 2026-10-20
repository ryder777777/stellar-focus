package focus

import (
	"errors"
	"strings"
)

// ErrValidateIgnorepathsPatternSyntax138EmptyTarget is returned when the target key is blank.
var ErrValidateIgnorepathsPatternSyntax138EmptyTarget = errors.New("focus: target key cannot be empty")

// ValidateIgnorepathsPatternSyntax138Report summarizes deterministic diagnostics for issue #138.
type ValidateIgnorepathsPatternSyntax138Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateValidateIgnorepathsPatternSyntax138 computes a deterministic ValidateIgnorepathsPatternSyntax138Report for the supplied inputs.
func EvaluateValidateIgnorepathsPatternSyntax138(targetKey string, countBase, deltaUsed int) (ValidateIgnorepathsPatternSyntax138Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return ValidateIgnorepathsPatternSyntax138Report{}, ErrValidateIgnorepathsPatternSyntax138EmptyTarget
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
	return ValidateIgnorepathsPatternSyntax138Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
