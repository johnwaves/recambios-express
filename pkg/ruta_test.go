package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRuta(t *testing.T) {
	ruta := NewRuta()

	assert.NotNil(t, ruta)
	assert.Empty(t, ruta.pedidos)
}

func TestErrorCoordenadas(t *testing.T) {
	ruta := NewRuta()
	pedidosIniciales := len(ruta.pedidos)

	direccion := NewDireccion("Calle Bartolo Escopeta", "Villarrobledo del Entresuelo", 12345)
	cliente, err := NewCliente("Manolo Escobar", *direccion)
	if err != nil {
		t.Logf("Error al crear el cliente: %v", err)
	}

	items := []ItemPedido{{pieza: "Ambientador con olor a pino", cantidad: 1}}
	pedido, err := cliente.HacerPedido(items)
	if err != nil {
		t.Logf("Error al hacer el pedido: %v", err)
	}

	err = ruta.AddPedido(pedido)
	if err != nil {
		t.Logf("Error al añadir el pedido a la ruta: %v", err)
	}

	assert.Equal(t, pedidosIniciales, len(ruta.pedidos), "Se esperaba que el pedido no se añadiera debido a una dirección no válida.")
}

func TestDistanciaAproximativa(t *testing.T) {
	ruta := NewRuta()

	cliente := struct {
		nombre                  string
		calle                   string
		ciudad                  string
		codigoPostal            int
		items                   []ItemPedido
		distanciaManualInferior float64
		distanciaManualSuperior float64
	}{
		nombre:                  "Manuel",
		calle:                   "Calle Aurora 3",
		ciudad:                  "Puerto Lumbreras",
		codigoPostal:            30890,
		items:                   []ItemPedido{{pieza: "Embrague", cantidad: 1}},
		distanciaManualInferior: 2.94,
		distanciaManualSuperior: 2.96,
	}

	direccion := NewDireccion(cliente.calle, cliente.ciudad, cliente.codigoPostal)
	nuevoCliente, err := NewCliente(cliente.nombre, *direccion)
	if err != nil {
		t.Fatalf("Error al crear el cliente: %v", err)
	}

	pedido, err := nuevoCliente.HacerPedido(cliente.items)
	if err != nil {
		t.Fatalf("Error al hacer el pedido: %v", err)
	}

	err = ruta.AddPedido(pedido)
	if err != nil {
		t.Fatalf("Error al añadir el pedido: %v", err)
	}

	distanciaReal := ruta.pedidos[len(ruta.pedidos)-1].Distancia

	t.Logf("Distancia real: %f", distanciaReal)

	assert.GreaterOrEqual(t, distanciaReal, cliente.distanciaManualInferior, "La distancia real es menor que la distancia manual inferior.")
	assert.LessOrEqual(t, distanciaReal, cliente.distanciaManualSuperior, "La distancia real es mayor que la distancia manual superior.")
}

func TestClienteNombre(t *testing.T) {
	direccion := NewDireccion("Calle Lope Gisbert 35", "Lorca", 30800)
	_, err := NewCliente("", *direccion)

	if err != nil {
		t.Logf("Error esperado: %v", err)
	} else {
		t.Errorf("Se esperaba un error cuando el nombre del cliente está vacío")
	}
}

func TestClienteDireccionCalle(t *testing.T) {
	direccion := NewDireccion("", "Lorca", 30800)
	_, err := NewCliente("Ángel", *direccion)

	if err != nil {
		t.Logf("Error esperado: %v", err)
	} else {
		t.Errorf("Se esperaba un error cuando la calle del cliente está vacía.")
	}
}

func TestClienteDireccionPoblacion(t *testing.T) {
	direccion := NewDireccion("Calle Lope Gisbert 35", "", 30800)
	_, err := NewCliente("Ángel", *direccion)

	if err != nil {
		t.Logf("Error esperado: %v", err)
	} else {
		t.Errorf("Se esperaba un error cuando la población del cliente está vacía.")
	}
}

func TestClienteDireccionCodigoPostal(t *testing.T) {
	direccion := NewDireccion("Calle Lope Gisbert 35", "Lorca", 0)
	_, err := NewCliente("Ángel", *direccion)

	if err != nil {
		t.Logf("Error esperado: %v", err)
	} else {
		t.Errorf("Se esperaba un error cuando el código postal del cliente es cero.")
	}
}

func TestClienteDireccionCodigoPostalNegativo(t *testing.T) {
	direccion := NewDireccion("Calle Lope Gisbert 35", "Lorca", -1)
	_, err := NewCliente("Ángel", *direccion)

	if err != nil {
		t.Logf("Error esperado: %v", err)
	} else {
		t.Errorf("Se esperaba un error cuando el código postal del cliente es negativo.")
	}
}

