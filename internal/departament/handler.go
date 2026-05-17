package departament

import "net/http"

type DepartHandlerDeps struct {
	LinkRepository *DepartRepository
}

type Departandler struct {
	LinkRepository *DepartRepository
}

func NewDepartHandler(router *http.ServeMux, deps DepartHandlerDeps) {
	handler := &DepartkHandler{
		LinkRepository: deps.DepartRepository,
	}
	router.HandleFunc("POST /departments", handler.CreateDepart())
	router.HandleFunc("POST /departments/{id}/employees/", handler.CreateEmployees())
	router.HandleFunc("PATCH /departments/{id}", handler.Update())
	router.HandleFunc("DELETE /departments/{id}", handler.Delete())
	router.HandleFunc("GET //departments/{id}", handler.GoTo())
}

func (handler *DepartHandler) Create() http.HandlerFunc {}

func (handler *DepartHandler) Create() http.HandlerFunc {}

func (handler *DepartHandler) Update() http.HandlerFunc {}

func (handler *DepartHandler) Delete() http.HandlerFunc {}

func (handler *DepartHandler) GoTo() http.HandlerFunc {}
