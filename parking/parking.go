package parking

import (
	"fmt"
	"time"
)

type Size string

const (
	Moto  Size = "Moto"
	Car   Size = "Car"
	Truck Size = "Truck"
)

type Price int

func (p Price) Times(val int) Price {
	return p * Price(val)
}

var hourlyRates = map[Size]Price{
	Moto:  50,
	Car:   100,
	Truck: 200,
}

type Ticket struct {
	Size      Size
	EntryTime time.Time
	LeaveTime time.Time
	Price     Price
}

type Parking struct{}

func NewParking() *Parking {
	return &Parking{}
}

func (p *Parking) Park(size Size, entryTime time.Time) (*Ticket, error) {
	return &Ticket{Size: size, EntryTime: entryTime}, nil
}

func (p *Parking) Leave(ticket *Ticket, leaveTime time.Time) (Price, error) {
	rate, ok := hourlyRates[ticket.Size]
	if !ok {
		return 0, fmt.Errorf("no tariff for size %q", ticket.Size)
	}
	ticket.LeaveTime = leaveTime
	hours := int(leaveTime.Sub(ticket.EntryTime)/time.Hour) + 1
	ticket.Price = rate.Times(hours)
	return ticket.Price, nil
}
