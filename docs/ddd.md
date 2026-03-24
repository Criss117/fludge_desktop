# 🧠 POS Desktop App — Contexto Completo (DDD + Go + Wails + sqlc)

## 📌 Concepto General

Aplicación **POS (Point of Sale) desktop**, diseñada para operar **offline-first**, con soporte **multitenant** mediante `organization_id` en todos los agregados.

### Stack

- **Wails v2** → aplicación desktop
- **Go** → backend (DDD + Clean Architecture)
- **React + TypeScript** → frontend
- **SQLite + sqlc** → persistencia (queries tipadas, sin ORM)

---

## 🧠 Principios Arquitectónicos

- **DDD real (no superficial)**
- **CQRS ligero**
- **Bounded Contexts bien definidos**
- **Separación estricta de capas**
- **Sin acoplamiento entre dominios**

---

## 🧩 Modelo Mental del Sistema

```
IAM        → quién eres
Catalog    → qué vendes
Inventory  → cuánto tienes
Sales      → qué vendiste
Credit     → quién debe
```

---

# 🧩 BOUNDED CONTEXTS

---

## 🔐 IAM (Identity & Access)

### Aggregate Roots:

- `Operator`
- `Organization`
- `AppState`

### Entidades hijas:

- `Member`
- `Team`
- `TeamMember`

### Responsabilidades:

- Autenticación (PIN)
- Gestión de organizaciones
- Permisos
- Estado activo (session)

---

## 📦 Catalog

### Aggregate Roots:

- `Product`
- `Category`
- `Supplier`

### Reglas:

- ❗ `Product` **NO tiene stock**
- Solo define datos estáticos
- Referencias por ID

---

## 📦 Inventory

### Aggregate Roots:

- `InventoryItem`

### Entidades hijas:

- `StockMovement`

### Responsabilidades:

- Estado actual (`quantity`, `min_stock`)
- Historial de movimientos
- Ajustes de stock

---

## 🛒 Sales (POS)

### Aggregate Roots:

- `Sale` (persistido como `ticket`)

### Entidades hijas:

- `SaleItem` (persistido como `ticket_detail`)

### Responsabilidades:

- Crear ventas
- Calcular totales
- Manejar pagos
- Emitir eventos

---

## 💳 Credit

### Aggregate Roots:

- `Customer`

### Entidades hijas:

- `CreditPayment`

### Responsabilidades:

- Límite de crédito
- Balance actual (**cacheado**)
- Registro de pagos

---

# 🔗 RELACIÓN ENTRE CONTEXTOS

```
IAM → todos

Catalog → Inventory, Sales

Sales → Inventory (evento)
Sales → Credit (evento)

Inventory → reacciona a eventos

Credit → depende de Sales
```

---

# 🧱 CAPAS DE ARQUITECTURA

---

## Domain

- Lógica de negocio pura
- Sin dependencias externas

Contiene:

```
aggregates/
valueobjects/
ports/
services/
```

---

## Application

- Orquesta casos de uso

Contiene:

```
usecases/
commands/
responses/
queries/ (interfaces, no sqlc)
```

---

## Infrastructure

- Implementaciones técnicas

Contiene:

```
repositories/
mappers/
sqlc (desde shared)
```

---

## Interfaces

- Adaptadores (Wails)

Contiene:

```
handlers/
session access
DTO mapping
```

---

# 📁 ESTRUCTURA DEL PROYECTO

```
internal/
  shared/
    db/
      schema/
        001_iam.sql
        002_catalog.sql
        003_inventory.sql
        004_sales.sql
        005_credit.sql

      iam_sqlc/
      catalog_sqlc/
      inventory_sqlc/
      sales_sqlc/
      credit_sqlc/

      db.go

    derrors/
    valueobjects/
    events/
    types/
    lib/

  appstate/

  platform/
    iam/
    catalog/
    inventory/
    sales/
    credit/
```

---

# 🗄️ BASE DE DATOS

## Principios

- Multitenant (`organization_id`)
- Soft delete (`deleted_at`)
- Constraints fuertes
- Índices optimizados
- DB garantiza unicidad

---

## Separación por Contexto

### IAM

```
operator
organization
app_state
member
team
team_member
```

### Catalog

```
product (SIN stock)
category
supplier
```

### Inventory

```
inventory_item  ← estado actual
stock_movement  ← historial
```

### Sales

```
ticket
ticket_detail
```

### Credit

```
customer
credit_payment
```

---

# ⚠️ REGLAS CRÍTICAS

1. **Product NO tiene stock**
2. **Inventory maneja TODO el stock**
3. **Contextos NO comparten agregados**
4. Comunicación solo vía:
   - IDs
   - eventos

5. **sqlc solo se usa en infrastructure**
6. **1 bounded context = 1 paquete sqlc**
7. **Use cases no instancian dependencias**
8. **Handlers no contienen lógica de negocio**
9. **Shared NO contiene lógica de dominio**

---

# 🧠 SHARED KERNEL

## Contiene únicamente:

```
db/           → sqlc + conexión
derrors/      → errores genéricos
valueobjects/ → Email, Phone
events/       → interfaces de eventos
types/        → pagination, filters
lib/          → uuid, time
```

## NO contiene:

- lógica de negocio
- agregados
- repositorios
- DTOs
- services de dominio

---

# ⚙️ SQLC CONFIG

- 1 config por bounded context (en el mismo yaml)
- 1 `queries.sql` por contexto
- 1 package generado por contexto

Ejemplo:

```
catalog_sqlc/
inventory_sqlc/
sales_sqlc/
```

---

# 📡 EVENTOS DE DOMINIO

Ejemplo:

```
SaleCompleted
  → Inventory descuenta stock
  → Credit genera deuda
```

---

# 🔒 SESSION

- `AppState` → IDs activos (dominio)
- `SessionState` → agregados cargados (app layer)

Handlers usan `SessionState`, no DB directa

---

# 🚀 OBJETIVO DEL SISTEMA

- Offline-first
- Alta coherencia de dominio
- POS rápido y consistente
- Escalable por módulos
- Mantenible a largo plazo

---

# 🧠 INSTRUCCIONES PARA LA IA

Cualquier cambio debe:

- respetar bounded contexts
- no acoplar dominios
- mantener separación de capas
- evitar lógica en handlers
- evitar uso de sqlc fuera de infrastructure
- mantener invariantes en domain
- tratar DB como detalle de infraestructura

Este es un sistema diseñado con DDD real y arquitectura limpia.  
Las decisiones deben priorizar consistencia del dominio sobre conveniencia técnica.
