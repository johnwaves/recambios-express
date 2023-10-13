package pkg

type Pedido interface {
    ID() int             // Obtener el identificador único del pedido
    ClienteID() int      // Obtener el identificador del cliente asociado al pedido
    AsignarRepartidor(r Repartidor) // Asignar un repartidor al pedido
}
