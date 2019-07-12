package main

import (
  "testing"
  
  
)

func TestGreeting(t *testing.T){
  resultado := Greeting("Code.education Rocks!")
  if resultado != "<b>Code.education Rocks!</b>"{
    t.Errorf("Resultado esperado esperada: %s","<b>Code.education Rocks!</b>")
  }
}

func TestSqrtLooping(t *testing.T){
  
  valor := SqrtLooping(1000000000)
  
  if valor <= 21081851051977{
    t.Errorf("Resultado precisa ser maior que: %f",valor)
  }
}