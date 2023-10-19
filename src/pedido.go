package src

type Pedido struct {
	id        int
	piezas    []Pieza
	ubicacion UbicacionCliente
}

func NuevoPedido(id int, piezas []Pieza, ubicacion UbicacionCliente) *Pedido {
	return &Pedido{
		id:        id,
		piezas:    piezas,
		ubicacion: ubicacion,
	}
}

func (p *Pedido) ID() int {
	return p.id
}

func (p *Pedido) PIEZAS() []Pieza {
	return p.piezas
}

func (p *Pedido) UBICACION() UbicacionCliente {
	return p.ubicacion
}
