package src

type UbicacionCliente struct {
	calle                   string
	numero                  int
	poblacion               string
	codigo_postal           int
	tiempo_estimado_minutos int
}

func NuevaUbicacionCliente(calle string, numero int, poblacion string, codigo_postal int, tiempo_estimado_minutos int) *UbicacionCliente {
	return &UbicacionCliente{
		calle:                   calle,
		numero:                  numero,
		poblacion:               poblacion,
		codigo_postal:           codigo_postal,
		tiempo_estimado_minutos: tiempo_estimado_minutos,
	}
}

func (u *UbicacionCliente) CALLE() string {
	return u.calle
}

func (u *UbicacionCliente) NUMERO() int {
	return u.numero
}

func (u *UbicacionCliente) POBLACION() string {
	return u.poblacion
}

func (u *UbicacionCliente) CODIGO_POSTAL() int {
	return u.codigo_postal
}

func (u *UbicacionCliente) TIEMPO_ESTIMADO() int {
	return u.tiempo_estimado_minutos
}
