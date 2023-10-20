package src

import (
	"sort"
	"time"
)

type JefeRecambios struct {
	pedidosPendientes       []Pedido
	repartidoresDisponibles []Repartidor
	repartosActivos         map[*Pedido]map[*Repartidor]*HoraEntrega //Ojo con esto que son punteros para que sean comparables.
}

func NuevoJefeRecambios() *JefeRecambios {
	return &JefeRecambios{}
}

func (j *JefeRecambios) NUEVOPEDIDO(p Pedido) {
	j.pedidosPendientes = append(j.pedidosPendientes, p)

	// Ordenar la lista de pedidos por tiempo estimado
	sort.Slice(j.pedidosPendientes, func(i, h int) bool {
		return j.pedidosPendientes[i].UBICACION().tiempoEstimadoMinutos > j.pedidosPendientes[h].UBICACION().tiempoEstimadoMinutos
	})
}

func (j *JefeRecambios) REPARTIDORLIBRE(r Repartidor) {
	j.repartidoresDisponibles = append(j.repartidoresDisponibles, r)
}

func (j *JefeRecambios) PEDIDOSPENDIENTES() []Pedido {
	return j.pedidosPendientes
}

func (j *JefeRecambios) REPARTIDORESDISPONIBLES() []Repartidor {
	return j.repartidoresDisponibles
}

func (j *JefeRecambios) ASIGNARREPARTO() Pedido {
	pedido := j.pedidosPendientes[0]
	repartidor := j.repartidoresDisponibles[0]

	tiempo := pedido.UBICACION().tiempoEstimadoMinutos
	ahora := time.Now()
	nuevaHora := ahora.Add(time.Duration(tiempo) * time.Minute)
	nuevaHoraComoString := nuevaHora.Format("15:04:05")

	horaEntrega := NuevoHoraEntrega(nuevaHoraComoString)

	reparto := make(map[*Repartidor]*HoraEntrega)
	reparto[&repartidor] = horaEntrega

	repartosActivos := make(map[*Pedido]map[*Repartidor]*HoraEntrega)
	repartosActivos[&pedido] = reparto

	return pedido
}

func (j *JefeRecambios) INFOPEDIDO(p *Pedido) map[*Repartidor]*HoraEntrega {
	return j.repartosActivos[p]
}
