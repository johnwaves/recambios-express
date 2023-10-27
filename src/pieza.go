package tienda

type Pieza struct {
	id     int
	nombre string
}

func NuevaPieza(id int) *Pieza {
	return &Pieza{
		id: id,
	}
}
