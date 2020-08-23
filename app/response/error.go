package response

import (
	"html/template"
	"net/http"
)

// ErrorData is a data for view.
type ErrorData struct {
	Code    int
	Message string
}

// Error responses a error view.
func (r *Response) Error(w http.ResponseWriter, code int) {
	e := &ErrorData{
		Code:    code,
		Message: handleErrorMessage(code),
	}

	tpl := template.Must(template.ParseFiles("view/layout/base.tpl", "view/error/index.tpl"))

	w.WriteHeader(e.Code)

	if err := tpl.ExecuteTemplate(w, "base", e); err != nil {
		w.Write([]byte(err.Error()))
	}
}

// handleErrorMessage handles an error message by code.
func handleErrorMessage(code int) string {
	switch code {
	case http.StatusInternalServerError:
		return "An unexepected condition has occured"
	}
	return ""
}
