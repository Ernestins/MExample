package main

import (
	"fmt"
	"time"
)

func main() {
	logging := true     //*
	name := "anton"     //*
	age := 18           //*
	pi := 3.141592      //*
	start := time.Now() //*

	fmt.Printf("boolean: %t, string: %s, integer: %d, float64: %e, object: %v", logging, name, age, pi, start)
}