package parking

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	calc := NewSimplePriceCalc(map[Size]Price{
		Moto:  50,
		Car:   100,
		Truck: 200,
	})

	t.Run("success create parking", func(t *testing.T) {
		slots := Slots{
			{
				Number: "01",
				Size:   Car,
			},
		}
		p := NewParking(slots, calc)
		require.NotNil(t, p)
	})
	t.Run("park", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			// Arrange
			slots := Slots{
				{
					Number: "01",
					Size:   Car,
				},
			}
			p := NewParking(slots, nil)

			now := mustParseTime("2026-01-01 01:00:00")

			// Act
			ticket, err := p.Park(Car, now)

			// Assert
			require.NoError(t, err)
			require.NotNil(t, ticket)
			require.Equal(t, Car, ticket.Size)
			require.Equal(t, now, ticket.EntryTime)
		})
		t.Run("no free slots", func(t *testing.T) {
			// Arrange
			slots := Slots{}
			p := NewParking(slots, calc)
			now := mustParseTime("2026-01-01 01:00:00")

			// Act
			ticket, err := p.Park(Car, now)

			// Assert
			require.ErrorIs(t, err, ErrNoFreeSlots)
			require.Nil(t, ticket)
		})
	})
	t.Run("leave", func(t *testing.T) {
		table := []struct {
			size      Size
			leaveTime string
			price     Price
			err       error
		}{
			// for Car
			{size: Car, leaveTime: "2026-01-01 01:00:00", price: 100, err: nil},
			{size: Car, leaveTime: "2026-01-01 01:30:00", price: 100, err: nil},
			{size: Car, leaveTime: "2026-01-01 01:59:59", price: 100, err: nil},
			{size: Car, leaveTime: "2026-01-01 02:00:00", price: 200, err: nil},

			// for Moto
			{size: Moto, leaveTime: "2026-01-01 01:00:00", price: 50, err: nil},
			{size: Moto, leaveTime: "2026-01-01 01:30:00", price: 50, err: nil},
			{size: Moto, leaveTime: "2026-01-01 01:59:59", price: 50, err: nil},
			{size: Moto, leaveTime: "2026-01-01 02:00:00", price: 100, err: nil},

			// for Truck
			{size: Truck, leaveTime: "2026-01-01 01:00:00", price: 200, err: nil},
			{size: Truck, leaveTime: "2026-01-01 01:30:00", price: 200, err: nil},
			{size: Truck, leaveTime: "2026-01-01 01:59:59", price: 200, err: nil},
			{size: Truck, leaveTime: "2026-01-01 02:00:00", price: 400, err: nil},
		}

		for _, tt := range table {
			t.Run(tt.leaveTime, func(t *testing.T) {
				// Arrange
				slots := Slots{
					{Number: "C1", Size: Car},
					{Number: "M1", Size: Moto},
					{Number: "T1", Size: Truck},
				}
				p := NewParking(slots, calc)
				parkTime := mustParseTime("2026-01-01 01:00:00")
				ticket, err := p.Park(tt.size, parkTime)

				// Act
				price, err := p.Leave(ticket, mustParseTime(tt.leaveTime))

				// Assert
				require.NoError(t, err)
				require.Equal(t, tt.price, price)
				require.Equal(t, tt.price, ticket.Price)
				require.Equal(t, parkTime, ticket.EntryTime)
				require.Equal(t, mustParseTime(tt.leaveTime), ticket.LeaveTime)
			})
		}
	})
}

func mustParseTime(str string) time.Time {
	t, err := time.Parse(time.DateTime, str)
	if err != nil {
		panic(err)
	}
	return t
}
