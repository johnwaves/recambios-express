package pkg

type Pedido struct {
	id      int
	cliente Cliente
}

func NuevoPedido(id int, cliente Cliente) *Pedido {
	return &Pedido{
		id:      id,
		cliente: cliente,
	}
}
