package departament

import (
	"TestTaskAPI/pkg/req"
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
	router.HandleFunc("PATCH /departments/{id}", handler.UpdateDepart())
	router.HandleFunc("DELETE /departments/{id}", handler.DeleteDepart())
	router.HandleFunc("GET /departments/{id}", handler.GetDepart())
}

func (handler *DepartHandler) CreateDepart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[DepartmentRequest](&w, r)
		if err != nil {
			return
		}
		CreateDep, err := handler.DepartRepository.CreateDepart(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Jsons(w, CreateDep, 201)
	}
}

func (handler *DepartHandler) CreateEmployees() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[EmployeeRequest](&w, r)
		if err != nil {
			return
		}
		CreateEmp, err := handler.DepartRepository.CreateEmpl(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Jsons(w, CreateEmp, 201)
	}
}

func (handler *DepartHandler) UpdateDepart() http.HandlerFunc {

}

func (handler *DepartHandler) DeleteDepart() http.HandlerFunc {

}

func (handler *DepartHandler) GetDepart() http.HandlerFunc {

}
