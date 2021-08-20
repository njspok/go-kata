package transaction

type OperationI interface {
	Execute() error
	Undo() error
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

type Transaction struct {
	chain   []*element
	undoErr error
	execErr error
}

func (t *Transaction) Empty() bool {
	return len(t.chain) == 0
}

func (t *Transaction) Execute() bool {
	if t.Empty() {
		return true
	}

	execErr, undoErr := t.firstElement().Execute()
	t.undoErr = undoErr
	t.execErr = execErr
	return t.undoErr == nil && t.execErr == nil
}

func (t *Transaction) Add(op OperationI) {
	newElement := &element{
		operation: op,
		next:      nil,
	}

	// link last element
	if !t.Empty() {
		t.lastElement().next = newElement
	}

	t.chain = append(t.chain, newElement)
}

func (t *Transaction) UndoErr() error {
	return t.undoErr
}

func (t *Transaction) ExecErr() error {
	return t.execErr
}

func (t *Transaction) lastElement() *element {
	return t.chain[len(t.chain)-1]
}

func (t *Transaction) firstElement() *element {
	return t.chain[0]
}

type element struct {
	operation OperationI
	next      *element
}

func (e *element) Execute() (error, error) {
	err := e.operation.Execute()
	if err != nil {
		return err, e.operation.Undo()
	}

	if e.next != nil {
		execErr, undoErr := e.next.Execute()

		if undoErr != nil {
			return execErr, undoErr
		}

		if execErr != nil {
			return execErr, e.operation.Undo()
		}

		return nil, nil
	}

	return nil, nil
}
