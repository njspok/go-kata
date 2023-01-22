package mock_version

import (
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
}
