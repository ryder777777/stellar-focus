package focus

import (
	"errors"
	"strings"
)

// ErrExtendFocusTbInterface147EmptyTarget is returned when the target key is blank.
var ErrExtendFocusTbInterface147EmptyTarget = errors.New("focus: target key cannot be empty")

// ExtendFocusTbInterface147Report summarizes deterministic diagnostics for issue #147.
type ExtendFocusTbInterface147Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateExtendFocusTbInterface147 computes a deterministic ExtendFocusTbInterface147Report for the supplied inputs.
func EvaluateExtendFocusTbInterface147(targetKey string, countBase, deltaUsed int) (ExtendFocusTbInterface147Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return ExtendFocusTbInterface147Report{}, ErrExtendFocusTbInterface147EmptyTarget
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
	return ExtendFocusTbInterface147Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
