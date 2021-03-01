package lamp

func NewWorkLamp() *WorkLamp {
	return &WorkLamp{
		light: lightOff,
	}
}

type lightState string

const (
	lightOn  lightState = "on"
	lightOff lightState = "off"
)

type WorkLamp struct {
	light lightState
}

func (l *WorkLamp) IsLighted() bool {
	return l.light == lightOn
}

func (l *WorkLamp) Switch() {
	if l.light == lightOff {
		l.light = lightOn
	} else {
		l.light = lightOff
	}
}
