package vanerrors

import (
	"fmt"
)

// The log data struct
type logData struct {
	BaseMes     string `json:"message"`
	Severity    any    `json:"severity"`
	Description string `json:"description"`
	Cause       error  `json:"cause"`
}

func (l logData) String(op LoggerOptions) string {
	var result string
	result += l.BaseMes
	if op.ShowSeverity {
		result += fmt.Sprintf(" %v", l.Severity)
	}
	if op.ShowDescription {
		result += " " + l.Description
	}
	if op.ShowCause {
		result += " " + l.Cause.Error()
	}
	return result
}
