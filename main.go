package main

import (
	"fmt"

	"employee-module/repositories"
	"employee-module/services"
)

func main() {
	employeeRepo := repositories.NewEmployeeRepository()
	changeLogRepo := repositories.NewChangeLogRepository()
	employeeService := services.NewEmployeeService(employeeRepo, changeLogRepo)

	fmt.Println("Принимаем сотрудников на работу:")

	ivanov := employeeService.HireEmployee(1, "Иванов Иван Иванович", "+79001112233", "ivanov@company.ru", 1, 1)
	fmt.Printf("- %s принят в отдел %d на должность %d\n", ivanov.FullName, ivanov.DepartmentID, ivanov.PositionID)

	petrova := employeeService.HireEmployee(2, "Петрова Мария Сергеевна", "+79004445566", "petrova@company.ru", 1, 2)
	fmt.Printf("- %s принята в отдел %d на должность %d\n", petrova.FullName, petrova.DepartmentID, petrova.PositionID)

	fmt.Println("\nПереводим сотрудника в другое подразделение:")
	if err := employeeService.TransferEmployee(2, 3, 2); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("- %s переведена в отдел %d\n", petrova.FullName, petrova.DepartmentID)
	}

	fmt.Println("\nУвольняем сотрудника:")
	if err := employeeService.DismissEmployee(1); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("- %s уволен\n", ivanov.FullName)
	}

	fmt.Println("\nАктивные сотрудники:")
	active := employeeService.ListActiveEmployees()
	if len(active) == 0 {
		fmt.Println("  Нет активных сотрудников")
	} else {
		for _, employee := range active {
			fmt.Printf("  - %s (отдел %d, должность %d)\n", employee.FullName, employee.DepartmentID, employee.PositionID)
		}
	}

	fmt.Println("\nИстория изменений сотрудника Иванова:")
	for _, log := range employeeService.GetChangeHistory(1) {
		fmt.Printf("  - [%s] %s -> %s\n", log.Action, log.OldValue, log.NewValue)
	}

	fmt.Println("\nИстория изменений сотрудницы Петровой:")
	for _, log := range employeeService.GetChangeHistory(2) {
		fmt.Printf("  - [%s] %s -> %s\n", log.Action, log.OldValue, log.NewValue)
	}
}
