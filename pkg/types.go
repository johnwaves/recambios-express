package pkg

import "fmt"

type Direccion struct {
	calle_numero  string
	poblacion     string
	codigo_postal int
}

func NewDireccion(calle_numero string, poblacion string, codigo_postal int) *Direccion {
	return &Direccion{
		calle_numero:  calle_numero,
		poblacion:     poblacion,
		codigo_postal: codigo_postal,
	}
}

func (d *Direccion) FormatDireccion() string {
	return fmt.Sprintf("%s, %s, %d", d.calle_numero, d.poblacion, d.codigo_postal)
}

type ItemPedido struct {
	pieza    string
	cantidad int
}

type PedidoConDistancia struct {
	Pedido    Pedido
	Distancia float64
}

type Coordenadas struct {
	Lat float64 `json:"lat,string"`
	Lon float64 `json:"lon,string"`
}
