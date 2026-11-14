// Package comunicados: módulo del sistema "Comunicados / cartelera virtual".
package comunicados

import "fmt"

// Publicar crea un aviso visible para los residentes (admin).
func Publicar(titulo string, cuerpo string, categoria string) (string, error) {
	if titulo == "" || cuerpo == "" {
		return "", fmt.Errorf("título y cuerpo son obligatorios")
	}
	return "", nil
}

// Editar actualiza un comunicado existente.
func Editar(idComunicado string, titulo string, cuerpo string) error {
	if idComunicado == "" {
		return fmt.Errorf("el id del comunicado es obligatorio")
	}
	return nil
}

// Archivar oculta o archiva un comunicado.
func Archivar(idComunicado string) error {
	if idComunicado == "" {
		return fmt.Errorf("el id del comunicado es obligatorio")
	}
	return nil
}

// ListarActivos muestra la cartelera para el residente (solo lectura).
func ListarActivos() []string {
	return nil
}

// ObtenerComunicado consulta un aviso por id.
func ObtenerComunicado(idComunicado string) (string, error) {
	if idComunicado == "" {
		return "", fmt.Errorf("el id del comunicado es obligatorio")
	}
	return "", nil
}
