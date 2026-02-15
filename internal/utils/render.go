package utils

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/fedotovmax/green-api-test/internal/keys"
)

func RenderTemplate(w http.ResponseWriter, r *http.Request, component templ.Component) error {
	w.Header().Set(keys.HeaderContentType, keys.ContentTypeTextHTMLUTF8)
	return component.Render(r.Context(), w)
}
