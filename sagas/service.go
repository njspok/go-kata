package sagas

import (
	"errors"
)

var (
	ErrOrderAlreadyProcessed = errors.New("order already processed")
	ErrSagaNotFound          = errors.New("saga not found")
	ErrSagaFinished          = errors.New("saga is finished")
)

type Stock interface {
	Reserve(itemId int, qty int) (int, error)
	CancelReserve(id int) error
}

type Payment interface {
	Pay(clientId int, sum int) (int, error)
	CancelPay(id int) error
}

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

func (s Scenario) Run(saga *Saga) error {
	for n := saga.StepN(); n < len(s); n++ {
		saga.SetStepN(n)

		step := s[n]

		saga.AddLog("%s Process", step.name)

		err := step.Run()
		if err != nil {
			saga.AddLog("%s Fail: %v", step.name, err)
			return err
		}

		saga.AddLog("%s Success", step.name)
	}

	saga.Finish()
	return nil
}

func NewOrderSagaService(stock Stock, payment Payment) *SagaService {
	srv := &SagaService{
		list:    make(map[int]*OrderSaga),
		stock:   stock,
		payment: payment,
	}
	return srv
}

type SagaService struct {
	list     map[int]*OrderSaga
	stock    Stock
	payment  Payment
	scenario Scenario
}

func (s *SagaService) OrderSaga(id int) *OrderSaga {
	if saga, ok := s.list[id]; ok {
		return saga
	}
	return nil
}

func (s *SagaService) Run(order *Order) (int, error) {
	if s.list[order.id] != nil {
		return 0, ErrOrderAlreadyProcessed
	}

	saga := NewOrderSaga(order, s.stock, s.payment)

	// todo saga start

	s.list[saga.ID()] = saga

	err := saga.Run()
	if err != nil {
		return saga.ID(), err
	}

	// todo saga success

	return saga.ID(), nil
}

func (s *SagaService) Rollback(sagaId int) error {
	// todo implement!!!
	panic("not implemented")
}

func (s *SagaService) TryAgain(sagaId int) error {
	saga := s.list[sagaId]
	if saga == nil {
		return ErrSagaNotFound
	}

	err := saga.TryAgain()
	if err != nil {
		return err
	}

	return nil
}
