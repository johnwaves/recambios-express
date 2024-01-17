# Herramienta para la elaboración y gestión de tests y pruebas

## Criterios
La elección de herramientas para las pruebas y los tests que se realizarán antes de entregar un PMV deben seguir la misma línea de los criterios establecidos para los gestores de dependencias y tareas. Estas herramientas, en concreto y ante todo, deben ofrecer compatibilidad con Go, permitir una integración sin problemas con las demás medios y ofrecer facilidad para su configuración.

Además, la herramienta de test debe:
- crear aserciones flexibles, rápidas y eficientes,
- si es necesario, agrupar las aserciones en subtests,
- planificar todos los test necesarios,
- incluir soporte para varios tipos de test, si se estima oportuno,
- generar informes o mostrar, simplemente, información sobre el resultado de los test,
- tener la capacidad de  automatizar los test para mejorar la eficiencia y la consistencia del proceso de prueba.

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