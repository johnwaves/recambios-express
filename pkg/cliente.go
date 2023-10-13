package pkg

type Cliente interface {
    ID() int          // Obtener el identificador único del cliente
    Distancia() float64  // Obtener la distancia del cliente al almacén
}
