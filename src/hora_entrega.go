package src

type HoraEntrega struct {
	horaEstimada string
}

func NuevoHoraEntrega(horaEstimada string) *HoraEntrega {
	return &HoraEntrega{
		horaEstimada: horaEstimada,
	}
}

func (h *HoraEntrega) HORAENTREGA() string {
	return h.horaEstimada
}
