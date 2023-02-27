package main

import (
	"fmt"
	"time"
)

type ExecTimer struct {
	name        string
	startTime   time.Time
	elapsedTime time.Duration
}

func NewExecTimer(name string) *ExecTimer {
	et := new(ExecTimer)

	et.name = name

	return et
}

func (et *ExecTimer) Start() {
	et.startTime = time.Now()
}

// TODO : fix inconsistent elapsed time calculation
func (et *ExecTimer) Finish() {
	et.elapsedTime = time.Since(et.startTime)
}

func (et *ExecTimer) Tell() {
	fmt.Println(et.name, "took", et.elapsedTime)
}
