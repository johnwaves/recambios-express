package pkg

import (
	"errors"
	"sort"
)

type Ruta struct {
	pedidos []PedidoConDistancia
}

func NewRuta() *Ruta {
	return &Ruta{}
}

func (r *Ruta) AddPedido(pedido *Pedido) error {
	cliente := pedido.Cliente()

	if cliente.Nombre() == "" {
		return errors.New("El nombre del cliente no puede estar vacío.")
	}

	direccion := cliente.direccion_entrega
	if direccion.calle_numero == "" || direccion.poblacion == "" || direccion.codigo_postal == 0 {
		return errors.New("Los datos de la dirección postal están incompletos.")
	}

	if len(pedido.items) == 0 {
		return errors.New("El pedido debe contener al menos una pieza.")
	} else {
		for _, item := range pedido.items {
			if item.pieza == "" {
				return errors.New("Cada pieza debe tener un nombre.")
			}
		}
	}

	for _, item := range pedido.items {
		if item.cantidad <= 0 {
			return errors.New("Cada pieza debe tener una cantidad mayor a cero.")
		}
	}

	coordenadasDestino, _ := ObtenerCoordenadas(direccion)

	if coordenadasDestino.Lat != 0.0 && coordenadasDestino.Lon != 0.0 {
		distancia := Haversine(AlmacenCoordenadas, coordenadasDestino)

		pedidoConDistancia := PedidoConDistancia{
			Pedido:    *pedido,
			Distancia: distancia,
		}

		r.pedidos = append(r.pedidos, pedidoConDistancia)
	} else {
		return errors.New("La dirección de entrega no es válida.")
	}

	return nil
}

func (r *Ruta) CantidadActual() int {
	return len(r.pedidos)
}

func (r *Ruta) distanciaMaxima() float64 {
	var maxDistancia float64
	for _, pedido := range r.pedidos {
		if pedido.Distancia > maxDistancia {
			maxDistancia = pedido.Distancia
		}
	}
	return maxDistancia
}

func (r *Ruta) distanciaMedia() float64 {
	var sumaDistancia float64
	for _, pedido := range r.pedidos {
		sumaDistancia += pedido.Distancia
	}
	return sumaDistancia / float64(len(r.pedidos))
}

func (r *Ruta) OrdenarPedidos() {
	distanciaMedia := r.distanciaMedia()
	distanciaMaxima := r.distanciaMaxima()

	var pedidosCortos, pedidosMedios, pedidosLargos []PedidoConDistancia
	for _, pedido := range r.pedidos {
		if pedido.Distancia < distanciaMedia {
			pedidosCortos = append(pedidosCortos, pedido)
		} else if pedido.Distancia < distanciaMaxima*porcentajePedidosLargos {
			pedidosMedios = append(pedidosMedios, pedido)
		} else {
			pedidosLargos = append(pedidosLargos, pedido)
		}
	}

	sort.Slice(pedidosLargos, func(i, j int) bool {
		return pedidosLargos[i].Distancia > pedidosLargos[j].Distancia
	})

	r.pedidos = []PedidoConDistancia{}

	if len(pedidosCortos) == 0 && len(pedidosMedios) == 0 {
		r.pedidos = append(r.pedidos, pedidosLargos...)
	} else {
		i, j, k := 0, 0, 0
		for i < len(pedidosLargos) || j < len(pedidosCortos) || k < len(pedidosMedios) {
			for count := 0; i < len(pedidosLargos) && count < 2; count++ {
				r.pedidos = append(r.pedidos, pedidosLargos[i])
				i++
			}

			if j < len(pedidosCortos) {
				r.pedidos = append(r.pedidos, pedidosCortos[j])
				j++
			} else if k < len(pedidosMedios) {
				r.pedidos = append(r.pedidos, pedidosMedios[k])
				k++
			}
		}
	}
}
