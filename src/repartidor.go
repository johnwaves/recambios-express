package src

type Repartidor struct {
	nombre_completo    string
	pedidos_pendientes []Pedido
}

func NuevoRepartidor(nombre_completo string) *Repartidor {
	return &Repartidor{
		nombre_completo: nombre_completo,
	}
}

func (r *Repartidor) NOMBRE() string {
	return r.nombre_completo
}

func (r *Repartidor) PEDIDOS() []Pedido {
	return r.pedidos_pendientes
}

func (r *Repartidor) ASIGNAR_PEDIDO(pedido Pedido) {
	r.pedidos_pendientes = append(r.pedidos_pendientes, pedido)
}
