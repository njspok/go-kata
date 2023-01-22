package mock_version

import "time"

type Auditer interface {
	AddRecord(
		lastRec *AuditRecord,
		name string,
		time time.Time,
		direction Direction,
	) (
		*AuditRecord,
		error,
	)
}

type AuditRep interface {
	LastRecordByName(name string) (*AuditRecord, error)
	Add(record *AuditRecord) error
}

func NewAuditService(rep AuditRep, manager Auditer) *AuditService {
	return &AuditService{
		manager: manager,
		rep:     rep,
	}
}

type AuditService struct {
	manager Auditer
	rep     AuditRep
}

func (s *AuditService) AddRecord(name string, time time.Time, direction string) error {
	rec, err := s.rep.LastRecordByName(name)
	if err != nil {
		return err
	}

	d, err := NewDirectionFromString(direction)
	if err != nil {
		return err
	}

	newRec, err := s.manager.AddRecord(rec, name, time, d)
	if err != nil {
		return err
	}

	return s.rep.Add(newRec)
}
