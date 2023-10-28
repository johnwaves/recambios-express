package pkg

type Cliente struct {
	nombreCliente         string
	calle                 string
	numero                int
	poblacion             string
	codigoPostal          int
	tiempoEstimadoMinutos int
}

func NuevoCliente(nombreCliente string, calle string, numero int, poblacion string, codigo_postal int, tiempo_estimado_minutos int) *Cliente {
	return &Cliente{
		nombreCliente:         nombreCliente,
		calle:                 calle,
		numero:                numero,
		poblacion:             poblacion,
		codigoPostal:          codigo_postal,
		tiempoEstimadoMinutos: tiempo_estimado_minutos,
	}
}
