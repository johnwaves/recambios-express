# Herramienta para la elaboración y gestión de tests y pruebas

## Criterios y requisitos
Al elegir las herramientas adecuadas para la elaboración y gestión de tests y pruebas en Go, es crucial establecer tanto requisitos básicos como criterios de selección detallados. Estos facilitarán la selección de herramientas más eficientes y compatibles con las necesidades de desarrollo.

### Requisitos básicos
- **Compatibilidad con el lenguaje Go y su toolchain actual**. Las herramientas deben integrarse sin problemas con el entorno de desarrollo Go existente.
- **Facilidad de configuración**. Deben ser fáciles de configurar y utilizar.

### Criterios de selección
- **Mantenimiento y actualizaciones regulares**. Las herramientas deben recibir actualizaciones y mantenimiento regulares para garantizar que no quedan obsoletas. Por ello, también convendría ser una herramienta lo suficientemente popular y usada como para tener una comunidad activa detrás para que se encuentre en continuo desarrollo.

- **Rendimiento de aserciones y pruebas**. Deben permitir la creación de aserciones rápidas, eficientes y flexibles. La herramienta en cuestión deberá otorgar la posibilidad de lanzar tests relacionados con cada función del código.

- **Simplicidad en la generación y consulta de informes y resultados**. La herramienta elegida será ideal si los informes correspondientes a las pruebas ejecutadas resultan ser simples de interpretar y comprobar. Si falla algún test realizado, debería generar la suficiente información como para saber qué se necesitaría cambiar en relación con la funcionalidad testeada, sin tener que hacer una búsqueda exhaustiva del error en todo el código proporcionado. 

- **Soporte para diversos tipos de pruebas**. Deben ser versátiles para soportar diferentes tipos de pruebas según sea necesario. Dicho de otra forma, la biblioteca de aserciones debe ser completa en cuanto a pruebas, ofreciendo diferentes maneras de comprobar y analizar el código escrito.

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
Personalmente, la combinación de Testify y el módulo test de Go representa una solución robusta y coherente para las pruebas durante el desarrollo de este proyecto. Haciendo mención a los criterios establecidos, Testify resulta ser una de las herramientas más [populares](https://trends.google.com/trends/explore?date=today%205-y&q=golang%20testify,golang%20goconvey,golang%20ginkgo,golang%20httpexpect,golang%20gomega). Esto, y la extensa comunidad que tiene detrás, hace que se permanezca al día en cuanto a nivel de mantenimiento se refiere. Por otra parte, Testify ofrece también una amplia gama de aserciones para realizar pruebas de forma más exhaustiva, aspecto que se ha comprobado consultando la [documentación oficial](https://pkg.go.dev/github.com/stretchr/testify/assert) de la biblioteca. 

Por último, esta elección ofrece simplicidad necesaria para garantizar las mejores prácticas de este lenguaje, asegurando un proceso de testing que es tanto eficiente como efectivo. La integración de estas herramientas con el entorno de desarrollo y su [configuración sencilla](https://github.com/stretchr/testify?tab=readme-ov-file#installation) cumplen con los criterios establecidos para la selección de herramientas.
