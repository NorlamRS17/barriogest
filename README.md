# BarrioGest

Sistema de gestión para **conjuntos residenciales, condominios, urbanizaciones cerradas y barrios**.  
Centraliza administración, finanzas (alícuotas), reservas de áreas comunes y control de visitas en garita.

---

## Concepto clave: módulo ≠ paquete ≠ `go.mod`

| Concepto | Qué es | En BarrioGest |
|---|---|---|
| **Módulo del sistema** | Área funcional del negocio (finanzas, reservas…) | 9 módulos |
| **Paquete Go** | Carpeta de código que implementa ese módulo | una carpeta por módulo, ej. `finanzas/` |
| **Módulo de Go** | El proyecto completo: la carpeta con **un solo** `go.mod` | `barriogest` (este repositorio) |

Reglas:

- **Un solo `go.mod`** en todo el repo.
- **Un paquete por módulo del sistema** (mismo nombre que la carpeta, minúsculas, sin guiones bajos).
- `main.go` solo arranca; no concentra la lógica.
- Funciones exportadas con **Mayúscula** inicial.
- Evitar paquetes genéricos tipo `utils`.
- Los paquetes no deben importarse entre sí en ciclo.

---

## Objetivo general

Desarrollar un sistema de gestión centralizada para optimizar procesos administrativos, financieros y de seguridad en conjuntos residenciales, mejorando la comunicación entre la administración y los residentes.

### Objetivos específicos

1. Finanzas: registro y control de alícuotas y conciliación de pagos.
2. Reservas: áreas comunes con validación de disponibilidad.
3. Accesos: pre-autorización de visitas y confirmación en garita.

---

## Roles

| Rol | Quién | Alcance |
|---|---|---|
| Administrador | Directiva o administrador contratado | Control total |
| Residente | Propietario o inquilino | Perfil, pagos, reservas, visitas |
| Guardia | Seguridad en garita | Visitas del día, ingreso/salida |

---

## Alcance

### Sí hace

- Unidades y asignación de residentes
- Alícuotas mensuales y pagos con comprobante (manual)
- Reservas de áreas comunes con validación de horario
- Visitas pre-autorizadas + confirmación del guardia

### No incluye esta versión

- SRI / facturación electrónica
- Pasarela de pagos (tarjeta, Stripe, PayPal)
- Hardware (cámaras, plumas)
- App móvil nativa (luego: web responsive)

---

## Módulos del sistema (9)

1. Unidades  
2. Residentes  
3. Finanzas (alícuotas y pagos)  
4. Reservas de áreas comunes  
5. Accesos / visitas (garita)  
6. Paquetería  
7. Incidencias / mantenimiento  
8. Comunicados / cartelera  
9. Votaciones / encuestas  

SRI, pasarela y hardware están **fuera** de alcance.

---

## Árbol del repositorio

```text
barriogest/                      ← módulo de Go (un solo go.mod)
├── go.mod
├── README.md
├── main.go                      ← package main (solo arranque)
├── unidades/unidades.go
├── residentes/residentes.go
├── finanzas/finanzas.go
├── reservas/reservas.go
├── accesos/accesos.go
├── paqueteria/paqueteria.go
├── incidencias/incidencias.go
├── comunicados/comunicados.go
└── votaciones/votaciones.go
```

---

## Tablas de funciones por módulo (para el PDF)

### 1. Unidades (`unidades`)

| Función | Parámetros | Retorna | Qué hace |
|---|---|---|---|
| `CrearUnidad` | `codigo, tipo, bloque string` | `error` | Registra casa/depto |
| `ModificarUnidad` | `codigo, tipo, bloque string` | `error` | Actualiza datos |
| `EliminarUnidad` | `codigo string` | `error` | Da de baja la unidad |
| `ObtenerUnidad` | `codigo string` | `string, error` | Consulta una unidad |
| `ListarUnidades` | — | `[]string` | Lista todas |

### 2. Residentes (`residentes`)

| Función | Parámetros | Retorna | Qué hace |
|---|---|---|---|
| `CrearResidente` | `nombre, cedula, email, telefono string` | `error` | Alta de residente |
| `AsignarAUnidad` | `cedula, codigoUnidad string` | `error` | Vincula a una unidad |
| `ActualizarContacto` | `cedula, email, telefono string` | `error` | Actualiza contacto |
| `ObtenerPerfil` | `cedula string` | `string, error` | Consulta perfil |
| `CambiarContrasena` | `cedula, contrasenaActual, contrasenaNueva string` | `error` | Cambia contraseña |
| `ListarResidentes` | — | `[]string` | Lista residentes |

