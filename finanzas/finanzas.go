// Package finanzas: módulo del sistema "Finanzas" (alícuotas / cuotas y pagos).
package finanzas

import "fmt"

// GenerarAlicuotasMensuales crea la deuda del mes para todas las unidades.
func GenerarAlicuotasMensuales(anio int, mes int, monto float64) error {
	if anio < 2000 || mes < 1 || mes > 12 {
		return fmt.Errorf("año o mes inválido")
	}
	if monto <= 0 {
		return fmt.Errorf("el monto debe ser mayor a cero")
	}
	return nil
}

// ReportarPago registra un pago informado por el residente (referencia + comprobante).
func ReportarPago(cedula string, codigoUnidad string, monto float64, referencia string, rutaComprobante string) error {
	if cedula == "" || codigoUnidad == "" || referencia == "" {
		return fmt.Errorf("cédula, unidad y referencia son obligatorias")
	}
	if monto <= 0 {
		return fmt.Errorf("el monto debe ser mayor a cero")
	}
	return nil
}

// AprobarPago marca un pago reportado como válido (admin).
func AprobarPago(idPago string) error {
	if idPago == "" {
		return fmt.Errorf("el id del pago es obligatorio")
	}
	return nil
}

// RechazarPago marca un pago reportado como inválido (admin).
func RechazarPago(idPago string, motivo string) error {
	if idPago == "" {
		return fmt.Errorf("el id del pago es obligatorio")
	}
	return nil
}

// ConsultarEstadoCuenta devuelve el resumen de saldos de una unidad/residente.
func ConsultarEstadoCuenta(codigoUnidad string) (string, error) {
	if codigoUnidad == "" {
		return "", fmt.Errorf("el código de unidad es obligatorio")
	}
	return "", nil
}

// ListarAlDia lista unidades/residentes al día en pagos.
func ListarAlDia() []string {
	return nil
}

// ListarEnMora lista unidades/residentes con pagos pendientes.
func ListarEnMora() []string {
	return nil
}
