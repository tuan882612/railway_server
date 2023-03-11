package response

import (
	"net/http"
)

type Base struct {
	Data any  `json:"data"`
	Meta Meta `json:"meta"`
}

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Custom(Data any, code int, message string) Base {
	return Base{
		Data: Data,
		Meta: Meta{
			Code:    code,
			Message: message,
		},
	}
}

func Error(err string) Base {
	return Base{
		Data: map[string]string{},
		Meta: Meta{
			Code:    http.StatusBadRequest,
			Message: err,
		},
	}
}

func Empty() Base {
	return Base{
		Data: map[string]string{},
		Meta: Meta{
			Code: http.StatusNoContent,
			Message: "Table doesn't contain data",
		},
	}
}