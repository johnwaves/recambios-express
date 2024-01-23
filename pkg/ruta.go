package pkg

const MaxPedidos int = 10

type Ruta struct {
	pedidosConDistancia [MaxPedidos]PedidoConDistancia
	cantidadActual      int
}
