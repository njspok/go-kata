package sagas

func NewOrderSaga(order *Order, scenario Scenario) *OrderSaga {
	return &OrderSaga{
		Saga: NewSaga(order.id, scenario),

		order:     order,
		reserveId: 0,
		payId:     0,
	}
}

// OrderSaga todo generalize saga abstract class?
type OrderSaga struct {
	*Saga

	order     *Order
	reserveId int
	payId     int
}

func (i *OrderSaga) Run() error {
	if i.IsFinished() {
		return ErrSagaFinished
	}

	return i.scenario.Run(i)
}

func (i *OrderSaga) TryAgain() error {
	if i.IsFinished() {
		return ErrSagaFinished
	}

	return i.Run()
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

func (i *OrderSaga) PayID() int {
	return i.payId
}
