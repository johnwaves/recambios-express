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
