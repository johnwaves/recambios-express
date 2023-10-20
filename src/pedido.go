package src

type Pedido struct {
	id        int
	piezas    map[Pieza]int
	ubicacion UbicacionCliente
}

func NuevoPedido(id int, piezas map[Pieza]int, ubicacion UbicacionCliente) *Pedido {
	return &Pedido{
		id:        id,
		piezas:    piezas,
		ubicacion: ubicacion,
	}
}

func (p *Pedido) ID() int {
	return p.id
}

func (p *Pedido) PIEZAS() map[Pieza]int {
	return p.piezas
}

func (p *Pedido) UBICACION() UbicacionCliente {
	return p.ubicacion
}
