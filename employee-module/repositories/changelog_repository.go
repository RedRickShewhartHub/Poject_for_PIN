package repositories

import "employee-module/models"

// ChangeLogRepository хранит записи об изменениях данных сотрудников.
type ChangeLogRepository struct {
	logs []*models.ChangeLog
}

// NewChangeLogRepository создаёт пустой репозиторий журнала изменений.
func NewChangeLogRepository() *ChangeLogRepository {
	return &ChangeLogRepository{
		logs: make([]*models.ChangeLog, 0),
	}
}

// Add добавляет новую запись в журнал изменений.
func (r *ChangeLogRepository) Add(log *models.ChangeLog) {
	r.logs = append(r.logs, log)
}

// GetByEmployeeID возвращает все записи изменений конкретного сотрудника.
func (r *ChangeLogRepository) GetByEmployeeID(employeeID int) []*models.ChangeLog {
	result := make([]*models.ChangeLog, 0)
	for _, log := range r.logs {
		if log.EmployeeID == employeeID {
			result = append(result, log)
		}
	}
	return result
}
