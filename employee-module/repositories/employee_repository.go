package repositories

import "employee-module/models"

// EmployeeRepository инкапсулирует хранение сотрудников.
// В реальном проекте это место для замены на работу с реальной БД,
// не затрагивая слой сервисов.
type EmployeeRepository struct {
	employees map[int]*models.Employee
}

// NewEmployeeRepository создаёт пустой репозиторий сотрудников.
func NewEmployeeRepository() *EmployeeRepository {
	return &EmployeeRepository{
		employees: make(map[int]*models.Employee),
	}
}

// GetByID возвращает сотрудника по идентификатору.
func (r *EmployeeRepository) GetByID(id int) (*models.Employee, bool) {
	employee, ok := r.employees[id]
	return employee, ok
}

// GetAll возвращает всех сотрудников.
func (r *EmployeeRepository) GetAll() []*models.Employee {
	result := make([]*models.Employee, 0, len(r.employees))
	for _, employee := range r.employees {
		result = append(result, employee)
	}
	return result
}

// Save сохраняет сотрудника (создание или обновление).
func (r *EmployeeRepository) Save(employee *models.Employee) {
	r.employees[employee.ID] = employee
}
