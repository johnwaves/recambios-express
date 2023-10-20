package src

type Repartidor struct {
	nombreCompleto    string
	pedidosPendientes []Pedido
}

func NuevoRepartidor(nombreCompleto string) *Repartidor {
	return &Repartidor{
		nombreCompleto: nombreCompleto,
	}
}

func (r *Repartidor) NOMBRE() string {
	return r.nombreCompleto
}

func (r *Repartidor) PEDIDOS() []Pedido {
	return r.pedidosPendientes
}

func (r *Repartidor) ASIGNAR_PEDIDO(pedido Pedido) {
	r.pedidosPendientes = append(r.pedidosPendientes, pedido)
}
