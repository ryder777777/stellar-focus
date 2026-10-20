package focus

import (
	"errors"
	"strings"
)

// ErrConfigureThelperAndGocritic155EmptyTarget is returned when the target key is blank.
var ErrConfigureThelperAndGocritic155EmptyTarget = errors.New("focus: target key cannot be empty")

// ConfigureThelperAndGocritic155Report summarizes deterministic diagnostics for issue #155.
type ConfigureThelperAndGocritic155Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateConfigureThelperAndGocritic155 computes a deterministic ConfigureThelperAndGocritic155Report for the supplied inputs.
func EvaluateConfigureThelperAndGocritic155(targetKey string, countBase, deltaUsed int) (ConfigureThelperAndGocritic155Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return ConfigureThelperAndGocritic155Report{}, ErrConfigureThelperAndGocritic155EmptyTarget
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
	return ConfigureThelperAndGocritic155Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
