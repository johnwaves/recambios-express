# Milestones

En este documento se recogen los hitos o _milestones_ del proyecto.

## [M-0](https://github.com/johnwaves/recambios-express/milestone/1): Modelo del problema. 
Este milestone tiene como objetivo generar un modelo del problema a partir de las _user-stories_, identificando las entidades participantes y los elementos clave. Lo que se entregará será una versión codificada donde se definirán las estructuras de datos para dichas entidades y las relaciones entre ellas.
La viabilidad se garantizará con la aprobación del gestor de producto o _product manager_.

## [M-1](): Creación de pedidos.
El propósito de este hito es la creación de pedidos. Los talleres deben poder realizar pedidos para el almacén y este, a su vez, obtener un listado de los pedidos realizados. Será viable cuando se consiga lo indicado.

## [M-2](): Asignación de pedidos.
Se precisa realizar la asignación de pedidos a los repartidores implementando una restricción. Se debe tener en cuenta el horario de trabajo de cada repartidor para entregar los pedidos, pues el horario de reparto no lo puede exceder. Se entregará un módulo al que se le añadirá esta restricción una vez superadas las pruebas oportunas.

## [M-3](): Estimación de la fecha de entrega.
Obtenidos todos los detalles sobre un pedido y su repartidor, se debe establecer una fecha de entrega aproximada en función del volumen de pedidos ya recibidos y del tiempo de desplazamiento para cada uno de ellos. Para implementar este aspecto se ha de tener en cuenta que si no se llega a repartir un determinado pedido un día, se le asignará máxima prioridad para el día siguiente y será el primero en ser entregado.