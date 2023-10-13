package cliente

type Cliente struct {
	id        int
	distancia float64
}

func NuevoCliente(id int, distancia float64) *Cliente {
	return &Cliente{
		id:        id,
		distancia: distancia,
	}
}

func (c *Cliente) ID() int {
	return c.id
}

func (c *Cliente) Distancia() float64 {
	return c.distancia
}
