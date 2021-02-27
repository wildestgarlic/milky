package exersises

import "TelebotOne/internal/repo/models"

type Exerciser interface {
	CreateExerciseTable(name string) error
	// Заполняет таблицу для упражнения
	FillExerciseTable(req *models.Exercise) error //fixme: пункт для индентификации
	// Удаляет таблицу упражнения по её ID
	DropExerciseTable(ID uint16) error
	// Получает список из id - name всех таблиц упражнений
	GetAllExercisesIDAndName() ([]string, error)
}
