package func_shop

import "github.com/samber/mo"

type Cart map[Sid]*Item

func (c Cart) Add(name string, sid Sid) Cart {
	item := c.Find(sid)
	if i, present := item.Get(); present {
		c[sid] = i.IncQty()
	} else {
		c[sid] = NewItem(name, sid)
	}
	return c
}

func (c Cart) Find(sid Sid) mo.Option[Item] {
	if i, exist := c[sid]; exist {
		return mo.Some(*i)
	}
	return mo.None[Item]()
}

type Sid string

type Item struct {
	name string
	sid  Sid
	qty  uint
}

func (i Item) Name() string {
	return i.name
}

func (i Item) Sid() Sid {
	return i.sid
}

func (i Item) Qty() uint {
	return i.qty
}

func (i Item) IncQty() *Item {
	i.qty++
	return &i
}

func (i Item) SetQty(q uint) *Item {
	i.qty = q
	return &i
}

func NewItem(name string, sid Sid) *Item {
	return &Item{
		name: name,
		sid:  sid,
		qty:  1,
	}
}
