// Package unidades: módulo del sistema "Unidades" (casas / departamentos).
// En Go, cada módulo del sistema se programa como un paquete.
package unidades

import "fmt"

// CrearUnidad registra una unidad nueva (ej. "Casa 10", "Depto 4B").
func CrearUnidad(codigo string, tipo string, bloque string) error {
	if codigo == "" {
		return fmt.Errorf("el código de la unidad es obligatorio")
	}
	return nil
}

// ModificarUnidad actualiza datos de una unidad existente.
func ModificarUnidad(codigo string, tipo string, bloque string) error {
	if codigo == "" {
		return fmt.Errorf("el código de la unidad es obligatorio")
	}
	return nil
}

// EliminarUnidad da de baja una unidad del conjunto.
func EliminarUnidad(codigo string) error {
	if codigo == "" {
		return fmt.Errorf("el código de la unidad es obligatorio")
	}
	return nil
}

// ObtenerUnidad consulta una unidad por su código.
func ObtenerUnidad(codigo string) (string, error) {
	if codigo == "" {
		return "", fmt.Errorf("el código de la unidad es obligatorio")
	}
	return "", nil
}

// ListarUnidades devuelve un resumen de todas las unidades.
func ListarUnidades() []string {
	return nil
}
