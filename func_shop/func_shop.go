package func_shop

import (
	"reflect"

	"github.com/samber/mo"
)

type Cart map[Sid]*Item

func NewCart() Cart {
	return make(Cart)
}

func (c Cart) Add(name string, sid Sid) Cart {
	item := c.Find(sid)
	if i, present := item.Get(); present {
		c[sid] = i.IncQty()
	} else {
		c[sid] = NewItem(name, sid)
	}
	return c
}

func (c Cart) Find(sid Sid) mo.Option[*Item] {
	if i, exist := c[sid]; exist {
		return mo.Some(Copy(i))
	}
	return mo.None[*Item]()
}

func (c Cart) Count() int {
	return len(c)
}

type Sid string

type Qty int

type Item struct {
	name string
	sid  Sid
	qty  Qty
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

func (i *Item) IncQty() *Item {
	i.qty++
	return i
}

func NewItem(name string, sid Sid) *Item {
	return &Item{
		name: name,
		sid:  sid,
		qty:  1,
	}
}

func Copy[T any](src T) T {
	v := reflect.ValueOf(src)
	if v.Kind() != reflect.Pointer {
		return src
	}
	if v.IsNil() {
		return src
	}

	cp := reflect.New(v.Elem().Type())
	cp.Elem().Set(v.Elem())
	return cp.Interface().(T)
}
