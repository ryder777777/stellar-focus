package focus

import (
	"errors"
	"strings"
)

// ErrAddPkgGoldenRequire146EmptyTarget is returned when the target key is blank.
var ErrAddPkgGoldenRequire146EmptyTarget = errors.New("focus: target key cannot be empty")

// AddPkgGoldenRequire146Report summarizes deterministic diagnostics for issue #146.
type AddPkgGoldenRequire146Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateAddPkgGoldenRequire146 computes a deterministic AddPkgGoldenRequire146Report for the supplied inputs.
func EvaluateAddPkgGoldenRequire146(targetKey string, countBase, deltaUsed int) (AddPkgGoldenRequire146Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return AddPkgGoldenRequire146Report{}, ErrAddPkgGoldenRequire146EmptyTarget
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
	return AddPkgGoldenRequire146Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
