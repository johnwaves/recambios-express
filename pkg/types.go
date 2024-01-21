package pkg

type Direccion struct {
	calle_numero  string
	poblacion     string
	codigo_postal int
}

func NewDireccion(calle_numero string, poblacion string, codigo_postal int) *Direccion {
	return &Direccion{
		calle_numero:  calle_numero,
		poblacion:     poblacion,
		codigo_postal: codigo_postal,
	}
}
