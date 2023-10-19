package src

type Pieza struct {
	id int
}

func NuevaPieza(id int) *Pieza {
	return &Pieza{
		id: id,
	}
}

func (p *Pieza) ID() int {
	return p.id
}
