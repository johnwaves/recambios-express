package main

import (
    "encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"sort"
)


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


const radioTierraKM = 6371

const (
	almacenLat = 37.5836344
	almacenLon = -1.7863288
)

var AlmacenCoordenadas = Coordenadas{
	Lat: almacenLat,
	Lon: almacenLon,
}

func ObtenerCoordenadas(direccion Direccion) (Coordenadas, error) {

	query := direccion.FormatDireccion()
	apiURL := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json&polygon=1&addressdetails=1", url.QueryEscape(query))

	resp, err := http.Get(apiURL)
	if err != nil {
		return Coordenadas{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Coordenadas{}, fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Coordenadas{}, err
	}

	var respuestaAPI []Coordenadas
	err = json.Unmarshal(body, &respuestaAPI)
	if err != nil {
		return Coordenadas{}, err
	}

	if len(respuestaAPI) == 0 {
		return Coordenadas{}, fmt.Errorf("no se encontraron coincidencias")
	}

	return respuestaAPI[0], nil
}

func GradosARadianes(grados float64) float64 {
	return grados * math.Pi / 180
}

func Haversine(coord1, coord2 Coordenadas) float64 {
	lat1 := GradosARadianes(coord1.Lat)
	lon1 := GradosARadianes(coord1.Lon)
	lat2 := GradosARadianes(coord2.Lat)
	lon2 := GradosARadianes(coord2.Lon)

	dLat := lat2 - lat1
	dLon := lon2 - lon1

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distancia := radioTierraKM * c
	return distancia
}

type Cliente struct {
	nombre_cliente    string
	direccion_entrega Direccion
}

func NewCliente(nombre_cliente string, direccion_entrega Direccion) *Cliente {
	return &Cliente{
		nombre_cliente:    nombre_cliente,
		direccion_entrega: direccion_entrega,
	}
}

func (c *Cliente) Nombre() string {
	return c.nombre_cliente
}

func (c *Cliente) FormatDireccionCliente() string {
	return c.direccion_entrega.FormatDireccion()
}

func (c *Cliente) HacerPedido(items []ItemPedido) *Pedido {
	pedido := NewPedido(items)
	pedido.asignarCliente(c)
	return pedido
}

type Pedido struct {
	cliente *Cliente
	items   []ItemPedido
}

func NewPedido(items []ItemPedido) *Pedido {
	return &Pedido{
		items: items,
	}
}

func (p *Pedido) asignarCliente(c *Cliente) {
	p.cliente = c
}

func (p *Pedido) Cliente() *Cliente {
	return p.cliente
}

const MaxPedidos int = 10

type Ruta struct {
	pedidos []PedidoConDistancia
}

func NewRuta() *Ruta {
	return &Ruta{}
}

func (r *Ruta) AddPedido(pedido *Pedido) {
	coordenadasDestino, _ := ObtenerCoordenadas(pedido.Cliente().direccion_entrega)
	distancia := Haversine(AlmacenCoordenadas, coordenadasDestino)
	fmt.Printf("Latitud: %f, Longitud: %f\n", coordenadasDestino.Lat, coordenadasDestino.Lon)
	fmt.Printf("Distancia: %f\n", distancia)
	if coordenadasDestino.Lat != 0.0 && coordenadasDestino.Lon != 0.0 {

		pedidoConDistancia := PedidoConDistancia{
			Pedido:    *pedido,
			Distancia: distancia,
		}

		r.pedidos = append(r.pedidos, pedidoConDistancia)
	}
	
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

    // Categorización de pedidos
    var pedidosCortos, pedidosMedios, pedidosLargos []PedidoConDistancia
    for _, pedido := range r.pedidos {
        if pedido.Distancia < distanciaMedia {
            pedidosCortos = append(pedidosCortos, pedido)
        } else if pedido.Distancia < distanciaMaxima * 0.75 {
            pedidosMedios = append(pedidosMedios, pedido)
        } else {
            pedidosLargos = append(pedidosLargos, pedido)
        }
    }

    // Ordenar los pedidos lejanos de mayor a menor distancia
    sort.Slice(pedidosLargos, func(i, j int) bool {
        return pedidosLargos[i].Distancia > pedidosLargos[j].Distancia
    })

    r.pedidos = []PedidoConDistancia{} // Vaciar el slice original para reordenarlo

    if len(pedidosCortos) == 0 && len(pedidosMedios) == 0 {
        // Solo hay pedidos lejanos
        r.pedidos = append(r.pedidos, pedidosLargos...)
    } else {
        // Intercalar pedidos
        i, j, k := 0, 0, 0
        for i < len(pedidosLargos) || j < len(pedidosCortos) || k < len(pedidosMedios) {
            // Agregar hasta dos pedidos largos si hay disponibles
            for count := 0; i < len(pedidosLargos) && count < 2; count++ {
                r.pedidos = append(r.pedidos, pedidosLargos[i])
                i++
            }

            // Agregar un pedido corto si hay disponible
            if j < len(pedidosCortos) {
                r.pedidos = append(r.pedidos, pedidosCortos[j])
                j++
            } else if k < len(pedidosMedios) { // Si no hay cortos, agregar un medio
                r.pedidos = append(r.pedidos, pedidosMedios[k])
                k++
            }
        }
    }
}


func (p *Pedido) Items() []ItemPedido {
    return p.items
}



func main() {
    // Crear un cliente
    direccion := NewDireccion("Calle Bartolo Escopeta", "Villarrobledo del Entresuelo", 12345)
    cliente := NewCliente("Juan Pérez", *direccion)

    // Añadir productos a un pedido
    items := []ItemPedido{
        {pieza: "Embrague", cantidad: 2},
        {pieza: "Pastillas de frenos", cantidad: 3},
    }
    pedido := cliente.HacerPedido(items)

    // Crear una ruta y añadir el pedido a la ruta
    ruta := NewRuta()
    ruta.AddPedido(pedido) // Corregido: quitado el manejo de error
    
    // **************************************************************************
    
    direccion = NewDireccion("Papa Juan Pablo II 7 1ºB", "Águilas", 30880)
    cliente = NewCliente("Antonia", *direccion)

    // Añadir productos a un pedido
    items = []ItemPedido{
        {pieza: "Limpiaparabrisas", cantidad: 7},
    }
    pedido = cliente.HacerPedido(items)
    ruta.AddPedido(pedido)
    
    // **************************************************************************
    
    direccion = NewDireccion("Sierra de la Pila 109", "Puerto de Mazarrón", 30860)
    cliente = NewCliente("Jerónimo", *direccion)

    // Añadir productos a un pedido
    items = []ItemPedido{
        {pieza: "Escobilla", cantidad: 1},
    }
    pedido = cliente.HacerPedido(items)
    ruta.AddPedido(pedido)
    
    // **************************************************************************
    
    direccion = NewDireccion("Plaza de Carruajes 10", "Lorca", 30800)
    cliente = NewCliente("David", *direccion)

    // Añadir productos a un pedido
    items = []ItemPedido{
        {pieza: "Válvula", cantidad: 1},
    }
    pedido = cliente.HacerPedido(items)
    ruta.AddPedido(pedido)
    
    // **************************************************************************
    
    direccion = NewDireccion("Camino de Las Casicas 17", "Las Casicas", 30890)
    cliente = NewCliente("Carmen", *direccion)

    // Añadir productos a un pedido
    items = []ItemPedido{
        {pieza: "Escobilla", cantidad: 1},
    }
    pedido = cliente.HacerPedido(items)
    ruta.AddPedido(pedido)
    
    // **************************************************************************
    
    direccion = NewDireccion("Calle San Francisco 6", "Puerto Lumbreras", 30890)
    cliente = NewCliente("Juanito", *direccion)

    // Añadir productos a un pedido
    items = []ItemPedido{
        {pieza: "Escobilla", cantidad: 1},
    }
    pedido = cliente.HacerPedido(items)
    ruta.AddPedido(pedido)
    
    // **************************************************************************
    
    direccion = NewDireccion("Alemania 39", "Puerto Lumbreras", 30890)
    cliente = NewCliente("Juana", *direccion)

    // Añadir productos a un pedido
    items = []ItemPedido{
        {pieza: "Escobilla", cantidad: 7},
    }
    pedido = cliente.HacerPedido(items)
    ruta.AddPedido(pedido)
    
    // **************************************************************************
    
    direccion = NewDireccion("La Salud 68", "La Hoya", 30816)
    cliente = NewCliente("Concepción", *direccion)

    // Añadir productos a un pedido
    items = []ItemPedido{
        {pieza: "Escobilla", cantidad: 1},
    }
    pedido = cliente.HacerPedido(items)
    ruta.AddPedido(pedido)
    
    // **************************************************************************
    
    direccion = NewDireccion("Mariana Pineda 30", "Alcantarilla", 30820)
    cliente = NewCliente("Concepción", *direccion)

    // Añadir productos a un pedido
    items = []ItemPedido{
        {pieza: "Escobilla", cantidad: 1},
    }
    pedido = cliente.HacerPedido(items)
    ruta.AddPedido(pedido)
    
    // **************************************************************************
    
    direccion = NewDireccion("Cuartelillo 2", "Totana", 30850)
    cliente = NewCliente("Concepción", *direccion)

    // Añadir productos a un pedido
    items = []ItemPedido{
        {pieza: "Escobilla", cantidad: 1},
    }
    pedido = cliente.HacerPedido(items)
    ruta.AddPedido(pedido)
    
    // **************************************************************************
    
    direccion = NewDireccion("Unamuno 3", "Las Torres de Cotillas", 30565)
    cliente = NewCliente("Concepción", *direccion)

    // Añadir productos a un pedido
    items = []ItemPedido{
        {pieza: "Escobilla", cantidad: 1},
    }
    pedido = cliente.HacerPedido(items)
    ruta.AddPedido(pedido)
    
    // **************************************************************************
    
    direccion = NewDireccion("Calle Aurora 3", "Puerto Lumbreras", 30890)
    cliente = NewCliente("Concepción", *direccion)

    // Añadir productos a un pedido
    items = []ItemPedido{
        {pieza: "Escobilla", cantidad: 1},
    }
    pedido = cliente.HacerPedido(items)
    ruta.AddPedido(pedido)
    
    // **************************************************************************

    // Imprimir los pedidos en la ruta
    fmt.Println("Pedidos en la ruta: \n")
    for _, pedidoConDistancia := range ruta.pedidos {
        fmt.Printf("Nombre del Cliente: %s\n", pedidoConDistancia.Pedido.Cliente().Nombre())
        fmt.Printf("Dirección de Entrega: %s\n", pedidoConDistancia.Pedido.Cliente().FormatDireccionCliente())

        fmt.Println("Items del Pedido:")
        for _, item := range pedidoConDistancia.Pedido.Items() { // Corregido: asumiendo GetItems() es un nuevo método en Pedido
            fmt.Printf("- %s, cantidad: %d\n", item.pieza, item.cantidad)
        }

        fmt.Printf("Distancia desde el almacén: %f km\n\n", pedidoConDistancia.Distancia)
        fmt.Printf("-----------------------------------------\n\n")
    }
    
    ruta.OrdenarPedidos()
    fmt.Printf("******   *******   *******  *************\n\n")
    // Imprimir los pedidos en la ruta después de ordenarlos
    fmt.Println("Pedidos en la ruta después de la ordenación: \n")
    for _, pedido := range ruta.pedidos {
        fmt.Printf("Nombre del Cliente: %s\n", pedido.Pedido.Cliente().Nombre())
        fmt.Printf("Dirección de Entrega: %s\n", pedido.Pedido.Cliente().FormatDireccionCliente())

        fmt.Println("Items del Pedido:")
        for _, item := range pedido.Pedido.Items() { // Asumiendo que existe un método Items en Pedido
            fmt.Printf("- %s, cantidad: %d\n", item.pieza, item.cantidad)
        }

        fmt.Printf("Distancia desde el almacén: %f km\n\n", pedido.Distancia)
        fmt.Printf("-----------------------------------------\n\n")
    }
    
    
}

