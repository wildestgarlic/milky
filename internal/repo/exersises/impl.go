package exersises

import (
	"TelebotOne/internal/db"
	"TelebotOne/internal/repo/models"
	"fmt"
	"gorm.io/gorm"
)

type exercisesRepo struct {
	conn *gorm.DB
}

func NewExercisesRepo() Exerciser {
	r := &exercisesRepo{
		conn: db.New().GetConnection(),
	}

	return r
}

func (r *exercisesRepo) CreateExerciseTable(name string) error {
	exercise := &models.Exercise{
		Name: name,
	}

	err := r.conn.AutoMigrate(exercise)
	if err != nil {
		return err
	}

	return nil
}

func (r *exercisesRepo) FillExerciseTable(req *models.Exercise) error {
	exercise := &models.Exercise{
		Name:          req.Name,
		SetsQuantity:  req.SetsQuantity,
		RepsInEachSet: req.RepsInEachSet,
		Description:   req.Description,
		MuscleGroup:   req.MuscleGroup,
	}

	err := r.conn.Create(exercise).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *exercisesRepo) DropExerciseTable(ID uint16) error {
	exercise := &models.Exercise{
		ID: ID,
	}

	err := r.conn.Migrator().DropTable(exercise)
	if err != nil {
		return err
	}

	return nil
}

func (r *exercisesRepo) GetAllExercisesIDAndName() ([]string, error) {
	var response []string
	exercises := &[]models.Exercise{}

	err := r.conn.Select("id", "name").Find(exercises).Error
	if err != nil {
		return nil, err
	}

	for _, exercise := range *exercises {
		idAndName := fmt.Sprintf("ID:%d - %s", exercise.ID, exercise.Name)
		response = append(response, idAndName)
	}

	return response, nil
}
