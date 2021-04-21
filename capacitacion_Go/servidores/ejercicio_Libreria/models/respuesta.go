package models

type Respuesta struct {
	Mensaje string `json:"mensaje"`
	Codigo  int    `json:"codigo"`
}
