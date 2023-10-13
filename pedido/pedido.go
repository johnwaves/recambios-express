package pedido

type Pedido struct {
	id         int
	clienteID  int
	repartidor Repartidor
}

func NuevoPedido(id, clienteID int) *Pedido {
	return &Pedido{
		id:        id,
		clienteID: clienteID,
	}
}

func (p *Pedido) ID() int {
	return p.id
}

func (p *Pedido) ClienteID() int {
	return p.clienteID
}

func (p *Pedido) AsignarRepartidor(r Repartidor) {
	p.repartidor = r
}
