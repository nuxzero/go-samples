package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/nuxzero/simplebank/util"
)

// Custom validator with go-playground/validator useable for gin-gonic/gin
var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
