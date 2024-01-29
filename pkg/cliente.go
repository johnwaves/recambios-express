package pkg

import "errors"

type Cliente struct {
	nombre_cliente    string
	direccion_entrega Direccion
}

func NewCliente(nombre_cliente string, direccion_entrega Direccion) (*Cliente, error) {

	if nombre_cliente == "" {
		return nil, errors.New("el nombre del cliente no puede estar vacío.")
	}

	errorMsg := ""
	if direccion_entrega.calle_numero == "" {
		errorMsg += "la calle del cliente no puede estar vacía. "
	}

	if direccion_entrega.poblacion == "" {
		errorMsg += "la población del cliente no puede estar vacía. "
	}

	if direccion_entrega.codigo_postal == 0 || direccion_entrega.codigo_postal < 0 {
		errorMsg += "el código postal del cliente no puede ser cero o negativo."
	}

	if errorMsg != "" {
		return nil, errors.New("ERROR en la dirección del cliente: " + errorMsg)
	}

	return &Cliente{
		nombre_cliente:    nombre_cliente,
		direccion_entrega: direccion_entrega,
	}, nil
}

func (c *Cliente) Nombre() string {
	return c.nombre_cliente
}

func (c *Cliente) FormatDireccionCliente() string {
	return c.direccion_entrega.FormatDireccion()
}

func (c *Cliente) HacerPedido(items []ItemPedido) (*Pedido, error) {
	pedido, err := NewPedido(items)
	if err != nil {
		return nil, err
	}

	pedido.asignarCliente(c)

	return pedido, nil
}
