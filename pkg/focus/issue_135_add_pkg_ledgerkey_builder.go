package focus

import (
	"errors"
	"strings"
)

// ErrAddPkgLedgerkeyBuilder135EmptyTarget is returned when the target key is blank.
var ErrAddPkgLedgerkeyBuilder135EmptyTarget = errors.New("focus: target key cannot be empty")

// AddPkgLedgerkeyBuilder135Report summarizes deterministic diagnostics for issue #135.
type AddPkgLedgerkeyBuilder135Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateAddPkgLedgerkeyBuilder135 computes a deterministic AddPkgLedgerkeyBuilder135Report for the supplied inputs.
func EvaluateAddPkgLedgerkeyBuilder135(targetKey string, countBase, deltaUsed int) (AddPkgLedgerkeyBuilder135Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return AddPkgLedgerkeyBuilder135Report{}, ErrAddPkgLedgerkeyBuilder135EmptyTarget
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
	return AddPkgLedgerkeyBuilder135Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
