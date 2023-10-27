package src

type JefeRecambios struct {
	pedidosPendientes       []Pedido
	repartidoresDisponibles []Repartidor
	repartosActivos         map[*Pedido]map[*Repartidor]*HoraEntrega //Ojo con esto que son punteros para que sean comparables.
}

func NuevoJefeRecambios() *JefeRecambios {
	return &JefeRecambios{}
}
