# Milestones

En este documento se recogen los hitos o _milestones_ del proyecto.

## [M-0](https://github.com/johnwaves/recambios-express/milestone/1): Modelo del problema. 
Este milestone tiene como objetivo generar un modelo del problema a partir de las _user-stories_, identificando las entidades participantes y los elementos clave. Lo que se entregará será una versión codificada donde se definirán las estructuras de datos para dichas entidades y las relaciones entre ellas.
La viabilidad se garantizará con la aprobación del gestor de producto o _product manager_.

## [M-1](https://github.com/johnwaves/recambios-express/milestone/2): Creación y organización de pedidos.
El propósito de este hito es la creación y organización de pedidos. Los talleres deben poder realizar pedidos al almacén y este, una vez recibidos, indicará de forma óptima mediante un algoritmo cuáles tienen prioridad y, por tanto, se deben entregar primero durante el día en curso. El entregable será un módulo que incluya el código con las implementaciones mencionadas y será viable si supera las pruebas correspondientes. 

## [M-2](https://github.com/johnwaves/recambios-express/milestone/3): Asignación y planificación de pedidos.
Para este hito se precisa realizar la asignación de pedidos a los repartidores en función de una restricción. Se debe tener en cuenta el horario de trabajo de cada repartidor para poder entregar los pedidos, pues el horario de reparto no lo puede exceder. Por tanto, la planificación diaria reflejará los pedidos a repartir en un determinado día y qué repartidor se hará cargo de cada uno. Se entregará un módulo donde vendrá el código para estas funcionalidades y, nuevamente, podrá ser viable siempre y cuando realice con éxito las pruebas.

## [M-3](https://github.com/johnwaves/recambios-express/milestone/4): Estimación de la fecha de entrega.
Obtenidos todos los detalles sobre un pedido y su repartidor, se debe establecer una fecha de entrega aproximada en función del volumen de pedidos ya recibidos y del tiempo de desplazamiento para cada uno de ellos. Para implementar este aspecto se ha de tener en cuenta que si no se llega a repartir un determinado pedido un día, se le asignará máxima prioridad para el día siguiente y será el primero en ser entregado. Lo que se proporcionará será un módulo con el código que realice esta estimación una vez superadas todas las comprobocaciones.
