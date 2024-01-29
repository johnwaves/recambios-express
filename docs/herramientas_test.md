# Herramienta para la elaboración y gestión de tests y pruebas

## Criterios y requisitos
Al elegir las herramientas adecuadas para la elaboración y gestión de tests y pruebas en Go, es crucial establecer tanto requisitos básicos como criterios de selección detallados. Estos facilitarán la selección de herramientas más eficientes y compatibles con las necesidades de desarrollo.

### Requisitos básicos
- **Compatibilidad con el lenguaje Go y su toolchain actual**. Las herramientas deben integrarse sin problemas con el entorno de desarrollo Go existente.
- **Facilidad de configuración**. Deben ser fáciles de configurar y utilizar.

### Criterios de selección
- **Mantenimiento y actualizaciones regulares**. Las herramientas deben recibir actualizaciones y mantenimiento regulares para garantizar que no quedan obsoletas.
- **Rendimiento de aserciones y pruebas**. Deben permitir la creación de aserciones rápidas, eficientes y flexibles.
- **Generación de informes**. Otra cuestión de vital importancia es la capacidad que deben tener para generar informes o mostrar información relevante sobre los resultados de los tests.
- **Soporte para diversos tipos de pruebas**. Deben ser versátiles para soportar diferentes tipos de pruebas según sea necesario.
- **Automatización de pruebas**. Por último, el tener la capacidad para automatizar los tests mejorará la eficiencia y consistencia del proceso.

## Listado de herramientas para la elección
Considerando estos criterios, se evalúan múltiples opciones.

### Bibliotecas de aserciones
En Go, las aserciones [no forman parte del estándar](https://stackoverflow.com/questions/45683134/assertion-in-non-test-function-in-go), por lo que es común recurrir a bibliotecas externas. Las más destacadas son:

- [**Testify**](https://github.com/stretchr/testify). Se trata de una biblioteca muy popular que ofrece un conjunto de aserciones y mocks comunes. Se integra bien con la biblioteca estándar de Go y es conocida por su versatilidad y facilidad de uso. 

- [**Gomega**](https://github.com/onsi/gomega). Esta biblioteca funciona particularmente bien con el framework BDD Ginkgo, aunque también puede adaptarse a otros contextos. Presenta una manera elegante de escribir aserciones en Go, con un enfoque en la legibilidad y expresividad del código​​.

- [**Go-cmp**](https://github.com/google/go-cmp). Mencionada en la lista de recursos [Awesome Go](https://awesome-go.com/testing/), se utiliza para comparar valores en Go durante las pruebas, con una funcionalidad detallada y flexible que va más allá de las comparaciones básicas disponibles en la biblioteca estándar de Go.

### Test runners o frameworks
- [**Módulo test de Go**](https://go.dev/doc/tutorial/add-a-test). Integrado en la biblioteca estándar de Go, `go test` es una opción muy utilizada por su simplicidad y eficacia. Ofrece un enfoque directo para la ejecución de pruebas sin la necesidad de dependencias externas.

- [**Ginkgo**](https://github.com/onsi/ginkgo). Este es un framework BDD para Go ampliamente reconocido y que se utiliza en combinación con Gomega para una experiencia de testing más expresiva​.

​- [**GoConvey**](https://github.com/smartystreets/goconvey). También mencionado en la [Awesome Go](https://awesome-go.com/testing/), y se trata de un framework de pruebas BDD con una interfaz web para visualizar todas las pruebas en tiempo real.

## Decisión
Desde mi punto de vista, la combinación de Testify y el módulo test de Go representa una solución robusta y coherente para las pruebas durante el desarrollo de este proyecto. Tanto Testify como `go test` garantizan una ejecución de tests eficiente y bien integrada con el ecosistema de Go. 

Esta elección ofrece simplicidad necesaria para garantizar las mejores prácticas de este lenguaje, asegurando un proceso de testing que es tanto eficiente como efectivo. La integración de estas herramientas con el entorno de desarrollo y su configuración sencilla cumplen con los criterios establecidos para la selección de herramientas.
