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

type Step func(*SagaInfo) error

func NewOrderSagaService(stock Stock, payment Payment) *SagaService {
	srv := &SagaService{
		list:    make(map[int]*SagaInfo),
		stock:   stock,
		payment: payment,
	}

	srv.scenario = []Step{
		srv.reserve,
		srv.pay,
	}

	return srv
}

type SagaService struct {
	list     map[int]*SagaInfo
	stock    Stock
	payment  Payment
	scenario []Step
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

	err := s.process(info)
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

	err := s.process(info)
	if err != nil {
		return err
	}

	return nil
}

func (s *SagaService) process(info *SagaInfo) error {
	for step := info.Step(); step < len(s.scenario); step++ {
		info.SetStep(step)
		action := s.scenario[step]
		err := action(info)
		if err != nil {
			return err
		}
	}

	info.Finish()
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
