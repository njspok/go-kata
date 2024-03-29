package varops

import "errors"

var ErrSolveNotFound = errors.New("solve not found")

type VarName string
type VarValue int
type Operation func(Solver) VarValue

func NewSolver() Solver {
	return make(Solver)
}

type Solver map[VarName]func(Solver) VarValue

func (l Solver) Solve(varName VarName) (VarValue, error) {
	op, found := l[varName]
	if found {
		return op(l), nil
	}

	return 0, ErrSolveNotFound
}

func (l Solver) SetOperation(varName VarName, op Operation) {
	l[varName] = op
}

func (l Solver) SetValue(varName VarName, value VarValue) {
	l.SetOperation(varName, func(_ Solver) VarValue { return value })
}
