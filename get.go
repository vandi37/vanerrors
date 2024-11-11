package vanerrors

// Gets the VanError if the error is it
//
// e - any error
//
// if error is a van error returns the van error
// if error is nol a van error returns nil
func Get(e error) *VanError {
	vanErr, ok := e.(VanError)
	if !ok {
		return nil
	}
	return &vanErr
}
