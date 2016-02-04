package main

import "testing"

func TestSimples_A(t *testing.T) {
  numTelefone, _ := getTelefone("A")

  if numTelefone != "2" {
    t.Error("Número de telefone deveria ser '2', mas foi ", numTelefone)
  }
}

func TestSimples_AC(t *testing.T) {
  numTelefone, _ := getTelefone("AC")

  if numTelefone != "22" {
    t.Error("Número de telefone deveria ser '22', mas foi ", numTelefone)
  }
}

func TestSimples_ABC(t *testing.T) {
  numTelefone, _ := getTelefone("ABC")

  if numTelefone != "222" {
    t.Error("Número de telefone deveria ser '222', mas foi ", numTelefone)
  }
}

func TestSimples_ABCD(t *testing.T) {
  numTelefone, _ := getTelefone("ABCD")

  if numTelefone != "2223" {
    t.Error("Número de telefone deveria ser '2223', mas foi ", numTelefone)
  }
}

func TestSimples_ABCD_Hifen(t *testing.T) {
  numTelefone, _ := getTelefone("ABCD-")

  if numTelefone != "2223-" {
    t.Error("Número de telefone deveria ser '2223-', mas foi ", numTelefone)
  }
}

func TestSimples_ABCD_Espaço(t *testing.T) {
  _, err := getTelefone("ABCD ")

  if err == "" {
      t.Error("Deveria ter sido retornado um erro ")
  } 

}
