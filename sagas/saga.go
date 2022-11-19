package sagas

import "fmt"

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

func (i *Saga) ID() int {
	return i.id
}

func (s *Saga) Run() error {
	if s.IsFinished() {
		return ErrSagaFinished
	}

	return s.scenario.Run(s)
}

func (s *OrderSaga) TryAgain() error {
	if s.IsFinished() {
		return ErrSagaFinished
	}

	return s.Run()
}

func (i *Saga) Log() Log {
	return i.log
}

func (i *Saga) AddLog(s string, a ...any) {
	i.log = append(i.log, fmt.Sprintf(s, a...))
}

func (i *Saga) Finish() {
	i.isFinished = true
}

func (i *Saga) IsFinished() bool {
	return i.isFinished
}

func (i *Saga) SetStepN(step int) {
	i.stepN = step
}

func (i *Saga) StepN() int {
	return i.stepN
}
