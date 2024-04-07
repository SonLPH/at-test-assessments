package handler

import (
	"at-home-assessments/model"
	"at-home-assessments/repository"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

type SendRequest struct {
	BookingID string `json:"booking_id"`
}

func (svc *BookingService) CreateBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)

	var booking model.Booking

	err := json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = err.Error()
		return
	}

	booking.BookingID = uuid.New().String()
	booking.BookingStatus = "pending"
	booking.BookingDate = time.Now().Format("2006-01-02 15:04:05")

	// booking.OfEmployeeID = "1"
	repo := repository.BookingRepo{MongoCollection: svc.MongoCollection}

	sendReq := SendRequest{BookingID: booking.BookingID}
	sendRes := &Response{}
	sendBody, err := json.Marshal(sendReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		return
	}

	sendURL := "http://localhost:4446/send"
	sendResp, err := http.Post(sendURL, "application/json", bytes.NewBuffer(sendBody))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		return
	}
	defer sendResp.Body.Close()
	if err := json.NewDecoder(sendResp.Body).Decode(sendRes); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		return
	}

	if sendResp.StatusCode != http.StatusCreated {
		w.WriteHeader(sendResp.StatusCode)
		res.Error = sendRes.Error
		return
	}

	booking.OfEmployeeID = sendRes.Data.(string)

	insertID, err := repo.InsertBooking(&booking)

	fmt.Print(err)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		return
	}

	res.Data = booking.BookingID

	w.WriteHeader(http.StatusCreated)
	log.Println("Booking created successfully", insertID)
}
