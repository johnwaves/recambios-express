package pkg

type EncargadoAsignacionPedidosRutas struct {
	pedidosPendientes       []Pedido
	repartidoresDisponibles []Repartidor
	repartosActivos         map[*Pedido]map[*Repartidor]*HoraEntrega //Ojo con esto que son punteros para que sean comparables.
}

func NuevoEncargadoAsignacionPedidosRutas() *EncargadoAsignacionPedidosRutas {
	return &EncargadoAsignacionPedidosRutas{}
}
