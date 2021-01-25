package main

import (
  "time"
)

type UltimaCompra struct {
  Data              time.Time `json:"dataUltimaCompraCartao"`
  Valor             float64   `json:"valorUltimaCompraCartao"`
  Estabelecimento   string  `json:"estabelecimentoUltimaCompraCartao"`
}