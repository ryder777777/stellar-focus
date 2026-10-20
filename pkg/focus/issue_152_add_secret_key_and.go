package focus

import (
	"errors"
	"strings"
)

// ErrAddSecretKeyAnd152EmptyTarget is returned when the target key is blank.
var ErrAddSecretKeyAnd152EmptyTarget = errors.New("focus: target key cannot be empty")

// AddSecretKeyAnd152Report summarizes deterministic diagnostics for issue #152.
type AddSecretKeyAnd152Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateAddSecretKeyAnd152 computes a deterministic AddSecretKeyAnd152Report for the supplied inputs.
func EvaluateAddSecretKeyAnd152(targetKey string, countBase, deltaUsed int) (AddSecretKeyAnd152Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return AddSecretKeyAnd152Report{}, ErrAddSecretKeyAnd152EmptyTarget
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
	return AddSecretKeyAnd152Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
