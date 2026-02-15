package pages

import (
	"net/http"

	"github.com/fedotovmax/green-api-test/internal/templates/home"
	"github.com/fedotovmax/green-api-test/internal/utils"
	"github.com/go-chi/chi/v5"
)

type controller struct {
	staticDirectoryPath string
}

func New(staticDirectoryPath string) *controller {
	return &controller{
		staticDirectoryPath: staticDirectoryPath,
	}
}

func (c *controller) Register(router chi.Router) {

	router.Handle("/public/*", http.StripPrefix(
		"/public/",
		http.FileServer(http.Dir(c.staticDirectoryPath)),
	))

	router.Get("/", c.home)

}

func (c *controller) home(w http.ResponseWriter, req *http.Request) {

	err := utils.RenderTemplate(w, req, home.Page())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
