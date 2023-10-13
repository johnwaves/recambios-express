package pkg

type Pedido interface {
    ID() int             // Obtener el identificador único del pedido
    ClienteID() int      // Obtener el identificador del cliente asociado al pedido
    AsignarRepartidor(r Repartidor) // Asignar un repartidor al pedido
}

type PedidoRepository interface {
    ObtenerPorID(id int) (Pedido, error)   // Obtener un pedido por su ID
    Guardar(pedido Pedido) error           // Guardar un pedido
    AsignarRepartidor(pedido Pedido, r Repartidor) error // Asignar un repartidor a un pedido
}
