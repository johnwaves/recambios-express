# Algoritmo usado para la ordenación y el cálculo de la ruta de reparto óptima

El algoritmo propuesto para la ordenación de los pedidos sigue una lógica específica basada en la distancia entre la ubicación del almacén (punto desde donde se reparten pedidos) y la dirección donde se encuentra el cliente asociado a cada pedido. Cabe destacar que esta solución no se considera como única, simplemente se trata de la que mejor se aproxima a la resolución de este problema.

## Clasificación de pedidos
Los pedidos se dividen en tres categorías en función de su distancia al almacén:
- **Cortos**: distancia menor que la media.
- **Medios**: distancia mayor o igual a la media pero menor que el 75% de la distancia máxima.
- **Largos**: distancia igual o mayor que el 75% de la distancia máxima.

## Escala de prioridades
Si todos los pedidos son lejanos, se reparten en orden descendente por distancia. En cambio, si existe una mezcla de distancias, se procede a intercalar los pedidos.
- El objetido es equilibrar la atención entre pedidos lejanos y cercanos.
- Por cada dos pedidos lejanos, se intenta repartir un pedido cercano o corto. Esto ayuda a evitar que los pedidos cercanos queden para el final, especialmente en rutas donde predominan los pedidos lejanos.
- Si no hay pedidos cortos disponibles en ese momento, se reparte un pedido de distancia media.

Este patrón se repite, asegurando que se atiendan pedidos de todas las distancias de manera equilibrada.

## Rangos de distancia
Los rangos de distancia (corto, medio y largo) se calculan dinámicamente en función de las distancias de los pedidos actuales, y no están preestablecidos. La idea es adaptar los rangos a las condiciones específicas de cada conjunto de pedidos, lo que permite una mayor flexibilidad y adaptabilidad a diferentes situaciones. Para esta clasificación se considera:
- **Distancia media**. Se calcula la distancia media de todos los pedidos. Esta distancia media se utiliza como punto de referencia para diferenciar entre pedidos cortos y medios.
- **Distancia máxima**. Se identifica la distancia máxima entre todos los pedidos. Un porcentaje de esta distancia máxima (en este caso, el 75%) se usa para diferenciar entre pedidos medios y largos.
