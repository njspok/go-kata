package sagas

func NewOrderSaga(order *Order, stock Stock, payment Payment) *OrderSaga {
	orderSaga := &OrderSaga{
		order:     order,
		reserveId: 0,
		payId:     0,

		stock:   stock,
		payment: payment,
	}

	orderSaga.Saga = NewSaga(order.id, Scenario{
		{
			name:   "Reserve",
			action: orderSaga.reserve,
		},
		{
			name:   "Pay",
			action: orderSaga.pay,
		},
	})

	return orderSaga
}

// OrderSaga todo generalize saga abstract class?
type OrderSaga struct {
	*Saga

	order     *Order
	reserveId int
	payId     int

	stock   Stock
	payment Payment
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

func (i *OrderSaga) reserve(saga *OrderSaga) error {
	reserveId, err := i.stock.Reserve(saga.order.itemId, saga.order.qty)
	if err != nil {
		return err
	}

	// todo move to saga?
	saga.SetReserveID(reserveId)
	return nil
}

func (s *OrderSaga) pay(saga *OrderSaga) error {
	payId, err := s.payment.Pay(saga.order.clientId, saga.order.sum)
	if err != nil {
		return err
	}

	// todo move to saga?
	saga.SetPayID(payId)
	return nil
}
