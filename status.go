package main

type Status struct {
  Cpf               string            `json:"cpf"`
  Nome              string            `json:"nome"`
  Ultima_Consulta   UltimaConsulta   `json:"ultimaConsulta"`
  Ultima_Compra     UltimaCompra     `json:"ultimaCompra"`
}