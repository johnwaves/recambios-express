package tienda

type HoraEntrega struct {
	horaEstimada string
}

func NuevoHoraEntrega(horaEstimada string) *HoraEntrega {
	return &HoraEntrega{
		horaEstimada: horaEstimada,
	}
}
