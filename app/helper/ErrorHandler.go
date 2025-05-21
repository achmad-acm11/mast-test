package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"mast-integrator/app/exception"
	"strings"
)

func ErrorHandler(err error) {
	if err != nil {
		panic(exception.NewInternalServerError(err.Error()))
	}
}

func customMessage(field string, typeMessage string, param string) string {
	var msg string
	switch typeMessage {
	case "required":
		msg = fmt.Sprintf("%s %s", field, "field is required")
	case "oneof":
		msg = fmt.Sprintf("%s %s", field, "field is oneof "+param)
	}
	return msg
}

func ErrorHandlerValidator(err error) {
	if err != nil {
		errors := make(map[string]string)

		for _, e := range err.(validator.ValidationErrors) {
			msg := customMessage(strings.ToLower(e.Field()), e.Tag(), e.Param())
			errors[strings.ToLower(e.Field())] = msg
		}
		panic(exception.NewValidationError(errors))
	}
}

func GetMessageOneErrorValidator(err error, field string) string {
	var msg string
	for _, e := range err.(validator.ValidationErrors) {
		msg = customMessage(strings.ToLower(field), e.Tag(), e.Param())
	}
	return msg
}

func ValidateFormData(c *gin.Context, rules map[string]string, validate *validator.Validate) (map[string]string, bool) {
	messagesError := make(map[string]string)
	isError := false
	for key, rule := range rules {
		val := c.PostForm(key)
		if err := validate.Var(val, rule); err != nil {
			msg := GetMessageOneErrorValidator(err, key)
			messagesError[strings.ToLower(key)] = msg
			isError = true
		}
	}
	return messagesError, isError
}
