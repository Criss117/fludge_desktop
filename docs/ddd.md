# 🧠 CONTEXTO COMPLETO DEL PROYECTO — POS DESKTOP (WAILS + GO + REACT)

Este documento define la arquitectura, decisiones técnicas y organización del proyecto para que una IA pueda continuar el desarrollo de forma consistente.

---

# 🧩 VISIÓN GENERAL

Aplicación POS (Point of Sale) desktop con:

- Funcionamiento 100% local/offline
- Arquitectura modular basada en DDD
- Soporte multitenant (multi-organización)
- UI en React, backend en Go
- Comunicación directa vía bindings de Wails

---

# 🧱 STACK TECNOLÓGICO

- Wails v2 (desktop app)
- Go (backend, arquitectura DDD)
- React + TypeScript (frontend)
- SQLite (base de datos local)
- sqlc (queries tipadas)
- Zod (validaciones frontend)

---

# 🧠 PRINCIPIOS ARQUITECTÓNICOS

## 1. DDD

- Separación por Bounded Contexts
- Agregados como núcleo del dominio
- Value Objects para invariantes
- Use Cases para escritura
- Queries directas para lectura

---

## 2. CQRS (ligero)

Escritura → Use Cases + Aggregates  
Lectura → sqlc directo → responses

---

## 3. Clean Architecture

Handler → UseCase → Domain → Repository → DB

---

## 4. Multitenancy

organization_id en TODAS las entidades

---

## 5. Offline-first

Sin APIs externas  
Sin auth server  
Todo local

---

# 🧩 BOUNDED CONTEXTS

## 🔐 IAM (Identity & Access)

Responsable de:

- Operadores
- Organizaciones
- Membresías
- Equipos y permisos
- Sesión

---

### Aggregate Roots

#### Operator

Tipos:

- ROOT → puede crear organizaciones
- EMPLOYEE → pertenece a una sola organización

Campos:

- id
- name
- username
- email (VO)
- pin (VO)
- operator_type
- timestamps

---

#### Organization

- id
- name
- slug (VO)
- legal_name
- address
- Members[]
- Teams[]

---

#### AppState (persistido)

- active_operator_id
- active_organization_id

---

### Entidades hijas

#### Member

- operator_id
- organization_id
- role (ROOT | MEMBER)

---

#### Team

- name
- permissions
- members[]

---

#### TeamMember

- operator_id
- team_id

---

## 📦 Catálogo

Aggregate roots independientes:

- Product
- Category
- Supplier

Relaciones por ID (no agregados anidados)

---

## 📦 Inventario (pendiente)

- StockMovement

---

## 💳 POS (pendiente)

- Ticket
- TicketDetail

---

## 💰 Crédito (pendiente)

- Customer
- CreditPayment

---

# 🧠 MODELADO CLAVE

## OperatorType vs MemberRole

OperatorType → nivel GLOBAL  
MemberRole → nivel ORGANIZACIÓN

---

### OperatorType

- ROOT
- EMPLOYEE

---

### MemberRole

- ROOT
- MEMBER

---

Regla importante:

ROOT operator NO necesita pertenecer a teams

---

# 🧠 SESIÓN

## AppState (persistido)

```go
type AppState struct {
	ActiveOrganizationID *string
	ActiveOperatorID     *string
}
```

---

## SessionState (en memoria)

```go
type SessionState struct {
	ActiveOrganization *Organization
	ActiveOperator     *ActiveOperator
}

type ActiveOperator struct {
	Operator *Operator
	Member   *Member
	Teams    []*Team
}
```

---

## Construcción

```go
BuildSessionState(operator, org, member, teams)
```

---

## Actualización

UseCase → Handler → StateChangeEvent → App → SessionState + AppState

---

Regla clave:

El dominio NO conoce la sesión

---

# 🔄 EVENTOS DE ESTADO

```go
type StateChangeEvent struct {
	Type         OnStateChangeType
	Operator     *Operator
	Organization *Organization
	Member       *Member
	Teams        []*Team
}
```

Tipos:

- SignUp
- SignIn
- SignOut
- SwitchOrganization

---

# 🧠 USE CASES IMPORTANTES

## Auth Flow

Register → SignIn → acceso

---

## RegisterRootOperator

- crea operator ROOT
- no maneja sesión
- handler emite evento

---

## SignIn

- valida credenciales
- retorna operator

---

## SwitchOrganization

Use case obligatorio (no query)

Valida:

- operador activo
- pertenencia

Retorna:

- Operator
- Organization
- Member
- Teams

---

# 🔥 TRANSACCIONES

Responsabilidad:

UseCase controla la transacción

---

## TransactionManager

```go
type TxManager interface {
	WithTx(ctx context.Context, fn func(q *db.Queries) error) error
}
```

---

## Uso

```go
txManager.WithTx(ctx, func(q *db.Queries) error {
	// operaciones
})
```

---

# 🧠 SQLC

## Config

```yaml
version: "2"
sql:
  - engine: "sqlite"
    queries: "internal/shared/db/queries.sql"
    schema: "internal/shared/db/schema.sql"
    gen:
      go:
        package: "db"
        out: "internal/shared/db"
        emit_json_tags: true
        overrides:
          - db_type: "blob"
            go_type: "encoding/json.RawMessage"
```

---

## Ubicación

internal/shared/db/

- schema.sql
- queries.sql
- código generado

---

# 🧠 HANDLERS

```go
type IamHandler struct {
	app           *application.Container
	getCtx        func() context.Context
	onStateChange func(StateChangeEvent)
}
```

---

Flujo:

React → Handler → UseCase → Event → App

---

Regla:

NO guardar ctx en struct

---

# 🧠 APP (COMPOSITION ROOT)

Responsable de:

- inicializar DB
- crear repositorios
- crear containers
- inyectar handlers
- manejar SessionState

---

## StateChange handling

```go
func (a *App) handleStateChange(e StateChangeEvent)
```

Actualiza:

- SessionState (memoria)
- AppState (DB)

---

# 🧠 SHARED

Ubicación:

internal/shared/

---

## Contenido

- db/ → sqlc + schema + queries
- dbutils/ → helpers time/null
- derrors/ → errores compartidos
- lib/ → uuid, helpers
- valueobjects/ → Email

---

Regla:

Shared NO contiene lógica de dominio

---

# 🧠 REGLAS IMPORTANTES

- Aggregates con campos públicos
- NewX valida, ReconstituteX no valida
- UseCases solo para escritura
- Queries no usan aggregates
- Repos no contienen lógica de negocio
- SessionState fuera del dominio
- organization_id obligatorio en todo
- Eventos para sincronizar estado

---

# 🎯 OBJETIVO

Mantener una arquitectura:

- consistente
- predecible
- extensible
- fácil de razonar

para soportar crecimiento del POS sin romper el dominio.
