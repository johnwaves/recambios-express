package pkg

type Ruta struct {
	pedidosPendientes []Pedido
}

func NuevoRepartidor() *Ruta {
	return &Ruta{}
}
