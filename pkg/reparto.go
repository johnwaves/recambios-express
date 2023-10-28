package pkg

type Reparto struct {
	pedidosPendientes []Pedido
	ruta              map[*Pedido]*HoraEntrega //Aqui ya se incluyen ordenados por prioridad.
}

func NuevoReparto() *Reparto {
	return &Reparto{}
}
