package lamp

import "errors"

var ErrLampNotWorking = errors.New("lamp not working")

func NewMan(patience uint) *Man {
	return &Man{
		patience: patience,
	}
}

type lamper interface {
	IsLighted() bool
	Switch()
}

type Man struct {
	patience uint
}

func (m *Man) LightOn(lamp lamper) error {
	if lamp.IsLighted() {
		return nil
	}

	if m.tryLightOn(lamp) {
		return nil
	}

	if m.impatient() {
		return ErrLampNotWorking
	}

	for i := uint(0); i < m.patience; i++ {
		if m.tryLightOn(lamp) {
			return nil
		}
	}

	return ErrLampNotWorking
}

func (m *Man) impatient() bool {
	return m.patience == 0
}

func (m *Man) tryLightOn(lamp lamper) bool {
	lamp.Switch()
	return lamp.IsLighted()
}
