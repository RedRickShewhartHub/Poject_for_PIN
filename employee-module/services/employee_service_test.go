package services

import (
	"testing"

	"employee-module/repositories"
)

func createTestService() *EmployeeService {
	employeeRepo := repositories.NewEmployeeRepository()
	changeLogRepo := repositories.NewChangeLogRepository()

	return NewEmployeeService(employeeRepo, changeLogRepo)
}

func TestHireEmployee(t *testing.T) {
	service := createTestService()

	employee := service.HireEmployee(
		1,
		"Иванов Иван Иванович",
		"+79001112233",
		"ivanov@company.ru",
		1,
		2,
	)

	if employee == nil {
		t.Fatal("сотрудник не должен быть nil")
	}

	if employee.ID != 1 {
		t.Errorf("ожидался ID 1, получен %d", employee.ID)
	}

	if employee.FullName != "Иванов Иван Иванович" {
		t.Errorf(
			"ожидалось ФИО Иванов Иван Иванович, получено %s",
			employee.FullName,
		)
	}

	if employee.Status != "active" {
		t.Errorf(
			"ожидался статус active, получен %s",
			employee.Status,
		)
	}
}

func TestDismissEmployee(t *testing.T) {
	service := createTestService()

	service.HireEmployee(
		1,
		"Иванов Иван Иванович",
		"+79001112233",
		"ivanov@company.ru",
		1,
		2,
	)

	err := service.DismissEmployee(1)

	if err != nil {
		t.Fatalf("неожиданная ошибка: %v", err)
	}

	history := service.GetChangeHistory(1)

	if len(history) != 1 {
		t.Fatalf(
			"ожидалась 1 запись в истории, получено %d",
			len(history),
		)
	}

	log := history[0]

	if log.Action != "status_change" {
		t.Errorf(
			"ожидалось действие status_change, получено %s",
			log.Action,
		)
	}

	if log.OldValue != "active" {
		t.Errorf(
			"ожидалось старое значение active, получено %s",
			log.OldValue,
		)
	}

	if log.NewValue != "dismissed" {
		t.Errorf(
			"ожидалось новое значение dismissed, получено %s",
			log.NewValue,
		)
	}
}

func TestDismissEmployeeNotFound(t *testing.T) {
	service := createTestService()

	err := service.DismissEmployee(999)

	if err == nil {
		t.Error("ожидалась ошибка при увольнении несуществующего сотрудника")
	}
}

func TestTransferEmployee(t *testing.T) {
	service := createTestService()

	service.HireEmployee(
		1,
		"Иванов Иван Иванович",
		"+79001112233",
		"ivanov@company.ru",
		1,
		2,
	)

	err := service.TransferEmployee(1, 3, 5)

	if err != nil {
		t.Fatalf("неожиданная ошибка: %v", err)
	}

	history := service.GetChangeHistory(1)

	if len(history) != 1 {
		t.Fatalf(
			"ожидалась 1 запись в истории, получено %d",
			len(history),
		)
	}

	if history[0].Action != "department_change" {
		t.Errorf(
			"ожидалось действие department_change, получено %s",
			history[0].Action,
		)
	}

	if history[0].OldValue != "1" {
		t.Errorf(
			"ожидалось старое подразделение 1, получено %s",
			history[0].OldValue,
		)
	}

	if history[0].NewValue != "3" {
		t.Errorf(
			"ожидалось новое подразделение 3, получено %s",
			history[0].NewValue,
		)
	}
}

func TestTransferEmployeeNotFound(t *testing.T) {
	service := createTestService()

	err := service.TransferEmployee(999, 3, 5)

	if err == nil {
		t.Error("ожидалась ошибка при переводе несуществующего сотрудника")
	}
}

func TestListActiveEmployees(t *testing.T) {
	service := createTestService()

	service.HireEmployee(
		1,
		"Иванов Иван Иванович",
		"+79001112233",
		"ivanov@company.ru",
		1,
		1,
	)

	service.HireEmployee(
		2,
		"Петрова Мария Сергеевна",
		"+79004445566",
		"petrova@company.ru",
		1,
		2,
	)

	err := service.DismissEmployee(1)

	if err != nil {
		t.Fatalf("неожиданная ошибка: %v", err)
	}

	active := service.ListActiveEmployees()

	if len(active) != 1 {
		t.Fatalf(
			"ожидался 1 активный сотрудник, получено %d",
			len(active),
		)
	}

	if active[0].ID != 2 {
		t.Errorf(
			"ожидался активный сотрудник с ID 2, получен ID %d",
			active[0].ID,
		)
	}

	if active[0].Status != "active" {
		t.Errorf(
			"ожидался статус active, получен %s",
			active[0].Status,
		)
	}
}

