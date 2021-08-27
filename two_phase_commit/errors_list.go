package two_phase_commit

import "fmt"

type ErrorsList []error

func (l *ErrorsList) Add(err error) {
	*l = append(*l, err)
}

func (l *ErrorsList) Empty() bool {
	return len(*l) == 0
}

func (l *ErrorsList) Count() int {
	return len(*l)
}

func (l ErrorsList) Error() string {
	if l.Empty() {
		return "list without errors"
	}
	return fmt.Sprintf("list of %v errors", l.Count())
}
