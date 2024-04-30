package core

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


func BindParamAndValidate[T interface{}](c *gin.Context, obj *T) {
    if c.Request.Method == http.MethodGet {
        if err := c.ShouldBindQuery(obj); err != nil {
			panic(NewKnownError(http.StatusBadRequest,nil,err.Error()))
        }
    } else if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut || c.Request.Method == http.MethodDelete {
        if err := c.ShouldBindJSON(obj); err != nil {
			panic(NewKnownError(http.StatusBadRequest,nil,err.Error()))
        }
    }

	// 使用binding tag来验证就好，这里不再重复验证，validator验证出参
    // if err := ValidateStruct(obj); err != nil {
    //     // 如果是验证错误，返回422 Unprocessable Entity
    //     // ErrorHandler(c, NewKnownError(http.StatusUnprocessableEntity,nil,err.Error()))
	// 	panic(NewKnownError(http.StatusUnprocessableEntity,nil,err.Error()))
    //     // return false
    // }
}

var Validate = validator.New()

// ValidateStruct validates the request structure
func ValidateStruct(i interface{}) error {
	err := Validate.Struct(i)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, err := range err.(validator.ValidationErrors) {
			var errorstr string
			switch err.Tag() {
			case "required":
				errorstr = fmt.Sprintf("Field '%s' is required", err.Field())
			case "email":
				errorstr = fmt.Sprintf("Field '%s' must be a valid email address", err.Field())
			case "gte":
				errorstr = fmt.Sprintf("The value of field '%s' must be at least %s", err.Field(), err.Param())
			case "lte":
				errorstr = fmt.Sprintf("The value of field '%s' must be at most %s", err.Field(), err.Param())
			default:
				errorstr = fmt.Sprintf("Field '%s' failed validation (%s)", err.Field(), validationErrorToText(err))
			}
			return fmt.Errorf("%s", errorstr)
		}

	}
	return nil
}


// validationErrorToText converts validation errors to more descriptive text
func validationErrorToText(e validator.FieldError) string {
	switch e.Tag() {
	case "min":
		return fmt.Sprintf("At least %s characters", e.Param())
	case "max":
		return fmt.Sprintf("At most %s characters", e.Param())
	case "len":
		return fmt.Sprintf("Must be %s characters", e.Param())
	// Add more cases as needed
	default:
		return "Does not meet validation rules"
	}
}

