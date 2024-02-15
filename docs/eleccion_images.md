# Elección de imagen para el contenedor del proyecto

El uso de imágenes de contenedores se ha establecido como un enfoque estándar en la actualidad para generar despliegues uniformes y reproducibles, aplicables a una amplia gama de aplicaciones y situaciones, que van desde pruebas hasta la implementación de microservicios.

## Criterios
Los criterios que se seguirán para la elección de la imagen ideal se establecen a continuación: 

- **Tamaño y rendimiento**. Optar por una imagen ligera es crucial para garantizar un arranque rápido. Si se eligen imágenes más pequeñas, aunque es posible que ofrezcan menos funcionalidades, puede resultar un compromiso viable cuando el proyecto no requiere de una extensa gama de bibliotecas. En este caso, se optará por una imagen ligera pero lo suficientemente completa.

- **Seguridad y actualizaciones**. Es importante seleccionar una imagen que esté actualizada y cuente con un soporte robusto. Las versiones más recientes suelen corregir fallos de seguridad identificados en versiones anteriores.

- **Compatibilidad con Golang**. Resulta ideal que la imagen venga con Go ya instalado para facilitar el proceso y evitar conflictos de versión. Además, utilizar versiones recientes permite aprovechar nuevas funcionalidades y estándares en el desarrollo, fomentando, ante todo, el progreso.

## Imágenes candidatas
En la siguiente lista se citan algunas de las imágenes consideradas como candidatas para ser usadas para este proyecto.

### [Ubuntu](https://hub.docker.com/_/ubuntu) o [Debian](https://hub.docker.com/_/debian):
Debian y Ubuntu ofrecen un balance entre funcionalidad y rendimiento. Son adecuados para proyectos que pueden requerir una gama más amplia de bibliotecas y dependencias. Ofrecen actualizaciones regulares y mejoras frecuentes de seguridad. Sin embargo, en ambos sistemas en necesario instalar Go y las dependencias de este proyecto, y esto, junto a las herramientas adicionales que incluyen, podría suponer un aumento considerable del tamaño de la imagen a instalar y del tiempo de instalación, si lo que se busca es agilidad en el desarrollo.

### [Golang con Alpine](https://hub.docker.com/_/alpine)
Esta es una de las opciones que ofrece un entorno minimalista y eficiente para desplegar servicios ligeros escritos en Go, dado que incorpora un compilador y algunas herramientas para desarrollar y ejecutar aplicaciones escritas en este lenguaje (como pueden ser `Go fmt`, `Go test` o `Go build`). Pese a disponer de estas ventajas, cabe mencionar que Alpine tiene un repositorio aún creciente y el número de paquetes disponibles es menor en comparación con otras distribuciones más grandes. Esto puede significar que ciertas herramientas o bibliotecas no estén disponibles o requieran una compilación manual.

### [Bitnami con Golang](https://hub.docker.com/r/bitnami/golang)
Bitnami es una de las imágenes que simplifica el proceso de despliegue, pues viene con herramientas preconfiguradas para el entorno Go y con actualizaciones muy frecuentes. De hecho, durante la redacción de este documento se ha publicado una actualización y aplicado un nuevo parche de seguridad. Además, el entorno de desarrollo no está cargado con demasiadas herramientas que puedan resultar innecesarias.

### [Debian (versión slim)](https://hub.docker.com/_/debian)
Esta variante de Debian se presenta como ligeramente más pequeña. Ofrece un sistema operativo base, sin herramientas añadidas, pudiendo nosotros instalar lo necesario para el proyecto. En cuanto a mantenimiento, las actualizaciones son constantes y ofrece un grado alto de seguridad. 

## Conclusión 
Tras evaluar cada una de las opciones disponibles y basando esta conclusión en los criterios establecidos, la imagen que se utilizará para llevar a cabo las pruebas de este proyecto será **Debian en su versión slim**.

Como se trata de una imagen que no dispone del entorno de desarrollo de Golang, la instalación del mismo y de la propia imagen en el contenedor se realiza en dos fases: construcción e instalación (fase final). En la fase de construcción se prepara el entorno a partir de la imagen oficial de Golang, y se instala `task`. En la fase de instalación se crea la imagen final, Debian (stable-slim) y se copia todo el entorno de Go, pero sin añadir dependencias adicionales. Cabe destacar que es necesario utilizar certificados SSL para poder permitir conexiones seguras cuando se gestionen las dependencias y establecer el PATH para incluir los directorios de los binarios de Go y `task`.

Esto, por tanto, supone una opción ideal para no incluir componentes innecesarios que pueden venir con vulnerabilidades o abundantes puntos débiles. Es por ello por lo que la imagen resultante está optimizada para la producción y el entorno utilizado es uno limpio.