package focus

import (
	"errors"
	"strings"
)

// ErrAchieveCompleteAssertionParity144EmptyTarget is returned when the target key is blank.
var ErrAchieveCompleteAssertionParity144EmptyTarget = errors.New("focus: target key cannot be empty")

// AchieveCompleteAssertionParity144Report summarizes deterministic diagnostics for issue #144.
type AchieveCompleteAssertionParity144Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateAchieveCompleteAssertionParity144 computes a deterministic AchieveCompleteAssertionParity144Report for the supplied inputs.
func EvaluateAchieveCompleteAssertionParity144(targetKey string, countBase, deltaUsed int) (AchieveCompleteAssertionParity144Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return AchieveCompleteAssertionParity144Report{}, ErrAchieveCompleteAssertionParity144EmptyTarget
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
	return AchieveCompleteAssertionParity144Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
