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

func (p Price) Times(v int) Price {
	return Price(v) * p
}

var hourlyRates = map[Size]Price{
	Moto:  50,
	Car:   100,
	Truck: 200,
}

type Slot struct {
	Number     string
	Size       Size
	IsOccupied bool
}

func (s *Slot) Match(size Size) bool {
	return s.Size == size && !s.IsOccupied
}

func (s *Slot) Park() {
	s.IsOccupied = true
}

func (s *Slot) Leave() {
	s.IsOccupied = false
}

type Slots []*Slot

func (ss Slots) Find(number string) *Slot {
	for _, s := range ss {
		if s.Number == number {
			return s
		}
	}
	return nil
}

func NewTicket(number string, size Size, tm time.Time) *Ticket {
	return &Ticket{
		Number:    number,
		Size:      size,
		EntryTime: tm,
		LeaveTime: time.Time{},
		Price:     0,
	}
}

type Ticket struct {
	Number    string
	Size      Size
	EntryTime time.Time
	LeaveTime time.Time
	Price     Price
}

type Parking struct {
	slots Slots
}

func NewParking(slots Slots) *Parking {
	return &Parking{slots: slots}
}

func (p *Parking) Park(size Size, entryTime time.Time) (*Ticket, error) {
	for _, slot := range p.slots {
		if !slot.Match(size) {
			continue
		}

		slot.Park()

		return NewTicket(slot.Number, size, entryTime), nil
	}
	return nil, ErrNoFreeSlots
}

func (p *Parking) Leave(ticket *Ticket, leaveTime time.Time) (Price, error) {
	rate, ok := hourlyRates[ticket.Size]
	if !ok {
		return 0, fmt.Errorf("no tariff for size %q", ticket.Size)
	}

	fullHours := int(leaveTime.Sub(ticket.EntryTime)/time.Hour) + 1

	ticket.Price = rate.Times(fullHours)
	ticket.LeaveTime = leaveTime

	p.slots.Find(ticket.Number).Leave()

	return ticket.Price, nil
}
