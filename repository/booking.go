package repository

import (
	"at-home-assessments/model"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type BookingRepo struct {
	MongoCollection *mongo.Collection
}

func (b *BookingRepo) InsertBooking(booking *model.Booking) (interface{}, error) {
	result, err := b.MongoCollection.InsertOne(context.Background(), booking)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (b *BookingRepo) FindBookingByID(bookingID string) (*model.Booking, error) {
	var booking model.Booking
	err := b.MongoCollection.FindOne(context.Background(), bson.M{"booking_id": bookingID}).Decode(&booking)
	if err != nil {
		return nil, err
	}

	return &booking, nil
}

func (b *BookingRepo) UpdateBooking(booking *model.Booking) error {
	_, err := b.MongoCollection.ReplaceOne(context.Background(), bson.M{"booking_id": booking.BookingID}, booking)
	if err != nil {
		return err
	}

	return nil
}

func (b *BookingRepo) DeleteBooking(bookingID string) error {
	_, err := b.MongoCollection.DeleteOne(context.Background(), bson.M{"booking_id": bookingID})
	if err != nil {
		return err
	}

	return nil
}

func (b *BookingRepo) GetBookingsByDayAndEmployee(day string, employeeID string) ([]*model.Booking, error) {
	filter := bson.M{
		"booking_date":   bson.M{"$regex": fmt.Sprintf("^%s", day)},
		"of_employee_id": employeeID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := b.MongoCollection.Find(ctx, filter)
	fmt.Println(cursor)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var bookings []*model.Booking

	for cursor.Next(ctx) {
		var booking model.Booking
		if err := cursor.Decode(&booking); err != nil {
			return nil, err
		}

		bookings = append(bookings, &booking)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(bookings) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return bookings, nil
}
