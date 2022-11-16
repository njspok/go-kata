package sagas

import "fmt"

func NewSagaInfo(order *Order) *SagaInfo {
	return &SagaInfo{
		id:        order.id,
		order:     order,
		reserveId: 0,
		log:       Log{},
	}
}

type SagaInfo struct {
	id        int
	order     *Order
	reserveId int
	payId     int
	log       Log
}

func (i *SagaInfo) ID() int {
	return i.id
}

func (i *SagaInfo) ReserveID() int {
	return i.reserveId
}

func (i *SagaInfo) SetReserveID(id int) {
	i.reserveId = id
}

func (i *SagaInfo) SetPayID(id int) {
	i.payId = id
}

func (i *SagaInfo) Log() Log {
	return i.log
}

func (i *SagaInfo) AddLog(s string, a ...any) {
	i.log = append(i.log, fmt.Sprintf(s, a...))
}

func (i *SagaInfo) PayID() int {
	return i.payId
}

type Log []string
