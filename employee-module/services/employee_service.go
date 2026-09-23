package services

import (
	"fmt"

	"employee-module/models"
	"employee-module/repositories"
)

// EmployeeService содержит бизнес-логику учёта сотрудников.
// Слой Application в многослойной архитектуре: не знает о деталях
// хранения данных, работает с репозиториями через их интерфейс.
type EmployeeService struct {
	employeeRepo  *repositories.EmployeeRepository
	changeLogRepo *repositories.ChangeLogRepository
}

// NewEmployeeService создаёт сервис с внедрёнными зависимостями (Dependency Injection).
func NewEmployeeService(employeeRepo *repositories.EmployeeRepository, changeLogRepo *repositories.ChangeLogRepository) *EmployeeService {
	return &EmployeeService{
		employeeRepo:  employeeRepo,
		changeLogRepo: changeLogRepo,
	}
}

// HireEmployee принимает нового сотрудника на работу.
func (s *EmployeeService) HireEmployee(id int, fullName, phone, email string, departmentID, positionID int) *models.Employee {
	employee := models.NewEmployee(id, fullName, phone, email, departmentID, positionID)
	s.employeeRepo.Save(employee)
	return employee
}

// DismissEmployee увольняет сотрудника и фиксирует изменение в журнале.
func (s *EmployeeService) DismissEmployee(id int) error {
	employee, ok := s.employeeRepo.GetByID(id)
	if !ok {
		return fmt.Errorf("сотрудник с id %d не найден", id)
	}

	oldStatus := employee.Status
	employee.Dismiss()
	s.employeeRepo.Save(employee)

	s.changeLogRepo.Add(&models.ChangeLog{
		EmployeeID: id,
		Action:     "status_change",
		OldValue:   oldStatus,
		NewValue:   employee.Status,
	})

	return nil
}

// TransferEmployee переводит сотрудника в другое подразделение/должность
// и фиксирует изменение в журнале.
func (s *EmployeeService) TransferEmployee(id, departmentID, positionID int) error {
	employee, ok := s.employeeRepo.GetByID(id)
	if !ok {
		return fmt.Errorf("сотрудник с id %d не найден", id)
	}

	oldDepartment := fmt.Sprintf("%d", employee.DepartmentID)
	employee.Transfer(departmentID, positionID)
	s.employeeRepo.Save(employee)

	s.changeLogRepo.Add(&models.ChangeLog{
		EmployeeID: id,
		Action:     "department_change",
		OldValue:   oldDepartment,
		NewValue:   fmt.Sprintf("%d", departmentID),
	})

	return nil
}

// ListActiveEmployees возвращает список сотрудников со статусом "active".
func (s *EmployeeService) ListActiveEmployees() []*models.Employee {
	active := make([]*models.Employee, 0)
	for _, employee := range s.employeeRepo.GetAll() {
		if employee.Status == "active" {
			active = append(active, employee)
		}
	}
	return active
}

// GetChangeHistory возвращает историю изменений конкретного сотрудника.
func (s *EmployeeService) GetChangeHistory(employeeID int) []*models.ChangeLog {
	return s.changeLogRepo.GetByEmployeeID(employeeID)
}