### 3. Finanzas (`finanzas`)

| Función | Parámetros | Retorna | Qué hace |
|---|---|---|---|
| `GenerarAlicuotasMensuales` | `anio, mes int, monto float64` | `error` | Genera deuda del mes |
| `ReportarPago` | `cedula, codigoUnidad string, monto float64, referencia, rutaComprobante string` | `error` | Residente reporta pago |
| `AprobarPago` | `idPago string` | `error` | Admin aprueba |
| `RechazarPago` | `idPago, motivo string` | `error` | Admin rechaza |
| `ConsultarEstadoCuenta` | `codigoUnidad string` | `string, error` | Estado de cuenta |
| `ListarAlDia` | — | `[]string` | Unidades al día |
| `ListarEnMora` | — | `[]string` | Unidades en mora |

### 4. Reservas (`reservas`)

| Función | Parámetros | Retorna | Qué hace |
|---|---|---|---|
| `CrearArea` | `nombre, horaInicio, horaFin string` | `error` | Alta de área común |
| `EditarArea` | `nombre, horaInicio, horaFin string` | `error` | Edita área |
| `ListarAreas` | — | `[]string` | Lista áreas |
| `Reservar` | `cedula, nombreArea, fecha, horaInicio, horaFin string` | `error` | Agenda (valida choque/mora) |
| `CancelarReserva` | `idReserva, motivo string` | `error` | Cancela reserva |

### 5. Accesos / visitas (`accesos`)

| Función | Parámetros | Retorna | Qué hace |
|---|---|---|---|
| `PreautorizarVisita` | `cedulaResidente, nombreVisita, documentoOPlaca, fecha string` | `error` | Residente pre-autoriza |
| `ListarVisitasDelDia` | `fecha string` | `[]string` | Lista del día (guardia) |
| `RegistrarIngreso` | `idVisita string` | `error` | Estado → En conjunto |
| `RegistrarSalida` | `idVisita string` | `error` | Estado → Salió |

### 6. Paquetería (`paqueteria`)

| Función | Parámetros | Retorna | Qué hace |
|---|---|---|---|
| `RegistrarLlegada` | `codigoUnidad, remitente, descripcion string` | `string, error` | Guardia registra paquete |
| `NotificarResidente` | `idPaquete, cedula string` | `error` | Avisa al residente |
| `ListarPendientesPorUnidad` | `codigoUnidad string` | `[]string` | Paquetes sin entregar |
| `MarcarEntregado` | `idPaquete string` | `error` | Marca como entregado |

### 7. Incidencias (`incidencias`)

| Función | Parámetros | Retorna | Qué hace |
|---|---|---|---|
| `CrearTicket` | `cedula, titulo, descripcion, rutaFoto string` | `string, error` | Residente reporta |
| `AsignarResponsable` | `idTicket, responsable string` | `error` | Admin asigna |
| `ActualizarEstado` | `idTicket, estado string` | `error` | Abierto / En progreso / Resuelto |
| `ListarTickets` | `estado string` | `[]string` | Lista incidencias |
| `ObtenerTicket` | `idTicket string` | `string, error` | Detalle de ticket |

### 8. Comunicados (`comunicados`)

| Función | Parámetros | Retorna | Qué hace |
|---|---|---|---|
| `Publicar` | `titulo, cuerpo, categoria string` | `string, error` | Admin publica aviso |
| `Editar` | `idComunicado, titulo, cuerpo string` | `error` | Edita comunicado |
| `Archivar` | `idComunicado string` | `error` | Archiva aviso |
| `ListarActivos` | — | `[]string` | Cartelera (lectura) |
| `ObtenerComunicado` | `idComunicado string` | `string, error` | Consulta un aviso |

### 9. Votaciones (`votaciones`)

| Función | Parámetros | Retorna | Qué hace |
|---|---|---|---|
| `CrearEncuesta` | `titulo string, opciones []string, fechaCierre string` | `string, error` | Admin crea encuesta |
| `Votar` | `idEncuesta, codigoUnidad, opcion string` | `error` | 1 voto por unidad |
| `CerrarEncuesta` | `idEncuesta string` | `error` | Cierra votación |
| `ConsultarResultados` | `idEncuesta string` | `string, error` | Conteo de votos |
| `ListarEncuestas` | `soloActivas bool` | `[]string` | Lista encuestas |

**Total:** 9 módulos / paquetes y más de 8 funciones.

---

## Cómo compilar y ejecutar

```bash
cd barriogest
go build ./...
go run .
```

Las funciones están declaradas con validación básica de parámetros. El proyecto compila con `go build ./...`.
