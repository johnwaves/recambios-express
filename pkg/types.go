package pkg

const ALMACEN_LON float32 = 37.5836344
const ALMACEN_LAT float32 = -1.7863288

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
