package daysteps

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
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
		return err
	}
	ds.Steps = steps

	durationSteps, err := time.ParseDuration(sliceData[1])
	if err != nil {
		return err
	}
	ds.Duration = durationSteps

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
}
