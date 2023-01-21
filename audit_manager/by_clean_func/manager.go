package mock_version

import (
	"time"

	"github.com/pkg/errors"
)

func NewDirectionFromString(s string) (Direction, error) {
	if s == string(In) {
		return In, nil
	}
	if s == string(Out) {
		return Out, nil
	}

	return "", errors.New("unknown direction")
}

type Direction string

const (
	In  Direction = "in"
	Out Direction = "out"
)

var (
	ErrInvalidDirection = errors.New("invalid direction")
	ErrNameMismatch     = errors.New("name mismatch")
	ErrInvalidTime      = errors.New("invalid time")
)

type AuditRecord struct {
	Name      string
	Time      time.Time
	Direction Direction
}

func NewAuditManager() *AuditManager {
	return &AuditManager{}
}

type AuditManager struct {
}

func (am *AuditManager) AddRecord(
	lastRec *AuditRecord,
	name string,
	time time.Time,
	direction Direction,
) (
	*AuditRecord,
	error,
) {
	if lastRec == nil {
		if direction == In {
			return &AuditRecord{
				Direction: direction,
				Time:      time,
				Name:      name,
			}, nil
		}
		return nil, ErrInvalidDirection
	}

	if lastRec.Time.After(time) {
		return nil, ErrInvalidTime
	}

	if lastRec.Name != name {
		return nil, ErrNameMismatch
	}

	if lastRec.Direction == direction {
		return nil, ErrInvalidDirection
	}

	return &AuditRecord{
		Direction: direction,
		Time:      time,
		Name:      name,
	}, nil
}
