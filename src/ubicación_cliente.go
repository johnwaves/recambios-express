package tienda

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
