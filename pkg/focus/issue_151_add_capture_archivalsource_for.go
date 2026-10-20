package focus

import (
	"errors"
	"strings"
)

// ErrAddCaptureArchivalsourceFor151EmptyTarget is returned when the target key is blank.
var ErrAddCaptureArchivalsourceFor151EmptyTarget = errors.New("focus: target key cannot be empty")

// AddCaptureArchivalsourceFor151Report summarizes deterministic diagnostics for issue #151.
type AddCaptureArchivalsourceFor151Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateAddCaptureArchivalsourceFor151 computes a deterministic AddCaptureArchivalsourceFor151Report for the supplied inputs.
func EvaluateAddCaptureArchivalsourceFor151(targetKey string, countBase, deltaUsed int) (AddCaptureArchivalsourceFor151Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return AddCaptureArchivalsourceFor151Report{}, ErrAddCaptureArchivalsourceFor151EmptyTarget
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
	return AddCaptureArchivalsourceFor151Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
