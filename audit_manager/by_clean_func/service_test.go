package mock_version

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAuditService(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		now := time.Now()

		rep := &MockAuditRep{}
		rep.EXPECT().
			LastRecordByName("ivan").
			Return(nil, nil).
			Once()
		rep.EXPECT().
			Add(&AuditRecord{
				Direction: In,
				Time:      now,
				Name:      "ivan",
			}).
			Return(nil).
			Once()

		manager := NewAuditManager()

		srv := NewAuditService(rep, manager)
		err := srv.AddRecord("ivan", now, "in")
		require.NoError(t, err)
	})
	t.Run("fail get last record", func(t *testing.T) {
		now := time.Now()

		rep := &MockAuditRep{}
		rep.EXPECT().
			LastRecordByName("ivan").
			Return(nil, errors.New("shit happens")).
			Once()

		manager := NewAuditManager()

		srv := NewAuditService(rep, manager)
		err := srv.AddRecord("ivan", now, "in")
		require.EqualError(t, err, "shit happens")
	})
	t.Run("invalid direction", func(t *testing.T) {
		now := time.Now()

		rep := &MockAuditRep{}
		rep.EXPECT().
			LastRecordByName("ivan").
			Return(nil, nil).
			Once()

		manager := NewAuditManager()

		srv := NewAuditService(rep, manager)
		err := srv.AddRecord("ivan", now, "incoming")
		require.EqualError(t, err, "unknown direction")
	})
	t.Run("cant append record", func(t *testing.T) {
		now := time.Now()

		rep := &MockAuditRep{}
		rep.EXPECT().
			LastRecordByName("ivan").
			Return(&AuditRecord{
				Direction: In,
				Time:      now.Add(-time.Hour),
				Name:      "ivan",
			}, nil).
			Once()

		manager := NewAuditManager()

		srv := NewAuditService(rep, manager)
		err := srv.AddRecord("ivan", now, "in")
		require.EqualError(t, err, "invalid direction")
	})
	t.Run("fail when add record in rep", func(t *testing.T) {
		now := time.Now()

		rep := &MockAuditRep{}
		rep.EXPECT().
			LastRecordByName("ivan").
			Return(nil, nil).
			Once()
		rep.EXPECT().
			Add(&AuditRecord{
				Direction: In,
				Time:      now,
				Name:      "ivan",
			}).
			Return(errors.New("shit happens")).
			Once()

		manager := NewAuditManager()

		srv := NewAuditService(rep, manager)
		err := srv.AddRecord("ivan", now, "in")
		require.EqualError(t, err, "shit happens")
	})
}
