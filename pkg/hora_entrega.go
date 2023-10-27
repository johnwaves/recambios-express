package pkg

type HoraEntrega struct {
	horaEstimada Time
}

func NuevoHoraEntrega(horaEstimada Time) *HoraEntrega {
	return &HoraEntrega{
		horaEstimada: horaEstimada,
	}
}
