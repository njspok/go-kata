package sagas

import "errors"

var (
	ErrOrderAlreadyProcessed = errors.New("order already processed")
)

type Stock interface {
	Reserve(itemId int, qty int) (int, error)
	CancelReserve(id int) error
}

func NewOrderSagaService(stock Stock) *SagaService {
	return &SagaService{
		list:  make(map[int]*SagaInfo),
		stock: stock,
	}
}

type SagaService struct {
	list  map[int]*SagaInfo
	stock Stock
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

	s.list[order.id] = info

	reserveStage := &Stage{
		Name:   "Reserve",
		Status: "Process",
		Error:  nil,
	}
	info.stages = append(info.stages, reserveStage)

	reserveId, err := s.stock.Reserve(order.itemId, order.qty)
	if err != nil {
		reserveStage.Status = "Fail"
		reserveStage.Error = err
		return 0, err
	}

	reserveStage.Status = "Success"
	info.SetReserveID(reserveId)

	return order.id, nil
}
