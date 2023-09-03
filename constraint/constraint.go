package constraint

import (
	"errors"

	"github.com/samber/lo"
)

type Result int

const (
	Correct Result = iota
	Incorrect
	Missing
)

type Solution[V comparable, D any] map[V]D

type Constraint[V comparable, D any] interface {
	Satisfied(Solution[V, D]) Result
	Variables() []V
}

func NewCSP[V comparable, D any](variables []V, domains map[V][]D) (*CSP[V, D], error) {
	csp := &CSP[V, D]{
		variables: variables,
		domains:   domains,
	}

	csp.constraints = make(map[V][]Constraint[V, D])
	for _, v := range csp.variables {
		csp.constraints[v] = []Constraint[V, D]{}
		if _, ok := csp.domains[v]; !ok {
			return nil, errors.New("every variable should have a domain assigned to it")
		}
	}

	return csp, nil
}

type CSP[V comparable, D any] struct {
	variables   []V
	domains     map[V][]D
	constraints map[V][]Constraint[V, D]
}

func (c *CSP[V, D]) AddConstraint(constr Constraint[V, D]) error {
	for _, v := range constr.Variables() {
		if lo.Contains(c.variables, v) {
			c.constraints[v] = append(c.constraints[v], constr)
		} else {
			return errors.New("variable in constraint not in CSP")
		}
	}

	return nil
}

func (c *CSP[V, D]) AddConstraints(list ...Constraint[V, D]) error {
	for _, constr := range list {
		if err := c.AddConstraint(constr); err != nil {
			return err
		}
	}
	return nil
}

func (c *CSP[V, D]) Consistent(v V, solution Solution[V, D]) bool {
	for _, constr := range c.constraints[v] {
		switch constr.Satisfied(solution) {
		case Missing:
			continue
		case Correct:
			continue
		case Incorrect:
			return false
		}
	}
	return true
}

func (c *CSP[V, D]) Search() Solution[V, D] {
	return c.backtrackingSearch(nil)
}

//func (c *CSP[V, D]) backtrackingSearch(solution Solution[V, D]) Solution[V, D] {
//	if solution == nil {
//		solution = make(Solution[V, D])
//	}
//
//	// find all assignments for variables
//	if len(solution) == len(c.variables) {
//		return solution
//	}
//
//	var unassigned []V
//	for _, v := range c.variables {
//		if _, ok := solution[v]; !ok {
//			unassigned = append(unassigned, v)
//		}
//	}
//
//	first := unassigned[0]
//	for _, value := range c.domains[first] {
//		localSolution := copyMap(solution)
//		localSolution[first] = value
//		if c.Consistent(first, localSolution) {
//			result := c.backtrackingSearch(localSolution)
//			if result != nil {
//				return result
//			}
//		}
//	}
//
//	return nil
//}

func (c *CSP[V, D]) backtrackingSearch(solution Solution[V, D]) Solution[V, D] {
	if solution == nil {
		solution = make(Solution[V, D])
	}

	// find all assignments for variables
	if len(solution) == len(c.variables) {
		return solution
	}

	var unassigned []V
	for _, v := range c.variables {
		if _, ok := solution[v]; !ok {
			unassigned = append(unassigned, v)
		}
	}

	first := unassigned[0]
	for _, value := range c.domains[first] {
		localSolution := copyMap(solution)
		localSolution[first] = value
		if c.Consistent(first, localSolution) {
			result := c.backtrackingSearch(localSolution)
			if result != nil {
				return result
			}
		}
	}

	return nil
}

func copyMap[V comparable, D any](m map[V]D) map[V]D {
	result := make(map[V]D, len(m))
	for v, d := range m {
		result[v] = d
	}
	return result
}
