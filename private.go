package vanerrors

import (
	"encoding/json"
	"fmt"
	"strings"
)

func topicView(topic string, data any) string {
	return fmt.Sprintf("%s: %v", topic, data)
}

func noTopicView(topic string, data any) string {
	return fmt.Sprintf("%v", data)
}

type viewStandard struct {
	name       string
	viewMethod func(topic string, data any) string
}

var viewOrder []viewStandard = []viewStandard{
	{
		name:       "date",
		viewMethod: noTopicView,
	},
	{
		name:       "code",
		viewMethod: noTopicView,
	},
	{
		name:       "main",
		viewMethod: noTopicView,
	},
	{
		name:       "description",
		viewMethod: topicView,
	},
	{
		name:       "cause",
		viewMethod: topicView,
	},
}

func (v JsonVanError) get(s string) any {
	switch s {
	case "date":
		return v.Date
	case "code":
		return v.Code
	case "main":
		return v.Main
	case "description":
		return v.Description
	case "cause":
		return v.Cause
	default:
		return nil
	}
}

func (v JsonVanError) toJson() string {
	result, err := json.Marshal(v)
	if err != nil {
		panic(NewWrap("Failed to convert error to json", err, EmptyHandler))
	}
	return string(result)
}

func (v JsonVanError) toString() string {
	var result string
	for _, s := range viewOrder {
		data := v.get(s.name)
		if data == nil || data == "" || data == 0 {
			continue
		}

		result += s.viewMethod(s.name, data)
		if s.name == "date" || s.name == "code" {
			result += " "
			continue
		}
		result += ", "
	}

	// So it trims only " " if it does not have a comma
	result = strings.TrimSuffix(result, " ")
	result = strings.TrimSuffix(result, ",")

	return result
}

func (err *VanError) toView(LogType bool) JsonVanError {
	opt := err.Options
	if LogType {
		opt = err.LoggerOptions.Options
	}

	var result JsonVanError

	// Adding date
	if opt.ShowDate {
		result.Date = err.Date.Format(DATE_FORMAT)
	}

	// Adding code
	if opt.ShowCode {
		result.Code = err.Code
	}

	// Adding name
	var main string
	main = err.Name

	// Adding message
	if msg := strings.TrimSpace(err.Message); msg != "" && opt.ShowMessage {
		main += ": " + err.Message
	}

	// Setting the main
	result.Main = main

	// Adding description
	if dsc := strings.TrimSpace(err.Description); dsc != "" && opt.ShowDescription {
		result.Description = err.Description
	}

	// Adding cause
	if opt.ShowCause && err.Cause != nil {
		result.Cause = err.Cause
	}
	return result
}

func (err *VanError) getView(LogType bool) string {
	view_map := err.toView(LogType)
	opt := err.Options
	if LogType {
		opt = err.LoggerOptions.Options
	}

	if opt.ShowAsJson {
		return view_map.toJson()
	}

	return view_map.toString()
}
