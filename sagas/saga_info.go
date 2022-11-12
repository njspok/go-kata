package sagas

func NewSagaInfo(id int) *SagaInfo {
	return &SagaInfo{
		id:        id,
		reserveId: 0,
		error:     nil,
	}
}

type SagaInfo struct {
	id        int
	reserveId int
	error     error
}

func (i *SagaInfo) ID() int {
	return i.id
}

func (i *SagaInfo) IsFail() bool {
	return !i.IsSuccess()
}

func (i *SagaInfo) IsSuccess() bool {
	return i.error == nil
}

func (i *SagaInfo) ReserveID() int {
	return i.reserveId
}

func (i *SagaInfo) SetReserveID(id int) {
	i.reserveId = id
}

func (i *SagaInfo) SetError(err error) {
	i.error = err
}

func (i *SagaInfo) Error() error {
	return i.error
}
