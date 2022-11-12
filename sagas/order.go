package sagas

func NewOrder(id int, itemId int, qty int) *Order {
	return &Order{
		id:     id,
		itemId: itemId,
		qty:    qty,
	}
}

type Order struct {
	id     int
	itemId int
	qty    int
}
