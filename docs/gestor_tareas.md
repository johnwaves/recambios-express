# Herramienta para la gestión de tareas

Los criterios establecidos para la selección de una herramienta adecuada para la gestión de las tareas en el proyecto son:
- compatibilidad, o en todo caso adaptabilidad, con el lenguaje del proyecto,
- soporte para pruebas,
- manuales y documentación que no resulten obsoletos.

Entre las diversas opciones disponibles se ha elegido **Task**. Además del soporte activo por parte de la comunidad y el equipo desarrollador, el uso de este gestor de tareas resulta ser más simple que otros, como es el caso de [Make](https://www.gnu.org/software/make/manual/make.html) o [Mage](https://docs.mage.ai/introduction/overview), dado que está escrito en Go y su sintaxis es sencilla. Su configuración se centra en un archivo con formato `YAML` donde se definen las tareas a realizar. Otra característica de Task es que también permite definir variables de entorno para las pruebas y los test que se ejecutarán antes del despliegue del proyecto, lo que posibilita establecer configuraciones específicas para cada una de ellas. Un motivo adicional para usar este gestor es el interés personal por explorar y aprender a utilizar herramientas con las que no se ha tratado anteriormente en el grado.

En cuanto a los demás gestores mencionados, Make es uno de los más populares, destacando por su versatilidad y alta eficiencia al comprobar únicamente aquellos archivos modificados entre ejecuciones de tareas, ofreciendo así un gran ahorro en recursos. En el caso de Mage, también está escrito en Go, aspecto que proporciona un nivel de integración elevado. Dispone de funciones (objetivos) que representan diferentes tareas que se pueden definir para compilar el proyecto, ejecutar pruebas o generar documentación, entre otros. Sin embargo, ambos gestones se han descartado por no ofrecer una sintaxis tan legible como la de Task, lo que se traduce en resultar más difícil la redacción de sus correspondientes scripts.
