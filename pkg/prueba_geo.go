package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "net/url"
)

// Coordenadas representa latitud y longitud
type Coordenadas struct {
    Lat float64 `json:"lat,string"`
    Lon float64 `json:"lon,string"`
}

type Direccion struct {
    calle_numero  string
    poblacion     string
    codigo_postal int
}

// FormatDireccion formatea y devuelve la dirección como una cadena
func (d *Direccion) FormatDireccion() string {
    return fmt.Sprintf("%s, %s, %d", d.calle_numero, d.poblacion, d.codigo_postal)
}

// ObtenerCoordenadas realiza una consulta a la API de Nominatim para obtener latitud y longitud
func ObtenerCoordenadas(direccion Direccion) (Coordenadas, error) {
    // Formatear la dirección para la consulta usando el método FormatDireccion
    query := direccion.FormatDireccion()

    // Construir la URL de la API
    apiURL := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json&polygon=1&addressdetails=1", url.QueryEscape(query))

    // Realizar la llamada HTTP
    resp, err := http.Get(apiURL)
    if err != nil {
        return Coordenadas{}, err
    }
    defer resp.Body.Close()

    // Verificar el código de estado HTTP
    if resp.StatusCode != http.StatusOK {
        return Coordenadas{}, fmt.Errorf("API request failed with status code %d", resp.StatusCode)
    }

    // Leer y decodificar la respuesta
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return Coordenadas{}, err
    }

    var respuestaAPI []Coordenadas
    err = json.Unmarshal(body, &respuestaAPI)
    if err != nil {
        return Coordenadas{}, err
    }

    // Asumir que la primera coincidencia es la correcta
    if len(respuestaAPI) == 0 {
        return Coordenadas{}, fmt.Errorf("no se encontraron coincidencias")
    }

    return respuestaAPI[0], nil
}

func main() {
    // Crear una instancia de Direccion
    direccion := Direccion{
        calle_numero:  "Calle Aurora 3",
        poblacion:     "Puerto Lumbreras",
        codigo_postal: 30890,
    }

    // Obtener coordenadas
    coords, err := ObtenerCoordenadas(direccion)
    if err != nil {
        fmt.Println("Error obteniendo coordenadas:", err)
        return
    }

    fmt.Println("Coordenadas:", coords)
}

