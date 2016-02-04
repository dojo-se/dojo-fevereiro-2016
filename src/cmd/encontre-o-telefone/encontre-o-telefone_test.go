package main

import "testing"

func TestSimples_A(t *testing.T) {
  numTelefone := getTelefone("A")

  if numTelefone != "2" {
    t.Error("Número de telefone deveria ser '2', mas foi ", numTelefone)
  }
}

func TestSimples_AC(t *testing.T) {
  numTelefone := getTelefone("AC")

  if numTelefone != "22" {
    t.Error("Número de telefone deveria ser '22', mas foi ", numTelefone)
  }
}
