package departament

import (
	"TestTaskAPI/pkg/res"
	"net/http"
)

type DepartHandlerDeps struct {
	DepartRepository *DepartRepository
}

type DepartHandler struct {
	DepartRepository *DepartRepository
}

func NewDepartHandler(router *http.ServeMux, deps DepartHandlerDeps) {
	handler := &DepartHandler{
		DepartRepository: deps.DepartRepository,
	}
	router.HandleFunc("POST /departments", handler.CreateDepart())
	router.HandleFunc("POST /departments/{id}/employees/", handler.CreateEmployees())
	router.HandleFunc("PATCH /departments/{id}", handler.Update())
	router.HandleFunc("DELETE /departments/{id}", handler.Delete())
	router.HandleFunc("GET /departments/{id}", handler.GoTo())
}

func (handler *DepartHandler) CreateDepart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		createdDepart, err := handler.DepartRepository.Create()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Jsons(w, createdDepart, 201)
	}
}

func (handler *DepartHandler) CreateEmployees() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		createdDepart, err := handler.EmployeeRequest.Create()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Jsons(w, createdDepart, 201)
	}
}

func (handler *DepartHandler) Update() http.HandlerFunc {

}

func (handler *DepartHandler) Delete() http.HandlerFunc {

}

func (handler *DepartHandler) GoTo() http.HandlerFunc {

}
