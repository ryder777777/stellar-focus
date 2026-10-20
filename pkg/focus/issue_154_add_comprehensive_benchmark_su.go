package focus

import (
	"errors"
	"strings"
)

// ErrAddComprehensiveBenchmarkSuite154EmptyTarget is returned when the target key is blank.
var ErrAddComprehensiveBenchmarkSuite154EmptyTarget = errors.New("focus: target key cannot be empty")

// AddComprehensiveBenchmarkSuite154Report summarizes deterministic diagnostics for issue #154.
type AddComprehensiveBenchmarkSuite154Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateAddComprehensiveBenchmarkSuite154 computes a deterministic AddComprehensiveBenchmarkSuite154Report for the supplied inputs.
func EvaluateAddComprehensiveBenchmarkSuite154(targetKey string, countBase, deltaUsed int) (AddComprehensiveBenchmarkSuite154Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return AddComprehensiveBenchmarkSuite154Report{}, ErrAddComprehensiveBenchmarkSuite154EmptyTarget
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
	return AddComprehensiveBenchmarkSuite154Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
