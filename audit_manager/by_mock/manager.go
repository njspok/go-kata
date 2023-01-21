package by_mock

import (
	"time"

	"github.com/pkg/errors"
)

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
	ID        int
	Name      string
	Time      time.Time
	Direction Direction
}

type AuditRep interface {
	LastRecordByName(name string) (*AuditRecord, error)
	Add(name string, time time.Time, direction Direction) error
}

func NewAuditManager(rep AuditRep) *AuditManager {
	return &AuditManager{
		auditRep: rep,
	}
}

type AuditManager struct {
	auditRep AuditRep
}

func (am *AuditManager) AddRecord(name string, time time.Time, direction Direction) error {
	rec, err := am.auditRep.LastRecordByName(name)
	if err != nil {
		return err
	}

	if rec == nil {
		if direction == In {
			return am.auditRep.Add(name, time, direction)
		}
		return ErrInvalidDirection
	}

	if rec.Time.After(time) {
		return ErrInvalidTime
	}

	if rec.Name != name {
		return ErrNameMismatch
	}

	if rec.Direction == direction {
		return ErrInvalidDirection
	}

	return am.auditRep.Add(name, time, direction)
}
