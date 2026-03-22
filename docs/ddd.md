# 🧠 POS Desktop App — Contexto Completo del Proyecto (DDD + Go + Wails)

## 📌 Concepto General

Esta aplicación es un **POS (Point of Sale) desktop**, diseñada para operar **offline-first**, con soporte **multitenant** mediante `organization_id` en todos los agregados.

Stack tecnológico:

- **Wails v2** (desktop app)
- **Go** (backend con DDD + Clean Architecture)
- **React + TypeScript** (frontend)
- **SQLite + sqlc** (persistencia con queries tipadas)

La arquitectura sigue:

- **DDD (Domain Driven Design)**
- **CQRS ligero**
- **Separación por Bounded Contexts**
- **Arquitectura por capas (Domain / Application / Infrastructure / Interfaces)**

---

## 🧠 Modelo Mental del Sistema

```
IAM        → quién eres
Catalog    → qué vendes
Inventory  → cuánto tienes
Sales      → qué vendiste
Credit     → quién debe
```

---

## 🧩 Bounded Contexts

---

### 🔐 IAM (Identity & Access)

Responsable de identidad, permisos y sesión activa.

#### Aggregate Roots:

- `Operator`
- `Organization`
- `AppState`

#### Entidades hijas:

- `Member`
- `Team`
- `TeamMember`

#### Responsabilidades:

- Autenticación (PIN)
- Gestión de organizaciones
- Gestión de miembros y equipos
- Permisos
- Estado activo de la app

---

### 📦 Catalog

Define los productos. **NO maneja estado dinámico**.

#### Aggregate Roots:

- `Product`
- `Category`
- `Supplier`

#### Reglas:

- `Product` NO contiene stock
- Referencias por ID (no agregados embebidos)
- Precios validados por invariantes

---

### 📦 Inventory

Gestiona el stock real de productos.

#### Aggregate Roots:

- `InventoryItem`

#### Entidades hijas:

- `StockMovement`

#### Responsabilidades:

- Cantidad actual (`InventoryItem`)
- Historial (`StockMovement`)
- Ajustes de stock
- Reorden (min_stock)

---

### 🛒 Sales (POS)

Core del sistema.

#### Aggregate Roots:

- `Sale` (persistido como `ticket`)

#### Entidades hijas:

- `SaleItem` (persistido como `ticket_detail`)

#### Responsabilidades:

- Crear ventas
- Calcular totales
- Manejar pagos
- Emitir eventos de negocio

---

### 💳 Credit

Gestión de crédito y cobranza.

#### Aggregate Roots:

- `Customer`
- `CreditAccount` (implícito en `Customer`)

#### Entidades hijas:

- `CreditPayment`

#### Responsabilidades:

- Límites de crédito
- Balance actual (cacheado)
- Registro de pagos

---

## 🔗 Relación entre Contextos

```
IAM → todos

Catalog → Inventory, Sales

Sales → Inventory (evento)
Sales → Credit (evento)

Inventory → independiente (reacciona a eventos)

Credit → depende de Sales
```

---

## 📁 Estructura de Carpetas

```
internal/
  shared/
    db/
      iam_sqlc/
      catalog_sqlc/
      inventory_sqlc/
      sales_sqlc/
    derrors/
    valueobjects/
    events/

  appstate/

  platform/
    iam/
      domain/
        aggregates/
        valueobjects/
        ports/
        services/
      application/
        usecases/
        commands/
        responses/
        queries/
      infrastructure/
        persistence/
          repositories/
          mappers/
        services/
      interfaces/

    catalog/
      domain/
        aggregates/
        valueobjects/
        ports/
      application/
        usecases/
        commands/
        responses/
        queries/
      infrastructure/
        persistence/
          repositories/
          mappers/
      interfaces/

    inventory/
      domain/
        aggregates/
        valueobjects/
        ports/
      application/
        usecases/
        commands/
        responses/
        queries/
      infrastructure/
        persistence/
          repositories/
          mappers/
      interfaces/

    sales/
      domain/
        aggregates/
        valueobjects/
        ports/
      application/
        usecases/
        commands/
        responses/
        queries/
      infrastructure/
        persistence/
          repositories/
          mappers/
      interfaces/

    credit/
      domain/
        aggregates/
        valueobjects/
        ports/
      application/
        usecases/
        commands/
        responses/
        queries/
      infrastructure/
        persistence/
          repositories/
          mappers/
      interfaces/
```

---

## 🧱 Capas de la Arquitectura

### Domain

- Contiene lógica de negocio pura
- Sin dependencias externas

Incluye:

- `aggregates/`
- `valueobjects/`
- `ports/` (interfaces)
- `services/` (lógica compleja sin DB)

---

### Application

- Orquesta casos de uso

Incluye:

- `usecases/`
- `commands/`
- `responses/`
- `queries/` (interfaces, no sqlc)

---

### Infrastructure

- Implementaciones técnicas

Incluye:

- `sqlc` (en shared)
- `repositories/`
- `mappers/`
- `services/`

---

### Interfaces

- Adaptadores (Wails)

Incluye:

- handlers
- transformación DTO ↔ commands
- acceso a SessionState

---

## 🗄️ Base de Datos (SQLite)

### Principios clave

- Multitenant (`organization_id` en todo)
- Soft delete (`deleted_at`)
- Constraints fuertes (integridad)
- Índices para performance
- DB garantiza unicidad (no el use case)

---

### Separación por Contexto

#### IAM

- operator
- organization
- app_state
- member
- team
- team_member

#### Catalog

- product (SIN stock)
- category
- supplier

#### Inventory

- inventory_item (estado actual)
- stock_movement (historial)

#### Sales

- ticket (sale)
- ticket_detail (sale_item)

#### Credit

- customer
- credit_payment

---

## ⚠️ Reglas Arquitectónicas Críticas

1. **Product NO tiene stock**
2. **Inventory maneja TODO el stock**
3. **Contextos no comparten agregados**
4. **Comunicación entre contextos solo vía:**
   - IDs
   - eventos

5. **sqlc solo se usa en infrastructure**
6. **Use cases no instancian dependencias**
7. **Handlers no contienen lógica de negocio**
8. **Unicidad se maneja en DB (constraints)**

---

## 📡 Eventos de Dominio

Ejemplo:

```
SaleCompleted
  → Inventory descuenta stock
  → Credit registra deuda
```

---

## 🔒 Session & App State

- `AppState` (dominio): IDs activos
- `SessionState` (app layer): agregados cargados

Uso:

- handlers acceden a `SessionState`
- no se consulta DB para permisos

---

## 🎯 Objetivo del Sistema

- Offline-first
- Alta coherencia de dominio
- Escalable a múltiples módulos
- Preparado para concurrencia local (POS)
- Mantenible a largo plazo

---

## 🚀 Contexto para la IA

Este proyecto:

- Usa DDD real (no superficial)
- Tiene separación estricta por bounded contexts
- Usa sqlc (no ORM)
- Prioriza invariantes de dominio
- Está optimizado para POS real (ventas rápidas, consistencia)

Cualquier cambio debe:

- respetar bounded contexts
- no acoplar dominios
- mantener separación de capas
- evitar lógica en handlers
- evitar acceso directo a DB fuera de infrastructure
