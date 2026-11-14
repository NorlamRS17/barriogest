// Package residentes: módulo del sistema "Residentes" (propietarios / inquilinos).
package residentes

import "fmt"

// CrearResidente da de alta un residente (rol residente o datos de contacto).
func CrearResidente(nombre string, cedula string, email string, telefono string) error {
	if nombre == "" || cedula == "" {
		return fmt.Errorf("nombre y cédula son obligatorios")
	}
	return nil
}

// AsignarAUnidad vincula un residente a una unidad (casa/depto).
func AsignarAUnidad(cedula string, codigoUnidad string) error {
	if cedula == "" || codigoUnidad == "" {
		return fmt.Errorf("cédula y código de unidad son obligatorios")
	}
	return nil
}

// ActualizarContacto permite al residente actualizar email y teléfono.
func ActualizarContacto(cedula string, email string, telefono string) error {
	if cedula == "" {
		return fmt.Errorf("la cédula es obligatoria")
	}
	return nil
}

// ObtenerPerfil consulta los datos del residente.
func ObtenerPerfil(cedula string) (string, error) {
	if cedula == "" {
		return "", fmt.Errorf("la cédula es obligatoria")
	}
	return "", nil
}

// CambiarContrasena actualiza la contraseña del residente.
func CambiarContrasena(cedula string, contrasenaActual string, contrasenaNueva string) error {
	if cedula == "" || contrasenaNueva == "" {
		return fmt.Errorf("cédula y contraseña nueva son obligatorias")
	}
	return nil
}

// ListarResidentes devuelve un resumen de residentes registrados.
func ListarResidentes() []string {
	return nil
}
