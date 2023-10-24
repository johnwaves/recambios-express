package src

type Pieza struct {
	id     int
	nombre string
}

func NuevaPieza(id int) *Pieza {
	return &Pieza{
		id: id,
	}
}

func (p *Pieza) NOMBRE() string {
	return p.nombre
}

func (p *Pieza) ID() int {
	return p.id
}
