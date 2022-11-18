package sagas

import "fmt"

func NewOrderSaga(order *Order, scenario Scenario) *OrderSaga {
	return &OrderSaga{
		id:        order.id,
		order:     order,
		reserveId: 0,
		log:       Log{},
		scenario:  scenario,
	}
}

type OrderSaga struct {
	id         int
	order      *Order
	reserveId  int
	payId      int
	log        Log
	stepN      int
	isFinished bool
	scenario   Scenario
}

func (i *OrderSaga) Run() error {
	return i.scenario.Run(i)
}

func (i *OrderSaga) TryAgain() error {
	if i.IsFinished() {
		return ErrSagaFinished
	}

	return i.Run()
}

func (i *OrderSaga) ID() int {
	return i.id
}

func (i *OrderSaga) ReserveID() int {
	return i.reserveId
}

func (i *OrderSaga) SetReserveID(id int) {
	i.reserveId = id
}

func (i *OrderSaga) SetPayID(id int) {
	i.payId = id
}

func (i *OrderSaga) Log() Log {
	return i.log
}

func (i *OrderSaga) AddLog(s string, a ...any) {
	i.log = append(i.log, fmt.Sprintf(s, a...))
}

func (i *OrderSaga) PayID() int {
	return i.payId
}

func (i *OrderSaga) SetStepN(step int) {
	i.stepN = step
}

func (i *OrderSaga) StepN() int {
	return i.stepN
}

func (i *OrderSaga) Finish() {
	i.isFinished = true
}

func (i *OrderSaga) IsFinished() bool {
	return i.isFinished
}

type Log []string
