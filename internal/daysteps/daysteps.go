package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	sliceData := strings.Split(datastring, ",")
	if len(sliceData) != 2 {
		return errors.New("неверное количество элементов")
	}
	steps, err := strconv.Atoi(sliceData[0])
	if err != nil {
		return fmt.Errorf("неправильный формат шагов: %w", err)
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше нуля")
	}
	ds.Steps = steps

	durationSteps, err := time.ParseDuration(sliceData[1])
	if err != nil {
		return fmt.Errorf("неверный формат продолжительности шагов: %w", err)
	}
	if durationSteps <= 0 {
		return errors.New("продолжительность должна быть больше нуля")
	}
	ds.Duration = durationSteps

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	dist := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("неверный расчет калорий: %w", err)
	}
	res := fmt.Sprintf("Количество шагов: %d.\n"+
		"Дистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.\n",
		ds.Steps, dist, calories)
	return res, nil
}
