package middleware

import (
	"github.com/globalxtreme/gobaseconf/response/error"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

type Validator struct{}

func (v Validator) Make(r *http.Request, rules interface{}) {
	validate := validator.New()

	err := validate.Struct(rules)
	if err != nil {
		var attributes []interface{}
		for _, e := range err.(validator.ValidationErrors) {
			attributes = append(attributes, map[string]interface{}{
				"param":   e.Field(),
				"message": getMessage(e.Error()),
			})
		}

		error.ErrXtremeValidation(attributes)
	}
}

func getMessage(errMsg string) string {
	splitMsg := strings.Split(errMsg, ":")
	key := 0
	if len(splitMsg) == 3 {
		key = 2
	} else if len(splitMsg) == 2 {
		key = 1
	}

	return splitMsg[key]
}
