package pkg

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
