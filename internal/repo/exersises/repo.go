package exersises

import "TelebotOne/internal/repo/models"

type Exerciser interface {
	// Создаёт таблицу для упражнения
	CreateExerciseTable(req *models.Exercise) error
	// Удаляет таблицу упражнения по её ID
	DropExerciseTable(ID uint16) error
	// Получает список из id - name всех таблиц упражнений
	GetAllExercisesIDAndName() ([]string, error)
}
