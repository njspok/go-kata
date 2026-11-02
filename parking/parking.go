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

type Slots struct {
	Moto  int
	Car   int
	Truck int
}

func (s *Slots) Of(size Size) (int, bool) {
	switch size {
	case Moto:
		return s.Moto, true
	case Car:
		return s.Car, true
	case Truck:
		return s.Truck, true
	}
	return 0, false
}

func (s *Slots) Add(size Size, delta int) {
	switch size {
	case Moto:
		s.Moto += delta
	case Car:
		s.Car += delta
	case Truck:
		s.Truck += delta
	}
}

type Ticket struct {
	Size      Size
	EntryTime time.Time
	LeaveTime time.Time
	Price     Price
}

type Parking struct {
	capacity Slots
	occupied Slots
}

func NewParking(capacity Slots) *Parking {
	return &Parking{capacity: capacity}
}

func (p *Parking) Park(size Size, entryTime time.Time) (*Ticket, error) {
	total, ok := p.capacity.Of(size)
	if !ok {
		return nil, fmt.Errorf("unknown size %q", size)
	}
	used, _ := p.occupied.Of(size)
	if used >= total {
		return nil, ErrNoFreeSlots
	}
	p.occupied.Add(size, 1)
	return &Ticket{Size: size, EntryTime: entryTime}, nil
}

func (p *Parking) Leave(ticket *Ticket, leaveTime time.Time) (Price, error) {
	rate, ok := hourlyRates[ticket.Size]
	if !ok {
		return 0, fmt.Errorf("no tariff for size %q", ticket.Size)
	}
	ticket.LeaveTime = leaveTime
	hours := int(leaveTime.Sub(ticket.EntryTime)/time.Hour) + 1
	ticket.Price = Price(hours) * rate
	p.occupied.Add(ticket.Size, -1)
	return ticket.Price, nil
}
