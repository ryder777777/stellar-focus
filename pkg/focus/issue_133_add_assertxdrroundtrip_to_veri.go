package focus

import (
	"errors"
	"strings"
)

// ErrAddAssertxdrroundtripToVerify133EmptyTarget is returned when the target key is blank.
var ErrAddAssertxdrroundtripToVerify133EmptyTarget = errors.New("focus: target key cannot be empty")

// AddAssertxdrroundtripToVerify133Report summarizes deterministic diagnostics for issue #133.
type AddAssertxdrroundtripToVerify133Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateAddAssertxdrroundtripToVerify133 computes a deterministic AddAssertxdrroundtripToVerify133Report for the supplied inputs.
func EvaluateAddAssertxdrroundtripToVerify133(targetKey string, countBase, deltaUsed int) (AddAssertxdrroundtripToVerify133Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return AddAssertxdrroundtripToVerify133Report{}, ErrAddAssertxdrroundtripToVerify133EmptyTarget
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
	return AddAssertxdrroundtripToVerify133Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
