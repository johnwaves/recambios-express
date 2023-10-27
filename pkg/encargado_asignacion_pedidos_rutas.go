package pkg

type EncargadoAsignacionPedidosRutas struct {
	pedidosPendientes []Pedido
	rutas             []Ruta
	repartosActivos   map[*Pedido]map[*Ruta]*HoraEntrega //Ojo con esto que son punteros para que sean comparables.
}

func NuevoEncargadoAsignacionPedidosRutas() *EncargadoAsignacionPedidosRutas {
	return &EncargadoAsignacionPedidosRutas{}
}
