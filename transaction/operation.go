package transaction

type Action func() error

func NewOperation(execute Action, undo Action) *Operation {
	return &Operation{
		execute: execute,
		undo:    undo,
	}
}

type Operation struct {
	execute Action
	undo    Action
}

func (o *Operation) Execute() error {
	return o.execute()
}

func (o *Operation) Undo() error {
	if o.undo == nil {
		return nil
	}
	return o.undo()
}
