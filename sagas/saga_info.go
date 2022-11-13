package sagas

func NewSagaInfo(id int) *SagaInfo {
	return &SagaInfo{
		id:        id,
		reserveId: 0,
		stages:    Stages{},
	}
}

type SagaInfo struct {
	id        int
	reserveId int
	stages    Stages
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

func (i *SagaInfo) Stages() Stages {
	return i.stages
}

type Stages []*Stage

type Stage struct {
	Name   string
	Status string
	Error  error
}
