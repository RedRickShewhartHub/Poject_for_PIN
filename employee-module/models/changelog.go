package models

// ChangeLog фиксирует изменение данных сотрудника.
type ChangeLog struct {
	EmployeeID int
	Action     string
	OldValue   string
	NewValue   string
}
