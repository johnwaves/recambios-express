package pkg

import (
	"errors"
	"fmt"
	"strings"
)

type Pedido struct {
	cliente *Cliente
	items   []ItemPedido
}

func NewPedido(items []ItemPedido) (*Pedido, error) {
	if len(items) == 0 {
		return nil, errors.New("el pedido debe contener al menos un ítem")
	}

	var errorMessages []string

	for i, item := range items {
		if item.pieza == "" {
			errorMessages = append(errorMessages, fmt.Sprintf("el nombre de la pieza en la posición %d no puede estar vacío.", i+1))
		}
		if item.cantidad <= 0 {
			errorMessages = append(errorMessages, fmt.Sprintf("la cantidad de la pieza '%s' en la posición %d debe ser mayor que cero.", item.pieza, i+1))
		}
	}

	if len(errorMessages) > 0 {
		return nil, errors.New(strings.Join(errorMessages, "; "))
	}

	return &Pedido{
		items: items,
	}, nil
}

func (p *Pedido) asignarCliente(c *Cliente) {
	p.cliente = c
}

func (p *Pedido) Cliente() *Cliente {
	return p.cliente
}

func (p *Pedido) Items() []ItemPedido {
	return p.items
}
