package src

type HoraEntrega struct {
	hora_estimada string
}

func NuevoHoraEntrega(hora_estimada string) *HoraEntrega {
	return &HoraEntrega{
		hora_estimada: hora_estimada,
	}
}

func (h *HoraEntrega) HORAENTREGA() string {
	return h.hora_estimada
}
