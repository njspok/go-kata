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

type PriceCalc interface {
	Calc(size Size, leave, entry time.Time) (Price, error)
}

func NewSimplePriceCalc(hourlyRates map[Size]Price) *SimplePriceCalc {
	return &SimplePriceCalc{hourlyRates: hourlyRates}
}

type SimplePriceCalc struct {
	hourlyRates map[Size]Price
}

func (s *SimplePriceCalc) Calc(size Size, entry, leave time.Time) (Price, error) {
	rate, ok := s.hourlyRates[size]
	if !ok {
		return 0, fmt.Errorf("no tariff for size %q", size)
	}

	fullHours := int(leave.Sub(entry)/time.Hour) + 1

	return rate.Times(fullHours), nil
}

func NewParking(slots Slots, calc PriceCalc) *Parking {
	return &Parking{slots: slots, calc: calc}
}

type Parking struct {
	slots Slots
	calc  PriceCalc
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
	price, err := p.calc.Calc(ticket.Size, ticket.EntryTime, leaveTime)
	if err != nil {
		return 0, err
	}

	ticket.Price = price
	ticket.LeaveTime = leaveTime

	p.slots.Find(ticket.Number).Leave()

	return ticket.Price, nil
}
