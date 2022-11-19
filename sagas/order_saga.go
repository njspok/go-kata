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

type OrderSaga struct {
	*Saga

	order     *Order
	reserveId int
	payId     int

	stock   Stock
	payment Payment
}

func (s *OrderSaga) ReserveID() int {
	return s.reserveId
}

func (s *OrderSaga) SetReserveID(id int) {
	s.reserveId = id
}

func (s *OrderSaga) SetPayID(id int) {
	s.payId = id
}

func (s *OrderSaga) PayID() int {
	return s.payId
}

func (s *OrderSaga) reserve() error {
	reserveId, err := s.stock.Reserve(s.order.itemId, s.order.qty)
	if err != nil {
		return err
	}

	s.SetReserveID(reserveId)
	return nil
}

func (s *OrderSaga) pay() error {
	payId, err := s.payment.Pay(s.order.clientId, s.order.sum)
	if err != nil {
		return err
	}

	s.SetPayID(payId)
	return nil
}
