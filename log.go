package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func Log(errorData ErrorData, loggerOptions LoggerOptions) error {
	// Setting the logger output
	log.SetOutput(errorData.Logger)

	// creating log data
	var data logData

	// Creating a string for logging
	var logStr string

	// Adding code
	if loggerOptions.ShowCode {
		logStr += fmt.Sprintf("%d ", errorData.Code)
	}

	// Adding name
	logStr += errorData.Name
	// Adding message
	if trimMessage := strings.TrimSpace(errorData.Message); trimMessage != "" && loggerOptions.ShowMessage {
		logStr += fmt.Sprintf(": %s", errorData.Message)
	}

	// Adding the string
	data.BaseMes = logStr

	// Adding severity
	if loggerOptions.ShowSeverity {
		var severity any
		if loggerOptions.IntSeverity {
			severity = SeverityArray[errorData.Severity]
		} else {
			severity = errorData.Severity
		}

		data.Severity = severity

	}

	// Adding description
	if loggerOptions.ShowDescription && errorData.Description != nil {
		description := make([]byte, 4096)
		n, err := errorData.Description.Read(description)
		if err == nil {
			description = description[:n]
		}

		data.Description = string(description)
	}

	// Adding cause
	if loggerOptions.ShowCause && errorData.Cause != nil {
		data.Cause = errorData.Cause
	}

	// Getting the result string
	logger := log.New(errorData.Logger, "", log.LstdFlags|log.Llongfile|log.Llongfile)
	if loggerOptions.LogWith {
		result, err := json.Marshal(data)
		if err != nil {
			err, _ := New(
				ErrorData{
					Name:     "severity is not valid",
					Message:  fmt.Sprintf("$d is les than 0 or bigger than 3", errorData.Severity),
					Code:     400,
					Severity: 2,
				},
				DefaultOptions,
				EmptyLoggerOptions,
			)
			return err
		}
		if errorData.Severity == 3 {

			logger.Fatalln(string(result))
		}
		logger.Println(string(result))
	} else {
		result := data.String(loggerOptions)
		if errorData.Severity == 3 {

			logger.Fatalln(result)
		}
		logger.Println(result)
	}
	return nil
}
