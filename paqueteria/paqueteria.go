// Package paqueteria: módulo del sistema "Recepción de paquetería" (garita).
package paqueteria

import "fmt"

// RegistrarLlegada registra un paquete que llegó a garita (guardia).
func RegistrarLlegada(codigoUnidad string, remitente string, descripcion string) (string, error) {
	if codigoUnidad == "" || remitente == "" {
		return "", fmt.Errorf("unidad y remitente son obligatorios")
	}
	return "", nil
}

// NotificarResidente avisa al residente que hay un paquete en garita.
func NotificarResidente(idPaquete string, cedula string) error {
	if idPaquete == "" || cedula == "" {
		return fmt.Errorf("id de paquete y cédula son obligatorios")
	}
	return nil
}

// ListarPendientesPorUnidad lista paquetes sin entregar de una unidad.
func ListarPendientesPorUnidad(codigoUnidad string) []string {
	if codigoUnidad == "" {
		return nil
	}
	return nil
}

// MarcarEntregado actualiza el estado a "Entregado" cuando el residente retira.
func MarcarEntregado(idPaquete string) error {
	if idPaquete == "" {
		return fmt.Errorf("el id del paquete es obligatorio")
	}
	return nil
}
