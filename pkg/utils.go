package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
)

const radioTierraKM = 6371

func ObtenerCoordenadas(direccion Direccion) (Coordenadas, error) {
	query := fmt.Sprintf("%s, %s, %d", direccion.calle_numero, direccion.poblacion, direccion.codigo_postal)

	apiURL := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json&polygon=1&addressdetails=1", url.QueryEscape(query))

	resp, err := http.Get(apiURL)
	if err != nil {
		return Coordenadas{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Coordenadas{}, fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Coordenadas{}, err
	}

	var respuestaAPI []Coordenadas
	err = json.Unmarshal(body, &respuestaAPI)
	if err != nil {
		return Coordenadas{}, err
	}

	if len(respuestaAPI) == 0 {
		return Coordenadas{}, fmt.Errorf("no se encontraron coincidencias")
	}

	return respuestaAPI[0], nil
}

func GradosARadianes(grados float64) float64 {
	return grados * math.Pi / 180
}

func Haversine(coord1, coord2 Coordenadas) float64 {
	lat1 := GradosARadianes(coord1.Lat)
	lon1 := GradosARadianes(coord1.Lon)
	lat2 := GradosARadianes(coord2.Lat)
	lon2 := GradosARadianes(coord2.Lon)

	dLat := lat2 - lat1
	dLon := lon2 - lon1

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distancia := radioTierraKM * c
	return distancia
}
