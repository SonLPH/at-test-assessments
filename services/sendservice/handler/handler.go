package handler

import (
	"at-home-assessments/model"
	"at-home-assessments/repository"
	"net/http"
	"strconv"

	"encoding/json"

	"go.mongodb.org/mongo-driver/mongo"
)

type SendService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

type Request struct {
	BookingID string `json:"booking_id"`
}

func (svc *SendService) MockEmployee(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 5; i++ {
		employee := model.Employee{
			EmployeeID: strconv.Itoa(i),
			Name:       "Employee " + strconv.Itoa(i),
			CurrentJob: "",
		}
		repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
		repo.InsertEmployee(&employee)
	}

	w.WriteHeader(http.StatusCreated)
}

func (svc *SendService) Send(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}

	var req Request
	var employee model.Employee

	defer json.NewEncoder(w).Encode(res)
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = err.Error()
		return
	}

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	freeEmployees, err := repo.GetEmployeesByJob("")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		return
	}

	if len(freeEmployees) == 0 {
		w.WriteHeader(http.StatusNotFound)
		res.Error = "No free employees available"
		return
	}

	employee = *freeEmployees[0]
	employee.CurrentJob = req.BookingID
	err = repo.UpdateEmployee(&employee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		return
	}

	res.Data = employee.EmployeeID
	w.WriteHeader(http.StatusCreated)
}
