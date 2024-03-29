# Recambios Express
Este repositorio contiene el proyecto que se realiza en la asignatura Infraestructura Virtual, perteneciente al Grado en Ingeniería Informática de la Universidad de Granada (curso 23/24). Su desarrollo tiene como base el modelo DevOps, el cual establece una serie de pasos a seguir para alcanzar los objetivos que se plantean y así lograr entregar un producto siguiendo las mejores prácticas. Más información [aquí](https://github.com/JJ/IV#material-docente-para-la-asignatura-infraestructura-virtual).

## Problema a resolver
Mi empresa se dedica a la venta de recambios para automóviles. Al tener únicamente dos repartidores, existen algunos pedidos al final de la jornada que no se llegan a entregar. Esto sucede porque se priorizan los pedidos de los talleres (clientes) que más lejos se encuentran desde la ubicación del almacén, pues se emplea más tiempo en los trayectos de ida y vuelta. Los pedidos que por cercanía se pueden repartir el mismo día quedan a la espera hasta que se entregan los otros. Todo ello conlleva una demora de tiempo para los pedidos restantes, da lugar a una disminución de las ventas y emporan las reseñas. Es por ello por lo que la asignación de pedidos para cada repartidor no resulta ser equitativa.

Los pedidos deben ser asignados de la forma más eficiente posible a cada repartidor, sin priorizar únicamente aquellos que se entregan más lejos del almacén. También se deben tener en cuenta los horarios de trabajo de cada repartidor y el horario de apertura del almacén para establecer los intervalos de reparto y los pedidos a entregar.

## Documentación sobre el problema
A continuación se detalla la información referente al punto de partida del proyecto:
- [Historias de usuario](https://github.com/johnwaves/recambios-express/blob/main/docs/user-stories.md)
- [Hitos](https://github.com/johnwaves/recambios-express/blob/main/docs/milestones.md)

## Herramientas
### Gestor de dependencias
**Go Modules** es la herramienta encargada de gestionar las dependencias del proyecto. [Aquí](https://github.com/johnwaves/recambios-express/blob/Objetivo-3/docs/gestor_dependencias.md) se detallan los motivos de esta elección.

### Gestor de tareas
Para automatizar las tareas se utiliza **Task**. [Aquí](https://github.com/johnwaves/recambios-express/blob/Objetivo-3/docs/gestor_tareas.md) se detallan los motivos de esta elección. Las órdenes disponibles son:
- `task installdeps`: instala todas las dependencias que son requeridas en el proyecto.
- `task updatedeps`: actualiza las dependencias instaladas.
- `task check`: comprueba la sintaxis y señala los posibles errores que puedan existir.
- `task build`: compila el proyecto.
- `task test`: ejecuta las pruebas unitarias desarrolladas para comprobar la funcionalidad del código implementado en el proyecto.

### Herramienta de pruebas
**Testify** y **Go Test** son las herramientas elegidas para generar y ejecutar pruebas y excepciones que comprueban la funcionalidad del código implementado. [Aquí](https://github.com/johnwaves/recambios-express/blob/Objetivo-4/docs/herramientas_test.md) se puede consultar más información acerca de los criterios de elección.

## Contenedor para pruebas

### Imagen para el contenedor de pruebas
**Debian (versión slim)** es la imagen seleccionada para la ejecución de los tests unitarios sobre este proyecto. Tanto los criterios seguidos como las demás imágenes candidatas se detallan [aquí](https://github.com/johnwaves/recambios-express/blob/Objetivo-5/docs/eleccion_images.md).

### Pruebas
Este proyecto se puede probar en Docker con los siguientes comandos:
- Para construir el contenedor:
```
docker build -t johnwaves1/recambios-express .
```

- Para ejecutar el contenedor:
```
docker run -u 1001 -t -v `pwd`:/app/test johnwaves1/recambios-express
```

## Progreso del proyecto
Se ha avanzado hasta el [Objetivo 5](http://jj.github.io/IV/documentos/proyecto/5.Docker) de la asignatura.
Las tareas realizadas en objetivos anteriores se pueden consultar accediendo a las diferentes ramas de este repositorio.

## Configuración de GitHub
Para verificar la configuración del repositorio, pulsa [aquí](https://github.com/johnwaves/recambios-express/blob/main/docs/git-config.png).

## Más información
**NOTA:** Conforme vaya avanzando la asignatura esta descripción se irá actualizando. Asimismo, apareceán nuevas secciones y detalles.

