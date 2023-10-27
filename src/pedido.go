package tienda

type Pedido struct {
	id        int
	ubicacion UbicacionCliente
}

func NuevoPedido(id int, ubicacion UbicacionCliente) *Pedido {
	return &Pedido{
		id:        id,
		ubicacion: ubicacion,
	}
}
