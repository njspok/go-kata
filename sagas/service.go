package sagas

import "errors"

var (
	ErrOrderAlreadyProcessed = errors.New("order already processed")
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
	return &SagaService{
		list:    make(map[int]*SagaInfo),
		stock:   stock,
		payment: payment,
	}
}

type SagaService struct {
	list    map[int]*SagaInfo
	stock   Stock
	payment Payment
}

func (s *SagaService) SagaInfo(id int) *SagaInfo {
	if saga, ok := s.list[id]; ok {
		return saga
	}
	return nil
}

func (s *SagaService) Run(order *Order) (int, error) {
	if s.list[order.id] != nil {
		return 0, ErrOrderAlreadyProcessed
	}

	info := NewSagaInfo(order.id)

	// todo saga start

	s.list[order.id] = info

	info.AddLog("Reserve Process")
	reserveId, err := s.stock.Reserve(order.itemId, order.qty)
	if err != nil {
		info.AddLog("Reserve Fail: %v", err)
		return 0, err
	}
	info.SetReserveID(reserveId)
	info.AddLog("Reserve Success")

	info.AddLog("Pay Process")
	payId, err := s.payment.Pay(order.clientId, order.sum)
	if err != nil {
		info.AddLog("Pay Fail: %v", err)
		return 0, err
	}
	info.SetPayID(payId)
	info.AddLog("Pay Success")

	// todo saga success

	return order.id, nil
}
