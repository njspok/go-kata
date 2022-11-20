package sagas

import (
	"errors"
)

var (
	ErrOrderAlreadyProcessed = errors.New("order already processed")
	ErrSagaNotFound          = errors.New("saga not found")
)

type Stock interface {
	Reserve(itemId int, qty int) (int, error)
	CancelReserve(id int) error
}

type Payment interface {
	Pay(clientId int, sum int) (int, error)
	CancelPay(id int) error
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
	list    map[int]*OrderSaga
	stock   Stock
	payment Payment
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

	s.list[saga.ID()] = saga

	err := saga.Run()
	if err != nil {
		return saga.ID(), err
	}

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

	err := saga.Run()
	if err != nil {
		return err
	}

	return nil
}