func TestGetChangeHistory(t *testing.T) {
	service := createTestService()

	service.HireEmployee(
		1,
		"Иванов Иван Иванович",
		"+79001112233",
		"ivanov@company.ru",
		1,
		1,
	)

	service.DismissEmployee(1)

	history := service.GetChangeHistory(1)

	if len(history) != 1 {
		t.Fatalf(
			"ожидалась 1 запись истории, получено %d",
			len(history),
		)
	}

	if history[0].EmployeeID != 1 {
		t.Errorf(
			"ожидался EmployeeID 1, получен %d",
			history[0].EmployeeID,
		)
	}
}

// Интеграционный тест.
// Проверяет взаимодействие EmployeeService,
// EmployeeRepository и ChangeLogRepository.
func TestEmployeeLifecycleIntegration(t *testing.T) {
	employeeRepo := repositories.NewEmployeeRepository()
	changeLogRepo := repositories.NewChangeLogRepository()

	service := NewEmployeeService(
		employeeRepo,
		changeLogRepo,
	)

	// 1. Принимаем сотрудника на работу.
	employee := service.HireEmployee(
		1,
		"Иванов Иван Иванович",
		"+79001112233",
		"ivanov@company.ru",
		1,
		2,
	)

	if employee == nil {
		t.Fatal("сотрудник не был создан")
	}

	// Проверяем, что сотрудник действительно сохранился
	// в EmployeeRepository.
	savedEmployee, ok := employeeRepo.GetByID(1)

	if !ok {
		t.Fatal("сотрудник не найден в EmployeeRepository")
	}

	if savedEmployee.FullName != "Иванов Иван Иванович" {
		t.Errorf(
			"неверное ФИО сохранённого сотрудника: %s",
			savedEmployee.FullName,
		)
	}

	// 2. Переводим сотрудника.
	err := service.TransferEmployee(1, 3, 5)

	if err != nil {
		t.Fatalf("ошибка при переводе: %v", err)
	}

	// Проверяем изменение данных в репозитории.
	savedEmployee, ok = employeeRepo.GetByID(1)

	if !ok {
		t.Fatal("сотрудник пропал из EmployeeRepository")
	}

	if savedEmployee.DepartmentID != 3 {
		t.Errorf(
			"ожидался DepartmentID 3, получен %d",
			savedEmployee.DepartmentID,
		)
	}

	if savedEmployee.PositionID != 5 {
		t.Errorf(
			"ожидался PositionID 5, получен %d",
			savedEmployee.PositionID,
		)
	}

	// 3. Увольняем сотрудника.
	err = service.DismissEmployee(1)

	if err != nil {
		t.Fatalf("ошибка при увольнении: %v", err)
	}

	// Проверяем изменение статуса.
	savedEmployee, ok = employeeRepo.GetByID(1)

	if !ok {
		t.Fatal("сотрудник не найден после увольнения")
	}

	if savedEmployee.Status != "dismissed" {
		t.Errorf(
			"ожидался статус dismissed, получен %s",
			savedEmployee.Status,
		)
	}

	// 4. Проверяем журнал изменений.
	history := service.GetChangeHistory(1)

	if len(history) != 2 {
		t.Fatalf(
			"ожидалось 2 записи в истории, получено %d",
			len(history),
		)
	}

	// Первая запись — перевод.
	if history[0].Action != "department_change" {
		t.Errorf(
			"первая запись: ожидалось department_change, получено %s",
			history[0].Action,
		)
	}

	// Вторая запись — увольнение.
	if history[1].Action != "status_change" {
		t.Errorf(
			"вторая запись: ожидалось status_change, получено %s",
			history[1].Action,
		)
	}

	if history[1].OldValue != "active" {
		t.Errorf(
			"для увольнения ожидалось старое значение active, получено %s",
			history[1].OldValue,
		)
	}

	if history[1].NewValue != "dismissed" {
		t.Errorf(
			"для увольнения ожидалось новое значение dismissed, получено %s",
			history[1].NewValue,
		)
	}
}
