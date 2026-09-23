package models

// Employee представляет сотрудника компании.
type Employee struct {
	ID           int
	FullName     string
	Phone        string
	Email        string
	Status       string // "active" или "dismissed"
	DepartmentID int
	PositionID   int
}

// NewEmployee создаёт нового сотрудника со статусом "active".
func NewEmployee(id int, fullName, phone, email string, departmentID, positionID int) *Employee {
	return &Employee{
		ID:           id,
		FullName:     fullName,
		Phone:        phone,
		Email:        email,
		Status:       "active",
		DepartmentID: departmentID,
		PositionID:   positionID,
	}
}

// Dismiss переводит сотрудника в статус "уволен".
func (e *Employee) Dismiss() {
	e.Status = "dismissed"
}

// Transfer переводит сотрудника в другое подразделение и на другую должность.
func (e *Employee) Transfer(departmentID, positionID int) {
	e.DepartmentID = departmentID
	e.PositionID = positionID
}
