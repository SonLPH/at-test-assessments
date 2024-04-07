package model

type Booking struct {
	BookingID     string `json:"booking_id,omitempty" bson:"booking_id"`
	OfEmployeeID  string `json:"of_employee_id,omitempty" bson:"of_employee_id"`
	OfUserID      string `json:"of_user_id,omitempty" bson:"of_user_id"`
	BookingTotal  string `json:"booking_total,omitempty" bson:"booking_total"`
	BookingType   string `json:"booking_type,omitempty" bson:"booking_type"`
	BookingDate   string `json:"booking_date,omitempty" bson:"booking_date"`
	BookingStatus string `json:"booking_status,omitempty" bson:"booking_status"`
}
