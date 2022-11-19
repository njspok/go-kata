package sagas

import (
	"errors"
	"fmt"
)

var (
	ErrSagaFinished = errors.New("saga is finished")
)

type Log []string

func NewSaga(id int, scenario Scenario) *Saga {
	return &Saga{
		id:         id,
		log:        Log{},
		isFinished: false,
		stepN:      0,
		scenario:   scenario,
	}
}

type Saga struct {
	id         int
	log        Log
	isFinished bool
	stepN      int
	scenario   Scenario
}

func (s *Saga) ID() int {
	return s.id
}

func (s *Saga) Run() error {
	if s.IsFinished() {
		return ErrSagaFinished
	}

	return s.scenario.Run(s)
}

func (s *Saga) TryAgain() error {
	if s.IsFinished() {
		return ErrSagaFinished
	}

	return s.Run()
}

func (s *Saga) Log() Log {
	return s.log
}

func (s *Saga) AddLog(format string, a ...any) {
	s.log = append(s.log, fmt.Sprintf(format, a...))
}

func (s *Saga) Finish() {
	s.isFinished = true
}

func (s *Saga) IsFinished() bool {
	return s.isFinished
}

func (s *Saga) SetStepN(step int) {
	s.stepN = step
}

func (s *Saga) StepN() int {
	return s.stepN
}
