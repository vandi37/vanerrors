package vanerrors

import (
	"log"
)

// Creates a log of VanError based on it settings
//
// The method could be used inside some methods (New(), VanError.Error()) and outside
// err := Default(Name, Message, Code, Logger)
// err.Log()
//
// It is a basic logger, using the standard log package
// If you want to have better log, set logs of.
func (e VanError) Log() {
	result := e.getView(true)

	// Getting the result string
	logger := log.New(e.Logger, "", 0)
	logger.Println(result)
}
