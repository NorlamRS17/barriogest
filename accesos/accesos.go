// Package accesos: módulo del sistema "Control de accesos / visitas" (garita).
package accesos

import "fmt"

// PreautorizarVisita registra una visita esperada (la crea el residente).
func PreautorizarVisita(cedulaResidente string, nombreVisita string, documentoOPlaca string, fecha string) error {
	if cedulaResidente == "" || nombreVisita == "" || fecha == "" {
		return fmt.Errorf("residente, nombre de visita y fecha son obligatorios")
	}
	return nil
}

// ListarVisitasDelDia muestra las visitas esperadas para una fecha (uso del guardia).
func ListarVisitasDelDia(fecha string) []string {
	if fecha == "" {
		return nil
	}
	return nil
}

// RegistrarIngreso marca la visita como "En conjunto".
func RegistrarIngreso(idVisita string) error {
	if idVisita == "" {
		return fmt.Errorf("el id de la visita es obligatorio")
	}
	return nil
}

// RegistrarSalida marca la visita como "Salió".
func RegistrarSalida(idVisita string) error {
	if idVisita == "" {
		return fmt.Errorf("el id de la visita es obligatorio")
	}
	return nil
}
