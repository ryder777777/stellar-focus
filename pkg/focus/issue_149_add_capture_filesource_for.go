package focus

import (
	"errors"
	"strings"
)

// ErrAddCaptureFilesourceFor149EmptyTarget is returned when the target key is blank.
var ErrAddCaptureFilesourceFor149EmptyTarget = errors.New("focus: target key cannot be empty")

// AddCaptureFilesourceFor149Report summarizes deterministic diagnostics for issue #149.
type AddCaptureFilesourceFor149Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateAddCaptureFilesourceFor149 computes a deterministic AddCaptureFilesourceFor149Report for the supplied inputs.
func EvaluateAddCaptureFilesourceFor149(targetKey string, countBase, deltaUsed int) (AddCaptureFilesourceFor149Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return AddCaptureFilesourceFor149Report{}, ErrAddCaptureFilesourceFor149EmptyTarget
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
	return AddCaptureFilesourceFor149Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
