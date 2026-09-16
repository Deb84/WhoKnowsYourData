package valueshandler

import (
	"whoknowsyourdata/handlers"
	valueService "whoknowsyourdata/services/values"
)

type Handler = handlers.Handler

type ValueHandler struct {
	*Handler
	ValueService *valueService.ValueService
}

func NewValueHandler(handler *Handler, valueService *valueService.ValueService) *ValueHandler {
	return &ValueHandler{
		Handler:      handler,
		ValueService: valueService,
	}
}
