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
