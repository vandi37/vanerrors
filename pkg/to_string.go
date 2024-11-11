package vanerrors

import "fmt"

func (e VanError) Error() string {
	var err string

	options := e.Options

	if options.ShowDate {
		err += e.Date.Format("2006-01-02 15:04:05 ")
	}

	if options.ShowSeverity {
		if options.IntSeverity {
			err += SeverityArray[e.Severity] + " "
		} else {
			err += fmt.Sprintf("level: %d ", e.Severity)
		}
	}

	if options.ShowCode {
		err += fmt.Sprintf("%d", e.Code) + " "
	}

	err += e.Name + " "

	if options.ShowMessage {
		err += e.Message + " "
	}

	if options.ShowDescription {
		description := make([]byte, 4096)
		n, errRead := e.Description.Read(description)
		if errRead == nil {
			description = description[:n]
		}

		err += "description: " + string(description) + " "
	}

	if options.ShowCause {
		err += "cause: " + e.Cause.Error() + " "
	}
	return err
}
