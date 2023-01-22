package by_mock

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAuditManager(t *testing.T) {
	t.Run("success new record with clean history", func(t *testing.T) {
		now := time.Now()

		rep := &MockAuditRep{}
		rep.EXPECT().
			LastRecordByName("ivan").
			Return(nil, nil).
			Once()
		rep.EXPECT().
			Add("ivan", now, In).
			Return(nil).
			Once()

		manager := NewAuditManager(rep)
		err := manager.AddRecord("ivan", now, In)
		require.NoError(t, err)
	})
	t.Run("fail get last record", func(t *testing.T) {
		now := time.Now()

		rep := &MockAuditRep{}
		rep.EXPECT().
			LastRecordByName("ivan").
			Return(nil, errors.New("shit happens")).
			Once()

		manager := NewAuditManager(rep)
		err := manager.AddRecord("ivan", now, In)
		require.EqualError(t, err, "shit happens")

	})
	t.Run("fail new record with clean history", func(t *testing.T) {
		rep := &MockAuditRep{}
		rep.EXPECT().
			LastRecordByName("ivan").
			Return(nil, nil).
			Once()

		manager := NewAuditManager(rep)
		err := manager.AddRecord("ivan", time.Now(), Out)
		require.ErrorIs(t, err, ErrInvalidDirection)
	})
	t.Run("success new record with history", func(t *testing.T) {
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
		rep.EXPECT().
			Add("ivan", now, Out).
			Return(nil).
			Once()

		manager := NewAuditManager(rep)
		err := manager.AddRecord("ivan", now, Out)
		require.NoError(t, err)
	})
	t.Run("fail with name mismatch", func(t *testing.T) {
		now := time.Now()

		rep := &MockAuditRep{}
		rep.EXPECT().
			LastRecordByName("ivan").
			Return(&AuditRecord{
				Direction: In,
				Time:      now.Add(-time.Hour),
				Name:      "peter",
			}, nil).
			Once()

		manager := NewAuditManager(rep)
		err := manager.AddRecord("ivan", now, Out)
		require.ErrorIs(t, err, ErrNameMismatch)
	})
	t.Run("fail with invalid direction", func(t *testing.T) {
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

		manager := NewAuditManager(rep)
		err := manager.AddRecord("ivan", now, In)
		require.ErrorIs(t, err, ErrInvalidDirection)
	})
	t.Run("fail with invalid time", func(t *testing.T) {
		now := time.Now()

		rep := &MockAuditRep{}
		rep.EXPECT().
			LastRecordByName("ivan").
			Return(&AuditRecord{
				Direction: In,
				Time:      now,
				Name:      "ivan",
			}, nil).
			Once()

		manager := NewAuditManager(rep)
		err := manager.AddRecord("ivan", now.Add(-time.Hour), In)
		require.ErrorIs(t, err, ErrInvalidTime)
	})
}
