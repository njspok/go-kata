package sagas

func NewOrder(id int, clientId int, itemId int, qty int) *Order {
	return &Order{
		id:       id,
		clientId: clientId,
		itemId:   itemId,
		qty:      qty,
	}
}

type Order struct {
	id       int
	clientId int
	itemId   int
	qty      int
}
