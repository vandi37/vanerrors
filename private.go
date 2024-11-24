package vanerrors

import (
	"encoding/json"
	"fmt"
	"io"
)

func topicView(topic string, data any) string {
	return fmt.Sprintf("%s: %v", topic, data)
}

func noTopicView(_ string, data any) string {
	return fmt.Sprintf("%v", data)
}

type viewStandard struct {
	name       string
	viewMethod func(string, any) string
}

var viewOrder []viewStandard = []viewStandard{
	{
		name:       "date",
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

type viewMap map[string]any

func (v viewMap) toJson() string {
	result, _ := json.Marshal(v)
	return string(result)
}

func (v viewMap) toString() string {
	var result string
	var i int = -1
	for _, s := range viewOrder {
		data := v[s.name]
		if data == nil {
			continue
		}
		i++
		result += s.viewMethod(s.name, data)
		if s.name == "date" {
			continue
		}
		if i < len(v)-1 {
			result += ", "
		}
	}
	return result
}

func (err VanError) toViewMap(LogType bool) viewMap {
	opt := err.Options
	if LogType {
		opt = err.LoggerOptions.Options
	}
	var result viewMap = viewMap{}

	// Adding date
	if opt.ShowDate {
		result["date"] = err.Date.Format("2006/01/02 15:04:05 ")
	}

	// Adding main info
	var main string

	// Adding code
	if opt.ShowCode {
		main += fmt.Sprintf("%d ", err.Code)
	}

	// Adding name
	main += err.Name

	// Adding message
	if opt.ShowMessage {
		main += ": " + err.Message
	}

	// Setting the main
	result["main"] = main

	// Adding description
	if opt.ShowDescription && err.Description != nil {
		description, _ := io.ReadAll(err.Description)
		if len(description) > 0 {
			result["description"] = string(description)
		}
	}

	// Adding cause
	if opt.ShowCause && err.Cause != nil {
		result["cause"] = err.Cause.Error()
	}

	return result
}

func (err VanError) getView(LogType bool) string {
	view_map := err.toViewMap(LogType)
	opt := err.Options
	if LogType {
		opt = err.LoggerOptions.Options
	}

	if opt.ShowAsJson {
		return view_map.toJson()
	}

	return view_map.toString()
}
