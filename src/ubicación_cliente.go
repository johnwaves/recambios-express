package src

type UbicacionCliente struct {
	calle                 string
	numero                int
	poblacion             string
	codigoPostal          int
	tiempoEstimadoMinutos int
}

func NuevaUbicacionCliente(calle string, numero int, poblacion string, codigo_postal int, tiempo_estimado_minutos int) *UbicacionCliente {
	return &UbicacionCliente{
		calle:                 calle,
		numero:                numero,
		poblacion:             poblacion,
		codigoPostal:          codigo_postal,
		tiempoEstimadoMinutos: tiempo_estimado_minutos,
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
	return u.codigoPostal
}

func (u *UbicacionCliente) TIEMPO_ESTIMADO() int {
	return u.tiempoEstimadoMinutos
}