func TestAddPedido(t *testing.T) {
	ruta := NewRuta()

	clientes := []struct {
		nombre       string
		calle        string
		ciudad       string
		codigoPostal int
		items        []ItemPedido
	}{
		{
			nombre:       "Juan Pérez",
			calle:        "Mably 6",
			ciudad:       "Puerto Lumbreras",
			codigoPostal: 30890,
			items: []ItemPedido{
				{pieza: "Embrague", cantidad: 2},
				{pieza: "Pastillas de frenos", cantidad: 3}},
		},
		{
			nombre:       "Antonio Gómez",
			calle:        "Papa Juan Pablo II 7 1ºB",
			ciudad:       "Águilas",
			codigoPostal: 30880,
			items:        []ItemPedido{{pieza: "Limpiaparabrisas", cantidad: 7}},
		},
		{
			nombre:       "",
			calle:        "Avda. Pedro García Rubio 6",
			ciudad:       "Puerto Lumbreras",
			codigoPostal: 30890,
			items:        []ItemPedido{{pieza: "Escobilla", cantidad: 1}},
		},
		{
			nombre:       "M. Rajoy",
			calle:        "",
			ciudad:       "Puerto Lumbreras",
			codigoPostal: 30890,
			items:        []ItemPedido{{pieza: "Escobilla", cantidad: 1}},
		},
		{
			nombre:       "Ginés Artero",
			calle:        "La Rosa 8",
			ciudad:       "",
			codigoPostal: 30890,
			items:        []ItemPedido{{pieza: "Escobilla", cantidad: 1}},
		},
		{
			nombre:       "Gabriel Jiménez",
			calle:        "Antonio Martínez Garro 2",
			ciudad:       "Puerto Lumbreras",
			codigoPostal: 0,
			items:        []ItemPedido{{pieza: "Escobilla", cantidad: 1}},
		},
		{
			nombre:       "Eusebio Ortega",
			calle:        "Ramón y Cajal 45",
			ciudad:       "Puerto Lumbreras",
			codigoPostal: 30890,
			items:        []ItemPedido{{pieza: "", cantidad: 1}, {pieza: "Alfombrilla", cantidad: -2}},
		},
		{
			nombre:       "José Ruíz",
			calle:        "Calle Estrella 86",
			ciudad:       "Puerto Lumbreras",
			codigoPostal: 30890,
			items:        []ItemPedido{{pieza: "Escobilla", cantidad: 0}},
		},
	}

	var numPedidosTotales int = 8
	var pedidosEsperados int = 2
	var pedidosAñadidos int
	var pedidosInvalidos int

	for _, clienteData := range clientes {
		direccion := NewDireccion(clienteData.calle, clienteData.ciudad, clienteData.codigoPostal)
		nuevoCliente, err := NewCliente(clienteData.nombre, *direccion)
		if err != nil {
			t.Logf("Error al crear el cliente: %v", err)
			pedidosInvalidos++
		} else {
			pedido, err := nuevoCliente.HacerPedido(clienteData.items)
			if err != nil {
				t.Logf("Error al hacer el pedido: %v", err)
				pedidosInvalidos++
			} else {
				err = ruta.AddPedido(pedido)
				if err != nil {
					t.Logf("Error al añadir el pedido: %v", err)
					pedidosInvalidos++
				} else {
					pedidosAñadidos++
				}
			}
		}
	}

	t.Logf("Número esperado de pedidos: %d", numPedidosTotales)
	t.Logf("Número real de pedidos añadidos: %d", pedidosAñadidos)
	t.Logf("Número de pedidos inválidos: %d", pedidosInvalidos)

	for i, pedido := range ruta.pedidos {
		t.Logf("\nPedido %d: %+v\n", i+1, pedido)
	}

	assert.Equal(t, pedidosEsperados, pedidosAñadidos, "El número de pedidos añadidos debería ser igual al número total de pedidos válidos.")
}

