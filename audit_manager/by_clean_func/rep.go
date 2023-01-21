package mock_version

type AuditRep interface {
	LastRecordByName(name string) (*AuditRecord, error)
	Add(record *AuditRecord) error
}
