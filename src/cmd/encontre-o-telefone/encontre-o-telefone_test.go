package main

import "testing"

var testValues = map[string]string {
    "A": "2",
    "AC": "22",
    "ABC": "222",
    "ABCD": "2223",
    "ABCD-": "2223-",
    "The-quick-brown-fox-jumps-over-the-lazy-dog": "843-78425-27696-369-58677-6837-843-5299-364",
}

func TestTestesDosValores(t *testing.T) {
    for key, value := range(testValues) {
        numTelefone, _ := getTelefone(key)
        
        if numTelefone != value {
            t.Error("Número de telefone deveria ser ", value, ", mas foi ", numTelefone)
        }
    }
}

func TestComNumeroInvalido(t *testing.T) {
  _, err := getTelefone("ABCD ")

  if err == "" {
      t.Error("Deveria ter sido retornado um erro ")
  } 

}
