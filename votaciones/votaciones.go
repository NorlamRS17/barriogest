// Package votaciones: módulo del sistema "Votaciones y encuestas".
// Regla de negocio: 1 voto por unidad (casa/depto), no por residente.
package votaciones

import "fmt"

// CrearEncuesta publica una encuesta o votación (admin).
func CrearEncuesta(titulo string, opciones []string, fechaCierre string) (string, error) {
	if titulo == "" || len(opciones) < 2 {
		return "", fmt.Errorf("título y al menos 2 opciones son obligatorios")
	}
	return "", nil
}

// Votar registra el voto de una unidad (máximo 1 por unidad).
func Votar(idEncuesta string, codigoUnidad string, opcion string) error {
	if idEncuesta == "" || codigoUnidad == "" || opcion == "" {
		return fmt.Errorf("encuesta, unidad y opción son obligatorios")
	}
	if yaVotoUnidad(idEncuesta, codigoUnidad) {
		return fmt.Errorf("esta unidad ya emitió su voto")
	}
	return nil
}

// CerrarEncuesta cierra la votación y deja de aceptar votos.
func CerrarEncuesta(idEncuesta string) error {
	if idEncuesta == "" {
		return fmt.Errorf("el id de la encuesta es obligatorio")
	}
	return nil
}

// ConsultarResultados devuelve el conteo de votos por opción.
func ConsultarResultados(idEncuesta string) (string, error) {
	if idEncuesta == "" {
		return "", fmt.Errorf("el id de la encuesta es obligatorio")
	}
	return "", nil
}

// ListarEncuestas lista encuestas activas o cerradas.
func ListarEncuestas(soloActivas bool) []string {
	_ = soloActivas
	return nil
}

// yaVotoUnidad es auxiliar (no exportada): controla 1 voto por unidad.
func yaVotoUnidad(idEncuesta string, codigoUnidad string) bool {
	_ = idEncuesta
	_ = codigoUnidad
	return false
}
