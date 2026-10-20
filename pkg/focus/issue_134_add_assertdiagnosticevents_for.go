package focus

import (
	"errors"
	"strings"
)

// ErrAddAssertdiagnosticeventsForSoroban134EmptyTarget is returned when the target key is blank.
var ErrAddAssertdiagnosticeventsForSoroban134EmptyTarget = errors.New("focus: target key cannot be empty")

// AddAssertdiagnosticeventsForSoroban134Report summarizes deterministic diagnostics for issue #134.
type AddAssertdiagnosticeventsForSoroban134Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateAddAssertdiagnosticeventsForSoroban134 computes a deterministic AddAssertdiagnosticeventsForSoroban134Report for the supplied inputs.
func EvaluateAddAssertdiagnosticeventsForSoroban134(targetKey string, countBase, deltaUsed int) (AddAssertdiagnosticeventsForSoroban134Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return AddAssertdiagnosticeventsForSoroban134Report{}, ErrAddAssertdiagnosticeventsForSoroban134EmptyTarget
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
	return AddAssertdiagnosticeventsForSoroban134Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
