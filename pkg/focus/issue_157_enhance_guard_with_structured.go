package focus

import (
	"errors"
	"strings"
)

// ErrEnhanceGuardWithStructured157EmptyTarget is returned when the target key is blank.
var ErrEnhanceGuardWithStructured157EmptyTarget = errors.New("focus: target key cannot be empty")

// EnhanceGuardWithStructured157Report summarizes deterministic diagnostics for issue #157.
type EnhanceGuardWithStructured157Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateEnhanceGuardWithStructured157 computes a deterministic EnhanceGuardWithStructured157Report for the supplied inputs.
func EvaluateEnhanceGuardWithStructured157(targetKey string, countBase, deltaUsed int) (EnhanceGuardWithStructured157Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return EnhanceGuardWithStructured157Report{}, ErrEnhanceGuardWithStructured157EmptyTarget
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
	return EnhanceGuardWithStructured157Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
