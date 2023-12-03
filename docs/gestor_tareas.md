# Herramienta para la gestión de tareas

Como herramienta de construcción, compilación y automatización de tareas se ha elegido **Task**, pues sigue los [criterios](https://github.com/johnwaves/recambios-express/issues/44) especificados.

Además del soporte activo por parte de la comunidad y el equipo desarrollador, el uso de este gestor de tareas resulta ser más simple que otros, como es el caso de Make o Mage, dado que está escrito en Go y su sintaxis es sencilla. Su configuración se centra en un archivo con formato `YAML` donde se definen las tareas a realizar. Task también permite definir variables de entorno para las pruebas y los test que se ejecutarán antes del despliegue del proyecto, lo que posibilita establecer configuraciones específicas para cada una de ellas. Otro motivo para usar este gestor es el interés personal por explorar y aprender a utilizar nuevas herramientas con las que no se ha tratado anteriormente en el grado.

Quedan, por tanto, descartados Make y Mage por no ofrecer una sintaxis tan legible como la de Task ni mayor facilidad en la redacción de scripts.
