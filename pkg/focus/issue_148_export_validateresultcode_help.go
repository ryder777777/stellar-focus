package focus

import (
	"errors"
	"strings"
)

// ErrExportValidateresultcodeHelperFor148EmptyTarget is returned when the target key is blank.
var ErrExportValidateresultcodeHelperFor148EmptyTarget = errors.New("focus: target key cannot be empty")

// ExportValidateresultcodeHelperFor148Report summarizes deterministic diagnostics for issue #148.
type ExportValidateresultcodeHelperFor148Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateExportValidateresultcodeHelperFor148 computes a deterministic ExportValidateresultcodeHelperFor148Report for the supplied inputs.
func EvaluateExportValidateresultcodeHelperFor148(targetKey string, countBase, deltaUsed int) (ExportValidateresultcodeHelperFor148Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return ExportValidateresultcodeHelperFor148Report{}, ErrExportValidateresultcodeHelperFor148EmptyTarget
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
	return ExportValidateresultcodeHelperFor148Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
