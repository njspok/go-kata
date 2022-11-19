package sagas

type Action func() error

type Step struct {
	name   string
	action Action
}

func (s Step) Run() error {
	return s.action()
}

func (s Step) Rollback() error {
	// todo implement!!!
	panic("not implemented")
}

type Scenario []Step

func (s Scenario) Run(saga *Saga) error {
	for n := saga.StepN(); n < len(s); n++ {
		saga.SetStepN(n)

		step := s[n]

		saga.AddLog("%s Process", step.name)

		err := step.Run()
		if err != nil {
			saga.AddLog("%s Fail: %v", step.name, err)
			return err
		}

		saga.AddLog("%s Success", step.name)
	}

	saga.Finish()
	return nil
}
