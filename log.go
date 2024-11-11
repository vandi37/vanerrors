package vanerrors

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func Log(vanError VanError) error {
	// Setting the logger output
	log.SetOutput(vanError.logger)

	// creating log data
	var data logData

	// Creating a string for logging
	var logStr string

	// Adding code
	if vanError.LoggerOptions.ShowCode {
		logStr += fmt.Sprintf("%d ", vanError.Code)
	}

	// Adding name
	logStr += vanError.Name
	// Adding message
	if trimMessage := strings.TrimSpace(vanError.Message); trimMessage != "" && vanError.LoggerOptions.ShowMessage {
		logStr += fmt.Sprintf(": %s", vanError.Message)
	}

	// Adding the string
	data.BaseMes = logStr

	// Adding severity
	if vanError.LoggerOptions.ShowSeverity {
		var severity any
		if vanError.LoggerOptions.IntSeverity {
			severity = SeverityArray[vanError.Severity]
		} else {
			severity = vanError.Severity
		}

		data.Severity = severity

	}

	// Adding description
	if vanError.LoggerOptions.ShowDescription && vanError.Description != nil {
		description := make([]byte, 4096)
		n, err := vanError.Description.Read(description)
		if err == nil {
			description = description[:n]
		}

		data.Description = string(description)
	}

	// Adding cause
	if vanError.LoggerOptions.ShowCause && vanError.Cause != nil {
		data.Cause = vanError.Cause
	}

	// Getting the result string
	logger := log.New(vanError.logger, "", log.LstdFlags|log.Llongfile|log.Llongfile)
	if vanError.LoggerOptions.LogWith {
		result, err := json.Marshal(data)
		if err != nil {
			err, _ := New(
				ErrorData{
					Name:     "severity is not valid",
					Message:  fmt.Sprintf("%d is les than 0 or bigger than 3", vanError.Severity),
					Code:     400,
					Severity: 2,
				},
				DefaultOptions,
				EmptyLoggerOptions,
			)
			return err
		}
		if vanError.Severity == 3 {

			logger.Fatalln(string(result))
		}
		logger.Println(string(result))
	} else {
		result := data.String(vanError.LoggerOptions)
		if vanError.Severity == 3 {

			logger.Fatalln(result)
		}
		logger.Println(result)
	}
	return nil
}
