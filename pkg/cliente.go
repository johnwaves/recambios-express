package pkg

type Cliente interface {
    ID() int          // Obtener el identificador único del cliente
    Distancia() float64  // Obtener la distancia del cliente al almacén
}

type ClienteRepository interface {
    ObtenerPorID(id int) (Cliente, error)   // Obtener un cliente por su ID
    Guardar(cliente Cliente) error         // Guardar un cliente
}
