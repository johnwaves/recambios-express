package pkg

type Pedido struct {
	cliente *Cliente
	items   []ItemPedido
}

func NewPedido(items []ItemPedido) *Pedido {
	return &Pedido{
		items: items,
	}
}

func (p *Pedido) asignarCliente(c *Cliente) {
	p.cliente = c
}

func (p *Pedido) Cliente() *Cliente {
	return p.cliente
}

func (p *Pedido) Items() []ItemPedido {
	return p.items
}
