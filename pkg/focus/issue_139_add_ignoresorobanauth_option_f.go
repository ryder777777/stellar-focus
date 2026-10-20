package focus

import (
	"errors"
	"strings"
)

// ErrAddIgnoresorobanauthOptionFor139EmptyTarget is returned when the target key is blank.
var ErrAddIgnoresorobanauthOptionFor139EmptyTarget = errors.New("focus: target key cannot be empty")

// AddIgnoresorobanauthOptionFor139Report summarizes deterministic diagnostics for issue #139.
type AddIgnoresorobanauthOptionFor139Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateAddIgnoresorobanauthOptionFor139 computes a deterministic AddIgnoresorobanauthOptionFor139Report for the supplied inputs.
func EvaluateAddIgnoresorobanauthOptionFor139(targetKey string, countBase, deltaUsed int) (AddIgnoresorobanauthOptionFor139Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return AddIgnoresorobanauthOptionFor139Report{}, ErrAddIgnoresorobanauthOptionFor139EmptyTarget
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
	return AddIgnoresorobanauthOptionFor139Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
