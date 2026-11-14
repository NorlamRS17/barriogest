// Package incidencias: módulo del sistema "Gestión de incidencias / mantenimiento".
package incidencias

import "fmt"

// CrearTicket abre una incidencia reportada por el residente.
func CrearTicket(cedula string, titulo string, descripcion string, rutaFoto string) (string, error) {
	if cedula == "" || titulo == "" {
		return "", fmt.Errorf("cédula y título son obligatorios")
	}
	return "", nil
}

// AsignarResponsable asigna el ticket a personal de mantenimiento (admin).
func AsignarResponsable(idTicket string, responsable string) error {
	if idTicket == "" || responsable == "" {
		return fmt.Errorf("id de ticket y responsable son obligatorios")
	}
	return nil
}

// ActualizarEstado cambia el estado (ej. Abierto → En progreso → Resuelto).
func ActualizarEstado(idTicket string, estado string) error {
	if idTicket == "" || estado == "" {
		return fmt.Errorf("id de ticket y estado son obligatorios")
	}
	return nil
}

// ListarTickets lista incidencias, opcionalmente filtradas por estado.
func ListarTickets(estado string) []string {
	return nil
}

// ObtenerTicket consulta el detalle de una incidencia.
func ObtenerTicket(idTicket string) (string, error) {
	if idTicket == "" {
		return "", fmt.Errorf("el id del ticket es obligatorio")
	}
	return "", nil
}
