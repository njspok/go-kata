package func_shop

import (
	"cmp"
	"iter"
	"maps"

	"github.com/samber/lo"
	"github.com/samber/mo"
)

type CouponName string

const (
	BigCart CouponName = "BIGCART"
)

type Coupon interface {
	Apply(cart Cart) Price
}

type BigCartCoupon struct{}

func (b BigCartCoupon) Apply(cart Cart) Price {
	return lo.Ternary(cart.Total() > 1000, cart.Total()-100, cart.Total())
}

var Coupons = map[CouponName]Coupon{
	BigCart: BigCartCoupon{},
}

func CalcDiscount(cart Cart, coupon CouponName, coupons map[CouponName]Coupon) Price {
	c := Lookup(coupons, coupon)
	if c.IsPresent() {
		return c.MustGet().Apply(cart)
	}
	return cart.Total()
}

type Cart map[Sid]*Item

func NewCart() Cart {
	return make(Cart)
}

func (c Cart) Total() Price {
	return Reduce(c, func(total Price, item *Item, _ Sid) Price {
		return total + item.Total()
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

func NewItem(name string, sid Sid, price Price) *Item {
	return &Item{
		name:  name,
		sid:   sid,
		qty:   1,
		price: price,
	}
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

func Filter[V any](seq iter.Seq[V], keep func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if keep(v) && !yield(v) {
				return
			}
		}
	}
}

func Lookup[K comparable, V any](m map[K]V, key K) mo.Option[V] {
	v, ok := m[key]
	if !ok {
		return mo.None[V]()
	}
	return mo.Some(v)
}
