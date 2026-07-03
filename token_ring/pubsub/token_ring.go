package pubsub

type Message any

type Marker struct {
	To  NodeID
	TTL int
}

type NodeID int

type MessageType string

func NewNode(nodeID NodeID, do func(), broker Broker) *Node {
	return &Node{
		nodeID:     nodeID,
		nextNodeID: nodeID,
		do:         do,
		broker:     broker,
	}
}

type Node struct {
	nodeID     NodeID
	nextNodeID NodeID
	do         func()
	broker     Broker
}

func (n *Node) ID() NodeID {
	return n.nodeID
}

func (n *Node) SetNextNodeID(next NodeID) {
	n.nextNodeID = next
}

func (n *Node) Send(msg Message) {
	switch m := msg.(type) {
	case Marker:
		if m.To == n.nodeID {
			if m.TTL == 0 {
				return
			}

			n.do()

			n.broker.Publish(Marker{
				To:  n.nextNodeID,
				TTL: m.TTL - 1,
			})
			return
		}
	default:
		panic("unexpected message type")
	}
}

type Broker interface {
	Publish(msg Message)
}

func NewLightBroker() *LightBroker {
	return &LightBroker{}
}

type LightBroker struct {
	nodes []*Node
}

func (b *LightBroker) Add(n *Node) {
	b.nodes = append(b.nodes, n)
}

func (b *LightBroker) Publish(msg Message) {
	for _, node := range b.nodes {
		node.Send(msg)
	}
}
