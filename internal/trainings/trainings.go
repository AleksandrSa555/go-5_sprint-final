package trainings

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
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
	t.Steps = steps

	t.TrainingType = sliceData[1]

	durationTraining, err := time.ParseDuration(sliceData[2])
	if err != nil {
		return err
	}
	t.Duration = durationTraining

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
}
