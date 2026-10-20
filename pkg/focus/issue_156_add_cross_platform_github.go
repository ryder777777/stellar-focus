package focus

import (
	"errors"
	"strings"
)

// ErrAddCrossPlatformGithub156EmptyTarget is returned when the target key is blank.
var ErrAddCrossPlatformGithub156EmptyTarget = errors.New("focus: target key cannot be empty")

// AddCrossPlatformGithub156Report summarizes deterministic diagnostics for issue #156.
type AddCrossPlatformGithub156Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateAddCrossPlatformGithub156 computes a deterministic AddCrossPlatformGithub156Report for the supplied inputs.
func EvaluateAddCrossPlatformGithub156(targetKey string, countBase, deltaUsed int) (AddCrossPlatformGithub156Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return AddCrossPlatformGithub156Report{}, ErrAddCrossPlatformGithub156EmptyTarget
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
	return AddCrossPlatformGithub156Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
