package validator

import (
	"log"
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/validator"
)

type ValidationErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()
var once sync.Once

func Init() {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})
}

func ValidateStruct(object interface{}) any {
	once.Do(func() {
		Init()
	})
	err := validate.Struct(object)
	if err != nil {
		var errors []ValidationErrorResponse
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationErrorResponse
			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, element)
			log.Println(err.StructNamespace())
		}
		return errors
	}

	return nil
}
