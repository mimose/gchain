package help

import (
	"errors"
	"fmt"
	"time"
)

const (
	STATUS_UNRIPE = iota
	STATUS_INIT
	STATUS_START
	STATUS_END
)

type StopWatch struct {
	start        time.Time
	duration     time.Duration
	status       int8
	points       []*point
	currentPoint *point
}

type point struct {
	name     string
	start    time.Time
	duration time.Duration
	status   int8
}

func NewStopWatch() StopWatch {
	return StopWatch{
		start:    time.Now(),
		duration: 0,
		status:   STATUS_INIT,
		points:   make([]*point, 0),
		currentPoint: &point{
			status: STATUS_UNRIPE,
		},
	}
}

func (sw *StopWatch) Start(name string) error {
	if len(name) == 0 {
		return errors.New("pls enter the name")
	}
	p := &point{
		name:     name,
		start:    time.Now(),
		duration: 0,
		status:   STATUS_START,
	}
	sw.points = append(sw.points, p)
	sw.currentPoint = p
	sw.status = STATUS_START
	return nil
}

func (sw *StopWatch) Stop() error {
	if *sw.currentPoint == (point{}) || sw.currentPoint.status != STATUS_START {
		return errors.New("no point in run")
	}
	stop := time.Now()
	sw.currentPoint.status = STATUS_END
	sw.currentPoint.duration = stop.Sub(sw.currentPoint.start)
	sw.status = STATUS_END
	sw.duration = stop.Sub(sw.start)
	return nil
}

func (sw StopWatch) Print() error {
	if sw.status != STATUS_END {
		return errors.New("pls stop first")
	}
	fmt.Println(sw.duration.String())
	for _, p := range sw.points {
		fmt.Println(p.name + ":" + p.duration.String())
	}
	return nil
}
