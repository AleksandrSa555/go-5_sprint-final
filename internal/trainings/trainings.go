package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	sliceData := strings.Split(datastring, ",")
	if len(sliceData) != 3 {
		return errors.New("неверное количество элементов")
	}
	steps, err := strconv.Atoi(sliceData[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше нуля")
	}
	t.Steps = steps

	t.TrainingType = sliceData[1]

	durationTraining, err := time.ParseDuration(sliceData[2])
	if err != nil {
		return err
	}
	if durationTraining <= 0 {
		return errors.New("продолжительность должна быть больше нуля")
	}
	t.Duration = durationTraining

	return nil
}

func (t Training) ActionInfo() (string, error) {

	dist := spentenergy.Distance(t.Steps, t.Personal.Height)
	avgSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	var calories float64
	var err error
	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}

	res := fmt.Sprintf("Тип тренировки: %s\n"+
		"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), dist, avgSpeed, calories)
	return res, nil
}
