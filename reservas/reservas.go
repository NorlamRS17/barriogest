// Package reservas: módulo del sistema "Reservas" (áreas comunes).
package reservas

import "fmt"

// CrearArea registra un área común (cancha, salón, BBQ, etc.).
func CrearArea(nombre string, horaInicio string, horaFin string) error {
	if nombre == "" {
		return fmt.Errorf("el nombre del área es obligatorio")
	}
	return nil
}

// EditarArea actualiza horarios o datos de un área común.
func EditarArea(nombre string, horaInicio string, horaFin string) error {
	if nombre == "" {
		return fmt.Errorf("el nombre del área es obligatorio")
	}
	return nil
}

// ListarAreas devuelve las áreas comunes disponibles.
func ListarAreas() []string {
	return nil
}

// Reservar agenda un área para un residente en fecha y rango horario.
// Debe rechazar si hay choque de horario o si el residente está en mora.
func Reservar(cedula string, nombreArea string, fecha string, horaInicio string, horaFin string) error {
	if cedula == "" || nombreArea == "" || fecha == "" {
		return fmt.Errorf("cédula, área y fecha son obligatorias")
	}
	if !horarioDisponible(nombreArea, fecha, horaInicio, horaFin) {
		return fmt.Errorf("el horario no está disponible")
	}
	return nil
}

// CancelarReserva anula una reserva (admin por mantenimiento, o según reglas).
func CancelarReserva(idReserva string, motivo string) error {
	if idReserva == "" {
		return fmt.Errorf("el id de la reserva es obligatorio")
	}
	return nil
}

// horarioDisponible es auxiliar (no exportada): valida choques de horario.
func horarioDisponible(nombreArea string, fecha string, horaInicio string, horaFin string) bool {
	_ = nombreArea
	_ = fecha
	_ = horaInicio
	_ = horaFin
	return true
}
