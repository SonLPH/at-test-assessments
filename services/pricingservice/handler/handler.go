package handler

import (
	"at-home-assessments/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
)

type PricingService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

type Request struct {
	Day        string `json:"day"`
	EmployeeID string `json:"employee_id"`
}

func (svc *PricingService) GetPricingByDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}

	var req Request

	defer json.NewEncoder(w).Encode(res)
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = err.Error()
		return
	}

	repo := repository.BookingRepo{MongoCollection: svc.MongoCollection}
	bookings, err := repo.GetBookingsByDayAndEmployee(req.Day, req.EmployeeID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		return
	}

	totalPricing := 0.0
	for _, booking := range bookings {
		bookingTotal, err := strconv.ParseFloat(booking.BookingTotal, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			res.Error = err.Error()
			return
		}
		totalPricing += bookingTotal
	}
	res.Data = map[string]interface{}{
		"employee_id":   req.EmployeeID,
		"total_pricing": totalPricing,
		"day":           req.Day,
	}
	w.WriteHeader(http.StatusOK)
}
