package models

import "testing"

func TestNewEmployee(t *testing.T) {
	employee := NewEmployee(
		1,
		"Иванов Иван Иванович",
		"+79001112233",
		"ivanov@company.ru",
		1,
		2,
	)

	if employee.ID != 1 {
		t.Errorf("ожидался ID 1, получен %d", employee.ID)
	}

	if employee.FullName != "Иванов Иван Иванович" {
		t.Errorf("неверное ФИО: %s", employee.FullName)
	}

	if employee.Phone != "+79001112233" {
		t.Errorf("неверный телефон: %s", employee.Phone)
	}

	if employee.Email != "ivanov@company.ru" {
		t.Errorf("неверный email: %s", employee.Email)
	}

	if employee.Status != "active" {
		t.Errorf("ожидался статус active, получен %s", employee.Status)
	}

	if employee.DepartmentID != 1 {
		t.Errorf("ожидался DepartmentID 1, получен %d", employee.DepartmentID)
	}

	if employee.PositionID != 2 {
		t.Errorf("ожидался PositionID 2, получен %d", employee.PositionID)
	}
}

func TestEmployeeDismiss(t *testing.T) {
	employee := NewEmployee(
		1,
		"Иванов Иван Иванович",
		"+79001112233",
		"ivanov@company.ru",
		1,
		2,
	)

	employee.Dismiss()

	if employee.Status != "dismissed" {
		t.Errorf(
			"ожидался статус dismissed, получен %s",
			employee.Status,
		)
	}
}

func TestEmployeeTransfer(t *testing.T) {
	employee := NewEmployee(
		1,
		"Иванов Иван Иванович",
		"+79001112233",
		"ivanov@company.ru",
		1,
		2,
	)

	employee.Transfer(3, 5)

	if employee.DepartmentID != 3 {
		t.Errorf(
			"ожидался DepartmentID 3, получен %d",
			employee.DepartmentID,
		)
	}

	if employee.PositionID != 5 {
		t.Errorf(
			"ожидался PositionID 5, получен %d",
			employee.PositionID,
		)
	}

	// После перевода статус должен остаться active.
	if employee.Status != "active" {
		t.Errorf(
			"после перевода ожидался статус active, получен %s",
			employee.Status,
		)
	}
}
