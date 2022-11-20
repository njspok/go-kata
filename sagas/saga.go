package sagas

import (
	"errors"
	"fmt"
)

var (
	ErrSagaFinished = errors.New("saga is finished")
)

type Log []string

type Action func() error

type Step struct {
	name   string
	action Action
}

func (s Step) Run() error {
	return s.action()
}

func (s Step) Rollback() error {
	// todo implement!!!
	panic("not implemented")
}

type Scenario []Step

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
	scenario   Scenario
	log        Log
	stepN      int
	isFinished bool
}

func (s *Saga) ID() int {
	return s.id
}

func (s *Saga) Run() error {
	if s.IsFinished() {
		return ErrSagaFinished
	}

	for n := s.StepN(); n < len(s.scenario); n++ {
		s.SetStepN(n)

		step := s.scenario[n]

		s.AddLog("%s Process", step.name)

		err := step.Run()
		if err != nil {
			s.AddLog("%s Fail: %v", step.name, err)
			return err
		}

		s.AddLog("%s Success", step.name)
	}

	s.Finish()

	return nil
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
