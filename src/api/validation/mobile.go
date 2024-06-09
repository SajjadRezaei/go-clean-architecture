package validation

import (
	"github.com/SajjadRezaei/go-clean-architecture/common"
	"github.com/go-playground/validator/v10"
)

func IranianMobileNumberValidator(fld validator.FieldLevel) bool {

	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}

	return common.ValidateMobileNumber(value)
}
