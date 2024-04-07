package repository

import (
	"context"

	"at-home-assessments/model"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type EmployeeRepo struct {
	MongoCollection *mongo.Collection
}

func (e *EmployeeRepo) InsertEmployee(employee *model.Employee) (interface{}, error) {
	result, err := e.MongoCollection.InsertOne(context.Background(), employee)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (e *EmployeeRepo) FindEmployeeByID(employeeID string) (*model.Employee, error) {
	var employee model.Employee
	err := e.MongoCollection.FindOne(context.Background(), bson.M{"employee_id": employeeID}).Decode(&employee)
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (e *EmployeeRepo) UpdateEmployee(employee *model.Employee) error {
	_, err := e.MongoCollection.ReplaceOne(context.Background(), bson.M{"employee_id": employee.EmployeeID}, employee)
	if err != nil {
		return err
	}

	return nil
}

func (e *EmployeeRepo) DeleteEmployee(employeeID string) error {
	_, err := e.MongoCollection.DeleteOne(context.Background(), bson.M{"employee_id": employeeID})
	if err != nil {
		return err
	}

	return nil
}

func (e *EmployeeRepo) GetEmployeesByJob(job string) ([]*model.Employee, error) {
	filter := bson.M{
		"current_job": "",
	}

	ctx := context.Background()
	cursor, err := e.MongoCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var employees []*model.Employee
	for cursor.Next(ctx) {
		var employee model.Employee
		err := cursor.Decode(&employee)
		if err != nil {
			return nil, err
		}

		employees = append(employees, &employee)
	}

	return employees, nil
}
