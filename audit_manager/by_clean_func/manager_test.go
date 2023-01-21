package mock_version

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAuditManager(t *testing.T) {
	t.Run("success new record with clean history", func(t *testing.T) {
		manager := NewAuditManager()

		now := time.Now()

		newRec, err := manager.AddRecord(nil, "ivan", now, In)
		require.NoError(t, err)
		require.Equal(t, &AuditRecord{
			Direction: In,
			Time:      now,
			Name:      "ivan",
		}, newRec)
	})
	t.Run("fail new record with clean history", func(t *testing.T) {
		manager := NewAuditManager()
		newRec, err := manager.AddRecord(nil, "ivan", time.Now(), Out)
		require.ErrorIs(t, err, ErrInvalidDirection)
		require.Nil(t, newRec)
	})
	t.Run("success new record with history", func(t *testing.T) {
		manager := NewAuditManager()

		now := time.Now()

		lastRec := &AuditRecord{
			Direction: In,
			Time:      now.Add(-time.Hour),
			Name:      "ivan",
		}

		newRec, err := manager.AddRecord(lastRec, "ivan", now, Out)
		require.NoError(t, err)
		require.Equal(t, &AuditRecord{
			Direction: Out,
			Time:      now,
			Name:      "ivan",
		}, newRec)
	})
	t.Run("fail with name mismatch", func(t *testing.T) {
		manager := NewAuditManager()

		now := time.Now()

		lastRec := &AuditRecord{
			Direction: In,
			Time:      now.Add(-time.Hour),
			Name:      "peter",
		}

		newRec, err := manager.AddRecord(lastRec, "ivan", now, Out)
		require.ErrorIs(t, err, ErrNameMismatch)
		require.Nil(t, newRec)
	})
	t.Run("fail with invalid direction", func(t *testing.T) {
		manager := NewAuditManager()

		now := time.Now()

		lastRec := &AuditRecord{
			Direction: In,
			Time:      now.Add(-time.Hour),
			Name:      "ivan",
		}

		newRec, err := manager.AddRecord(lastRec, "ivan", now, In)
		require.ErrorIs(t, err, ErrInvalidDirection)
		require.Nil(t, newRec)
	})
	t.Run("fail with invalid time", func(t *testing.T) {
		manager := NewAuditManager()

		now := time.Now()

		lastRec := &AuditRecord{
			Direction: In,
			Time:      now,
			Name:      "ivan",
		}

		newRec, err := manager.AddRecord(lastRec, "ivan", now.Add(-time.Hour), In)
		require.ErrorIs(t, err, ErrInvalidTime)
		require.Nil(t, newRec)
	})
}
