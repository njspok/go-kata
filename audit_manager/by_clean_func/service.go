package mock_version

import "time"

type AuditService struct {
	manager *AuditManager
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
