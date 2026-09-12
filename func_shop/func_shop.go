package func_shop

import (
	"cmp"
	"maps"

	"github.com/samber/mo"
)

type Cart map[Sid]*Item

func NewCart() Cart {
	return make(Cart)
}

func (c Cart) Total() Price {
	return Reduce(c, func(agg Price, val *Item, key Sid) Price {
		return agg + val.Total()
	}, Price(0))
}

func (c Cart) Del(sid Sid) Cart {
	cp := maps.Clone(c)
	delete(cp, sid)
	return cp
}

func (c Cart) Add(name string, sid Sid, price Price) Cart {
	cp := maps.Clone(c)
	item := cp.Find(sid)

	if i, present := item.Get(); present {
		cp[sid] = i.IncQty().SetPrice(price)
	} else {
		cp[sid] = NewItem(name, sid, price)
	}
	return cp
}

func (c Cart) Find(sid Sid) mo.Option[*Item] {
	if i, exist := c[sid]; exist {
		return mo.Some(i)
	}
	return mo.None[*Item]()
}

func (c Cart) Count() int {
	return len(c)
}

type Sid string

type Qty int

type Price int

func (p Price) Mult(v Qty) Price {
	return Price(v) * p
}

type Item struct {
	name  string
	sid   Sid
	qty   Qty
	price Price
}

func (i Item) Name() string {
	return i.name
}

func (i Item) Sid() Sid {
	return i.sid
}

func (i Item) Qty() Qty {
	return i.qty
}

func (i Item) IncQty() *Item {
	i.qty++
	return &i
}
func (i Item) SetPrice(price Price) *Item {
	i.price = price
	return &i
}

func (i Item) Price() Price {
	return i.price
}

func (i Item) Total() Price {
	return i.Price().Mult(i.Qty())
}

func NewItem(name string, sid Sid, price Price) *Item {
	return &Item{
		name:  name,
		sid:   sid,
		qty:   1,
		price: price,
	}
}

func Reduce[K cmp.Ordered, V any, R any](
	collection map[K]V,
	accumulator func(agg R, val V, key K) R,
	initial R,
) R {
	for i, item := range collection {
		initial = accumulator(initial, item, i)
	}

	return initial
}
