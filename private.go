package vanerrors

import (
	"encoding/json"
	"fmt"
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

func jsonView(topic string, data any) string {
	json_data, err := json.Marshal(data)
	if err != nil {
		json_data = []byte(fmt.Sprintf(`"%v"`, data))
	}
	return fmt.Sprintf(`"%s":%s`, topic, json_data)
}

type viewMap map[string]any

func (v viewMap) toJson() string {
	var result string = `{`
	var i int = -1
	for _, s := range viewOrder {
		data := v[s.name]
		if data == nil {
			continue
		}
		i++
		result += jsonView(s.name, data)
		if i < len(v)-1 {
			result += ","
		}
	}
	result += `}`
	return result
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
	var result = viewMap{}

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
		var description []byte
		for {
			buf := make([]byte, 2048)
			n, readErr := err.Description.Read(buf)
			if readErr != nil {
				if len(description) <= 0 {
					break
				}
				result["description"] = string(description)
				break
			}
			description = append(description, buf[:n]...)
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
