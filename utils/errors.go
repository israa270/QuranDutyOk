package utils

import (
	"github.com/go-playground/validator/v10"
)

// ErrorMsg struct
type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// GetErrorMsg fn
func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
		case "required":
			return "field is required"
		case "lte":
			return "Should be less than " + fe.Param()
		case "gte":
			return "Should be greater than " + fe.Param()
		case "min":
			return "Should be not less than " + fe.Param()
		case "max":
			return "should be not greater than " + fe.Param()
		case "e164":
			return "should be phone number" + fe.Param()
		case "postcode_iso3166_alpha2=GB":
			return "should be be postal code" + fe.Param()
		case "email":
			return "should be email" + fe.Param()
		case "alphanumunicode":
			return "should be letters, numbers, special character" + fe.Param()
		case "iso3166_1_alpha2":
			return "should be country code" + fe.Param()
		case "len":
			return "should be len "+fe.Param()
		case "eqfield":
			return "confirm password should be same password "+fe.Param()	
	}
	return "Unknown error"
}

// HandleError handle error from validator
func HandleError(ve validator.ValidationErrors) []ErrorMsg {
	out := make([]ErrorMsg, len(ve))
	for i, fe := range ve {
		out[i] = ErrorMsg{Field: fe.Field(), Message: GetErrorMsg(fe)}
	}
	return out
}
