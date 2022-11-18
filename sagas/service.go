package sagas

import "errors"

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

type Action func(*SagaInfo) error

type Step struct {
	action Action
}

func (s Step) Run(info *SagaInfo) error {
	return s.action(info)
}

func (s Step) Rollback(info *SagaInfo) error {
	// todo implement!!!
	panic("not implemented")
}

type Scenario []Step

func (s Scenario) Run(info *SagaInfo) error {
	for n := info.StepN(); n < len(s); n++ {
		info.SetStepN(n)

		step := s[n]
		err := step.Run(info)
		if err != nil {
			return err
		}
	}

	info.Finish()
	return nil
}

func NewOrderSagaService(stock Stock, payment Payment) *SagaService {
	srv := &SagaService{
		list:    make(map[int]*SagaInfo),
		stock:   stock,
		payment: payment,
	}

	srv.scenario = Scenario{
		{
			action: srv.reserve,
		},
		{
			action: srv.pay,
		},
	}

	return srv
}

type SagaService struct {
	list     map[int]*SagaInfo
	stock    Stock
	payment  Payment
	scenario Scenario
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

	info := NewSagaInfo(order)

	// todo saga start

	s.list[order.id] = info

	err := s.scenario.Run(info)
	if err != nil {
		return info.id, err
	}

	// todo saga success

	return info.id, nil
}

func (s *SagaService) Rollback(sagaId int) error {
	// todo implement!!!
	panic("not implemented")
}

func (s *SagaService) TryAgain(sagaId int) error {
	info := s.list[sagaId]
	if info == nil {
		return ErrSagaNotFound
	}

	if info.IsFinished() {
		return ErrSagaFinished
	}

	err := s.scenario.Run(info)
	if err != nil {
		return err
	}

	return nil
}

func (s *SagaService) pay(info *SagaInfo) error {
	info.AddLog("Pay Process")
	payId, err := s.payment.Pay(info.order.clientId, info.order.sum)
	if err != nil {
		info.AddLog("Pay Fail: %v", err)
		return err
	}

	info.SetPayID(payId)
	info.AddLog("Pay Success")
	return nil
}

func (s *SagaService) reserve(info *SagaInfo) error {
	info.AddLog("Reserve Process")
	reserveId, err := s.stock.Reserve(info.order.itemId, info.order.qty)
	if err != nil {
		info.AddLog("Reserve Fail: %v", err)
		return err
	}

	info.SetReserveID(reserveId)
	info.AddLog("Reserve Success")
	return nil
}
