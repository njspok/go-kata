package sagas

import "fmt"

func NewSagaInfo(id int) *SagaInfo {
	return &SagaInfo{
		id:        id,
		reserveId: 0,
		log:       Log{},
	}
}

type SagaInfo struct {
	id        int
	reserveId int
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

func (i *SagaInfo) Log() Log {
	return i.log
}

func (i *SagaInfo) AddLog(s string, a ...any) {
	i.log = append(i.log, fmt.Sprintf(s, a...))
}

type Log []string
