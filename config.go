package vanerrors

var SeverityArray [4]string = [4]string{
	"info",
	"warn",
	"error",
	"fatal",
}

var DefaultOptions Options = Options{
	ShowCode:    true,
	ShowMessage: true,
}

var DefaultLoggerOptions LoggerOptions = LoggerOptions{
	DoLog: false,
}

var EmptyLoggerOptions LoggerOptions = LoggerOptions{
	DoLog:           true,
	ShowMessage:     true,
	ShowCode:        true,
	ShowSeverity:    true,
	ShowDescription: true,
	ShowCause:       true,
}
