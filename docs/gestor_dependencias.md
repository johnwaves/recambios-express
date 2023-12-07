# Herramienta para la gestión de dependencias

Para la elección de un gestor de dependencias adecuado al proyecto se fijan los siguientes criterios:
- compatibilidad con el lenguaje,
- manejo de versiones, y
- soporte e integración con diversas herramientas de construcción.

A pesar de la existencia de varias herramientas que cumplen con estos requisitos, como pueden ser [Glide](https://glide.readthedocs.io/en/latest/), [Dep](https://golang.github.io/dep/docs/introduction.html) o [Govendor](https://pkg.go.dev/github.com/kardianos/govendor), la elegida para este proyecto es [**Go Modules**](https://go.dev/ref/mod#introduction). Este gestor asegura la compatibilidad con el Go, pues es la herramienta oficial que se introdujo a partir de la versión 1.11 de este lenguaje y que ofrece soporte por parte de la comunidad y del equipo de desarrollo de Go. Además, proporciona versionado de dependencias usando [SemVer](https://semver.org/lang/es/), aspecto clave para realizar el seguimiento de las versiones de forma sencilla y eficaz, y facilita la construcción del código integrando un soporte dedicado a comandos como `go build`.

En cuanto a las otras herramientas mencionadas, han sido en gran parte reemplazadas por Go Modules. Aunque en las primeras versiones presentaban ciertas ventajas, sus desventajas actuales clarifican el porqué del uso de la elegida. Por ejemplo, Dep fue la herramienta oficial de gestión de dependencias antes de la introducción de Go Modules, pero desde entonces ha sido descontinuada. Glide y Govendor también son herramientas de gestión de dependencias, pero no son tan completas o versátiles.



