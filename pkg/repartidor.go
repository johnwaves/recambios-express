package pkg

type Repartidor interface {
    ID() int          // Obtener el identificador único del repartidor
    Horario() Horario // Obtener el horario del repartidor
}

type Horario interface {
    Inicio() string  // Obtener el horario de inicio de trabajo del repartidor
    Fin() string     // Obtener el horario de fin de trabajo del repartidor
}

type RepartidorRepository interface {
    ObtenerPorID(id int) (Repartidor, error)   // Obtener un repartidor por su ID
    Guardar(repartidor Repartidor) error      // Guardar un repartidor
}
