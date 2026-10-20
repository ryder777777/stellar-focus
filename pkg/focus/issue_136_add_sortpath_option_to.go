package focus

import (
	"errors"
	"strings"
)

// ErrAddSortpathOptionTo136EmptyTarget is returned when the target key is blank.
var ErrAddSortpathOptionTo136EmptyTarget = errors.New("focus: target key cannot be empty")

// AddSortpathOptionTo136Report summarizes deterministic diagnostics for issue #136.
type AddSortpathOptionTo136Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateAddSortpathOptionTo136 computes a deterministic AddSortpathOptionTo136Report for the supplied inputs.
func EvaluateAddSortpathOptionTo136(targetKey string, countBase, deltaUsed int) (AddSortpathOptionTo136Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return AddSortpathOptionTo136Report{}, ErrAddSortpathOptionTo136EmptyTarget
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
	return AddSortpathOptionTo136Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
