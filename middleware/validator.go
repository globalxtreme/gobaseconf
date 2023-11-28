package middleware

import (
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/response/error"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
	"time"
)

type Validator struct{}

func (v Validator) Make(r *http.Request, rules interface{}) {
	err := config.XtremeValidate.Struct(rules)
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

func (v Validator) RegisterValidation(callback func(validate *validator.Validate)) {
	config.XtremeValidate = validator.New()

	_ = config.XtremeValidate.RegisterValidation("date_ddmmyyyy", dateDDMMYYYYValidation)
	_ = config.XtremeValidate.RegisterValidation("time_hhmm", dateHHMMValidation)
	_ = config.XtremeValidate.RegisterValidation("time_hhmmss", dateHHMMSSValidation)

	callback(config.XtremeValidate)
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

func dateDDMMYYYYValidation(fl validator.FieldLevel) bool {
	_, err := time.Parse("02/01/2006", fl.Field().String())
	return err == nil
}

func dateHHMMValidation(fl validator.FieldLevel) bool {
	_, err := time.Parse("15:04", fl.Field().String())
	return err == nil
}

func dateHHMMSSValidation(fl validator.FieldLevel) bool {
	_, err := time.Parse("15:04:05", fl.Field().String())
	return err == nil
}
