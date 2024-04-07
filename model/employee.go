package model

type Employee struct {
	EmployeeID string `json:"employee_id,omitempty" bson:"employee_id"`
	Name       string `json:"name,omitempty" bson:"name"`
	CurrentJob string `json:"current_job,omitempty" bson:"current_job"`
}