var clientesOrdena = []struct {
	nombre       string
	calle        string
	ciudad       string
	codigoPostal int
	items        []ItemPedido
}{
	{
		nombre:       "Juan Pérez",
		calle:        "Mably 6",
		ciudad:       "Puerto Lumbreras",
		codigoPostal: 30890,
		items: []ItemPedido{
			{pieza: "Embrague", cantidad: 2},
			{pieza: "Pastillas de frenos", cantidad: 3}},
	},
	{
		nombre:       "Antonio Gómez",
		calle:        "Papa Juan Pablo II 7 1ºB",
		ciudad:       "Águilas",
		codigoPostal: 30880,
		items:        []ItemPedido{{pieza: "Limpiaparabrisas", cantidad: 7}},
	},
	{
		nombre:       "Manoli Egea",
		calle:        "Sierra de la Pila 109",
		ciudad:       "Puerto de Mazarrón",
		codigoPostal: 30860,
		items:        []ItemPedido{{pieza: "Escobilla", cantidad: 1}},
	},
	{
		nombre:       "David López",
		calle:        "Plaza de Carruajes 10",
		ciudad:       "Lorca",
		codigoPostal: 30800,
		items:        []ItemPedido{{pieza: "Válvula", cantidad: 1}},
	},
	{
		nombre:       "Carmen Padilla",
		calle:        "Camino de Las Casicas 17",
		ciudad:       "Las Casicas",
		codigoPostal: 30890,
		items:        []ItemPedido{{pieza: "Escobilla", cantidad: 1}},
	},
	{
		nombre:       "Juanito Valderrama",
		calle:        "Calle San Francisco 6",
		ciudad:       "Puerto Lumbreras",
		codigoPostal: 30890,
		items:        []ItemPedido{{pieza: "Escobilla", cantidad: 1}},
	},
	{
		nombre:       "Andrea Muñoz",
		calle:        "Alemania 39",
		ciudad:       "Puerto Lumbreras",
		codigoPostal: 30890,
		items:        []ItemPedido{{pieza: "Escobilla", cantidad: 7}},
	},
	{
		nombre:       "Concepción Velasco",
		calle:        "La Salud 68",
		ciudad:       "La Hoya",
		codigoPostal: 30816,
		items:        []ItemPedido{{pieza: "Escobilla", cantidad: 1}},
	},
	{
		nombre:       "Ana Martínez",
		calle:        "Mariana Pineda 30",
		ciudad:       "Alcantarilla",
		codigoPostal: 30820,
		items:        []ItemPedido{{pieza: "Escobilla", cantidad: 1}},
	},
	{
		nombre:       "María José Miñarro",
		calle:        "Cuartelillo 2",
		ciudad:       "Totana",
		codigoPostal: 30850,
		items:        []ItemPedido{{pieza: "Escobilla", cantidad: 1}},
	},
	{
		nombre:       "Sebastián Correas",
		calle:        "Unamuno 3",
		ciudad:       "Las Torres de Cotillas",
		codigoPostal: 30565,
		items:        []ItemPedido{{pieza: "Escobilla", cantidad: 1}},
	},
	{
		nombre:       "Gregorio Navarro",
		calle:        "Calle Ojos 5",
		ciudad:       "Puerto Lumbreras",
		codigoPostal: 30890,
		items:        []ItemPedido{{pieza: "Escobilla", cantidad: 1}},
	},
}

func TestDistancias(t *testing.T) {
	ruta := NewRuta()

	for _, clienteData := range clientesOrdena {
		direccion := NewDireccion(clienteData.calle, clienteData.ciudad, clienteData.codigoPostal)
		nuevoCliente, err := NewCliente(clienteData.nombre, *direccion)
		if err != nil {
			t.Fatalf("Error al crear el cliente: %v", err)
		}

		pedido, err := nuevoCliente.HacerPedido(clienteData.items)
		if err != nil {
			t.Fatalf("Error al hacer el pedido: %v", err)
		}

		err = ruta.AddPedido(pedido)
		if err != nil {
			t.Fatalf("Error al añadir el pedido: %v", err)
		}
	}

	ruta.OrdenarPedidos()

	if ruta.CantidadActual() > 0 {
		distanciaMaxima := ruta.distanciaMaxima()
		for i, pedido := range ruta.pedidos {
			if i > 0 && i < len(ruta.pedidos)-1 {
				assert.LessOrEqual(t, pedido.Distancia, distanciaMaxima, "Un pedido tiene una distancia mayor a la máxima esperada.")
			}
		}
	}
}

func TestOrdenarPedidos(t *testing.T) {
	ruta := NewRuta()

	for _, clienteData := range clientesOrdena {
		direccion := NewDireccion(clienteData.calle, clienteData.ciudad, clienteData.codigoPostal)
		nuevoCliente, err := NewCliente(clienteData.nombre, *direccion)
		if err != nil {
			t.Fatalf("Error al crear el cliente: %v", err)
		}

		pedido, err := nuevoCliente.HacerPedido(clienteData.items)
		if err != nil {
			t.Fatalf("Error al hacer el pedido: %v", err)
		}

		err = ruta.AddPedido(pedido)
		if err != nil {
			t.Fatalf("Error al añadir el pedido: %v", err)
		}
	}

	ruta.OrdenarPedidos()

	cuentaLargosConsecutivos := 0
	for i, pedido := range ruta.pedidos {
		esPedidoLargo := pedido.Distancia >= ruta.distanciaMaxima()*0.75

		if esPedidoLargo {
			cuentaLargosConsecutivos++
		} else {
			assert.LessOrEqual(t, cuentaLargosConsecutivos, 2, "Hay más de dos pedidos largos consecutivos.")
			cuentaLargosConsecutivos = 0
		}

		if esPedidoLargo && i >= len(ruta.pedidos)-3 {
			assert.Fail(t, "Hay un pedido largo en las últimas posiciones.")
		}
	}

	for i, pedido := range ruta.pedidos {
		t.Logf("\nPedido %d: %+v\n", i+1, pedido)
	}
}
