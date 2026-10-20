package focus

import (
	"errors"
	"strings"
)

// ErrAddIgnoreledgerheaderOptionFor142EmptyTarget is returned when the target key is blank.
var ErrAddIgnoreledgerheaderOptionFor142EmptyTarget = errors.New("focus: target key cannot be empty")

// AddIgnoreledgerheaderOptionFor142Report summarizes deterministic diagnostics for issue #142.
type AddIgnoreledgerheaderOptionFor142Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateAddIgnoreledgerheaderOptionFor142 computes a deterministic AddIgnoreledgerheaderOptionFor142Report for the supplied inputs.
func EvaluateAddIgnoreledgerheaderOptionFor142(targetKey string, countBase, deltaUsed int) (AddIgnoreledgerheaderOptionFor142Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return AddIgnoreledgerheaderOptionFor142Report{}, ErrAddIgnoreledgerheaderOptionFor142EmptyTarget
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
	return AddIgnoreledgerheaderOptionFor142Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
