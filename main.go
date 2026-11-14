// Package main solo arranca BarrioGest.
// La lógica vive en los paquetes de cada módulo del sistema.
package main

import (
	"fmt"

	"barriogest/accesos"
	"barriogest/comunicados"
	"barriogest/finanzas"
	"barriogest/incidencias"
	"barriogest/paqueteria"
	"barriogest/reservas"
	"barriogest/residentes"
	"barriogest/unidades"
	"barriogest/votaciones"
)

func main() {
	fmt.Println("BarrioGest — gestión de condominios, conjuntos y barrios")

	_ = unidades.ListarUnidades()
	_ = residentes.ListarResidentes()
	_ = finanzas.ListarEnMora()
	_ = reservas.ListarAreas()
	_ = accesos.ListarVisitasDelDia("2026-11-15")
	_ = paqueteria.ListarPendientesPorUnidad("Casa-10")
	_ = incidencias.ListarTickets("")
	_ = comunicados.ListarActivos()
	_ = votaciones.ListarEncuestas(true)
}
