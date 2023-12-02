# Herramienta para la gestión de dependencias

A pesar de la existencia de varias herramientas capaces de gestionar dependencias en proyectos
con Go, como pueden ser Glide, Dep o Govendor, la elegida para este proyecto es **Go Modules**.

La elección se ha hecho en base a los criterios especificados en este [issue](https://github.com/johnwaves/recambios-express/issues/43), dado que:
- Go Modules asegura la compatibilidad con el lenguaje, pues es la herramienta oficial que
se introdujo a partir de la versión 1.11 del mismo y que ofrece soporte por parte de la comunidad y del equipo de desarrollo de Go,
- proporciona versionado de dependencias usando Versionado Semántico [SemVer](https://semver.org/lang/es/), aspecto clave para realizar el seguimiento de las versiones de forma sencilla y eficaz, y
- facilita la construcción del código integrando soporte para comandos como `go build`.

En cuanto a las otras herramientas mencionadas, han sido en gran parte reemplazadas por Go Modules. Aunque presentan diversas ventajas, las desventajas clarifican el porqué del uso de la que se ha elegido. Por ejemplo, Dep fue la herramienta oficial de gestión de dependencias antes de la introducción de Go Modules, pero desde entonces ha sido descontinuada. Glide y Govendor también son herramientas de gestión de dependencias, pero no son tan completas o versátiles como Go Modules.




