package pkg

type Direccion struct {
	calle_numero  string
	poblacion     string
	codigo_postal int
}

func NewDireccion(calleNumero string, poblacion string, codigoPostal int) *Direccion {
	return &Direccion{
		calle_numero:  calleNumero,
		poblacion:     poblacion,
		codigo_postal: codigoPostal,
	}
}
