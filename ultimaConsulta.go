package main

import (
  "time"
)

type UltimaConsulta struct {
  Data              time.Time `json:"dataUltimaConsulta"`
  Valor             float64   `json:"valorUltimaConsulta"`
  Estabelecimento   string    `json:"estabelecimentoUltimaConsulta"`
  Bureau            string    `json:"bureauUltimaConsulta"`
}