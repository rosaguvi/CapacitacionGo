package models

type DatConsulta struct {
	Parametro     string `json:"parametro"`
	Rango_inicial int    `json:"rango_inicial"`
	Rango_final   int    `json:"rango_final"`
}
