package focus

import (
	"errors"
	"strings"
)

// ErrAddNetworkPassphraseTracking150EmptyTarget is returned when the target key is blank.
var ErrAddNetworkPassphraseTracking150EmptyTarget = errors.New("focus: target key cannot be empty")

// AddNetworkPassphraseTracking150Report summarizes deterministic diagnostics for issue #150.
type AddNetworkPassphraseTracking150Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateAddNetworkPassphraseTracking150 computes a deterministic AddNetworkPassphraseTracking150Report for the supplied inputs.
func EvaluateAddNetworkPassphraseTracking150(targetKey string, countBase, deltaUsed int) (AddNetworkPassphraseTracking150Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return AddNetworkPassphraseTracking150Report{}, ErrAddNetworkPassphraseTracking150EmptyTarget
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
	return AddNetworkPassphraseTracking150Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
