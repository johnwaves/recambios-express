# Milestones

En este documento se recogen los hitos o _milestones_ del proyecto.

## [M-0](): Modelo del problema. 
Este milestone tiene como objetivo generar un modelo del problema a partir de las _user-stories_, identificando las entidades participantes y los elementos clave. Por ello, el entregable será un módulo que permita a los clientes realizar pedidos al almacén y asignarlos a los repartidores. La viabilidad la garantizarán los test previos a realizar.

## [M-1](): Implementación de rectricciones temporales.
Es preciso implementar una restricción de carácter temporal para entregar los pedidos, pues el horario de reparto no puede exceder el horario de trabajo establecido en la jornada laboral de los repartidores. Se entregará un módulo al que se le añadirá esta restricción una vez superadas las pruebas oportunas.

## [M-2](): Estimación de la fecha de entrega.
Obtenidos todos los detalles sobre un pedido y su repartidor, se debe establecer una fecha de entrega aproximada en función del volumen de pedidos ya recibidos y del tiempo de desplazamiento para cada uno de ellos. Para implementar este aspecto se ha de tener en cuenta que si no se llega a repartir un determinado pedido un día, se le asignará máxima prioridad para el día siguiente y será el primero en ser entregado.