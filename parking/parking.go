package parking

import (
	"errors"
	"fmt"
	"time"
)

type Size string

const (
	Moto  Size = "Moto"
	Car   Size = "Car"
	Truck Size = "Truck"
)

var ErrNoFreeSlots = errors.New("no free slots")

type Price int

var hourlyRates = map[Size]Price{
	Moto:  50,
	Car:   100,
	Truck: 200,
}

type Slot struct {
	Number string
	Size   Size
}

type Slots []Slot

type Ticket struct {
	Number    string
	Size      Size
	EntryTime time.Time
	LeaveTime time.Time
	Price     Price
}

type Parking struct {
	slots    Slots
	occupied map[string]bool
}

func NewParking(slots Slots) *Parking {
	return &Parking{slots: slots, occupied: make(map[string]bool)}
}

func (p *Parking) Park(size Size, entryTime time.Time) (*Ticket, error) {
	for _, s := range p.slots {
		if s.Size != size || p.occupied[s.Number] {
			continue
		}

		p.occupied[s.Number] = true

		return &Ticket{Number: s.Number, Size: size, EntryTime: entryTime}, nil
	}
	return nil, ErrNoFreeSlots
}

func (p *Parking) Leave(ticket *Ticket, leaveTime time.Time) (Price, error) {
	rate, ok := hourlyRates[ticket.Size]
	if !ok {
		return 0, fmt.Errorf("no tariff for size %q", ticket.Size)
	}

	ticket.LeaveTime = leaveTime

	hours := int(leaveTime.Sub(ticket.EntryTime)/time.Hour) + 1
	ticket.Price = Price(hours) * rate

	delete(p.occupied, ticket.Number)

	return ticket.Price, nil
}
