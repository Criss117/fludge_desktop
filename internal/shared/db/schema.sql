-- ============================================================
-- SCHEMA: LOCAL MULTITENANT — v2
-- ============================================================
-- Convenciones:
--   • Timestamps: UNIX epoch INTEGER (ms)
--   • Montos:     INTEGER en centavos (ej. $10.50 = 1050)
--   • Borrado:    Soft-delete via deleted_at IS NULL
--   • Índices:    Únicos parciales para registros activos
-- ============================================================

PRAGMA journal_mode = WAL;
PRAGMA foreign_keys = ON;

-- ============================================================
-- SECTION: IAM (Identity & Access Management)
-- ============================================================

CREATE TABLE IF NOT EXISTS "operator" (
    "id"            TEXT    PRIMARY KEY NOT NULL,
    "name"          TEXT    NOT NULL,
    "email"         TEXT    NOT NULL,
    "username"      TEXT    NOT NULL,
    "pin"           TEXT    NOT NULL,
    "operator_type" TEXT    NOT NULL,
    "created_at"    INTEGER NOT NULL,
    "updated_at"    INTEGER NOT NULL,
    "deleted_at"    INTEGER,
    
    CONSTRAINT "operator_type_valid" CHECK("operator_type" IN ('ROOT', 'EMPLOYEE')),
    CONSTRAINT "operator_pin_length" CHECK(length("pin") >= 4)
);

CREATE UNIQUE INDEX IF NOT EXISTS "operator_username_uidx" ON "operator" ("username") WHERE "deleted_at" IS NULL;
CREATE UNIQUE INDEX IF NOT EXISTS "operator_email_uidx"    ON "operator" ("email")    WHERE "deleted_at" IS NULL;

---

CREATE TABLE IF NOT EXISTS "organization" (
    "id"            TEXT    PRIMARY KEY NOT NULL,
    "name"          TEXT    NOT NULL,
    "slug"          TEXT    NOT NULL,
    "logo"          TEXT,
    "metadata"      BLOB,
    "legal_name"    TEXT    NOT NULL,
    "address"       TEXT    NOT NULL,
    "contact_phone" TEXT,
    "contact_email" TEXT,
    "created_at"    INTEGER NOT NULL,
    "updated_at"    INTEGER NOT NULL,
    "deleted_at"    INTEGER
);

CREATE UNIQUE INDEX IF NOT EXISTS "organization_slug_uidx"       ON "organization" ("slug")       WHERE "deleted_at" IS NULL;
CREATE UNIQUE INDEX IF NOT EXISTS "organization_legal_name_uidx" ON "organization" ("legal_name") WHERE "deleted_at" IS NULL;

---

-- Singleton: Control de estado de la aplicación local
CREATE TABLE IF NOT EXISTS "app_state" (
    "id"                     TEXT    PRIMARY KEY NOT NULL DEFAULT 'local',
    "active_organization_id" TEXT,
    "active_operator_id"     TEXT,
    "updated_at"             INTEGER NOT NULL,
    
    FOREIGN KEY ("active_organization_id") REFERENCES "organization"("id") ON DELETE SET NULL,
    FOREIGN KEY ("active_operator_id")     REFERENCES "operator"("id")     ON DELETE SET NULL,
    CONSTRAINT  "app_state_singleton"      CHECK("id" = 'local')
);

---

CREATE TABLE IF NOT EXISTS "member" (
    "id"              TEXT    PRIMARY KEY NOT NULL,
    "organization_id" TEXT    NOT NULL,
    "operator_id"     TEXT    NOT NULL,
    "role"            TEXT    NOT NULL,
    "created_at"      INTEGER NOT NULL,
    "updated_at"      INTEGER NOT NULL,
    "deleted_at"      INTEGER,
    
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE,
    FOREIGN KEY ("operator_id")     REFERENCES "operator"("id")     ON DELETE CASCADE,
    CONSTRAINT  "member_role_valid" CHECK("role" IN ('ROOT', 'EMPLOYEE'))
);

CREATE UNIQUE INDEX IF NOT EXISTS "member_org_operator_uidx" ON "member" ("organization_id", "operator_id") WHERE "deleted_at" IS NULL;

---

CREATE TABLE IF NOT EXISTS "team" (
    "id"              TEXT    PRIMARY KEY NOT NULL,
    "name"            TEXT    NOT NULL,
    "organization_id" TEXT    NOT NULL,
    "permissions"     BLOB    NOT NULL,
    "description"     TEXT,
    "created_at"      INTEGER NOT NULL,
    "updated_at"      INTEGER NOT NULL,
    "deleted_at"      INTEGER,
    
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS "team_name_org_uidx" ON "team" ("name", "organization_id") WHERE "deleted_at" IS NULL;

---

CREATE TABLE IF NOT EXISTS "team_member" (
    "id"              TEXT    PRIMARY KEY NOT NULL,
    "team_id"         TEXT    NOT NULL,
    "operator_id"     TEXT    NOT NULL,
    "organization_id" TEXT    NOT NULL,
    "created_at"      INTEGER NOT NULL,
    "updated_at"      INTEGER NOT NULL,
    "deleted_at"      INTEGER,
    
    FOREIGN KEY ("team_id")         REFERENCES "team"("id")         ON DELETE CASCADE,
    FOREIGN KEY ("operator_id")     REFERENCES "operator"("id")     ON DELETE CASCADE,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS "team_member_team_operator_uidx" ON "team_member" ("team_id", "operator_id") WHERE "deleted_at" IS NULL;


-- ============================================================
-- SECTION: CATALOG
-- ============================================================

CREATE TABLE IF NOT EXISTS "category" (
    "id"              TEXT    PRIMARY KEY NOT NULL,
    "name"            TEXT    NOT NULL,
    "description"     TEXT,
    "organization_id" TEXT    NOT NULL,
    "created_at"      INTEGER NOT NULL,
    "updated_at"      INTEGER NOT NULL,
    "deleted_at"      INTEGER,
    
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE,
    CONSTRAINT  "category_name_min_length" CHECK(length(trim("name")) >= 2)
);

CREATE UNIQUE INDEX IF NOT EXISTS "category_name_org_uidx" ON "category" ("name", "organization_id") WHERE "deleted_at" IS NULL;

---

CREATE TABLE IF NOT EXISTS "supplier" (
    "id"              TEXT    PRIMARY KEY NOT NULL,
    "name"            TEXT    NOT NULL,
    "contact_phone"   TEXT,
    "contact_email"   TEXT,
    "organization_id" TEXT    NOT NULL,
    "metadata"        TEXT,
    "created_at"      INTEGER NOT NULL,
    "updated_at"      INTEGER NOT NULL,
    "deleted_at"      INTEGER,
    
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS "supplier_name_org_uidx" ON "supplier" ("name", "organization_id") WHERE "deleted_at" IS NULL;

---

CREATE TABLE IF NOT EXISTS "product" (
    "id"              TEXT    PRIMARY KEY NOT NULL,
    "sku"             TEXT    NOT NULL,
    "name"            TEXT    NOT NULL,
    "description"     TEXT,
    "wholesale_price" INTEGER NOT NULL, -- Cents
    "sale_price"      INTEGER NOT NULL, -- Cents
    "cost_price"      INTEGER NOT NULL, -- Cents
    "category_id"     TEXT,
    "supplier_id"     TEXT,
    "organization_id" TEXT    NOT NULL,
    "created_at"      INTEGER NOT NULL,
    "updated_at"      INTEGER NOT NULL,
    "deleted_at"      INTEGER,
    
    FOREIGN KEY ("category_id")     REFERENCES "category"("id")     ON DELETE SET NULL,
    FOREIGN KEY ("supplier_id")     REFERENCES "supplier"("id")     ON DELETE SET NULL,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE,
    
    CONSTRAINT "prices_gte_zero"       CHECK("wholesale_price" >= 0 AND "sale_price" >= 0 AND "cost_price" >= 0),
    CONSTRAINT "sale_price_covers_cost" CHECK("sale_price" >= "cost_price")
);

CREATE UNIQUE INDEX IF NOT EXISTS "product_sku_org_uidx"  ON "product" ("sku", "organization_id")  WHERE "deleted_at" IS NULL;
CREATE UNIQUE INDEX IF NOT EXISTS "product_name_org_uidx" ON "product" ("name", "organization_id") WHERE "deleted_at" IS NULL;


-- ============================================================
-- SECTION: CREDIT
-- ============================================================

CREATE TABLE IF NOT EXISTS "customer" (
    "id"              TEXT    PRIMARY KEY NOT NULL,
    "name"            TEXT    NOT NULL,
    "email"           TEXT,
    "phone"           TEXT,
    "address"         TEXT,
    "credit_limit"    INTEGER NOT NULL DEFAULT 0,
    "current_balance" INTEGER NOT NULL DEFAULT 0,
    "organization_id" TEXT    NOT NULL,
    "created_at"      INTEGER NOT NULL,
    "updated_at"      INTEGER NOT NULL,
    
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE RESTRICT,
    CONSTRAINT  "customer_credit_limit_gte_zero"    CHECK("credit_limit"    >= 0),
    CONSTRAINT  "customer_current_balance_gte_zero"  CHECK("current_balance" >= 0),
    CONSTRAINT  "customer_current_balance_lte_limit" CHECK("current_balance" <= "credit_limit")
);

CREATE INDEX IF NOT EXISTS "customer_org_idx" ON "customer" ("organization_id");

---

CREATE TABLE IF NOT EXISTS "credit_payment" (
    "id"              TEXT    PRIMARY KEY NOT NULL,
    "customer_id"     TEXT    NOT NULL,
    "amount"          INTEGER NOT NULL,
    "notes"           TEXT,
    "organization_id" TEXT    NOT NULL,
    "operator_id"     TEXT    NOT NULL,
    "created_at"      INTEGER NOT NULL,
    
    FOREIGN KEY ("customer_id")     REFERENCES "customer"("id")     ON DELETE RESTRICT,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE RESTRICT,
    FOREIGN KEY ("operator_id")     REFERENCES "operator"("id")     ON DELETE RESTRICT,
    CONSTRAINT  "payment_amount_gt_zero" CHECK("amount" > 0)
);

CREATE INDEX IF NOT EXISTS "credit_payment_customer_idx" ON "credit_payment" ("customer_id", "created_at");


-- ============================================================
-- SECTION: SALES (POS)
-- ============================================================

CREATE TABLE IF NOT EXISTS "ticket" (
    "id"              TEXT    PRIMARY KEY NOT NULL,
    "customer_id"     TEXT,
    "total_amount"    INTEGER NOT NULL,
    "paid_amount"     INTEGER NOT NULL,
    "status"          TEXT    NOT NULL,
    "payment_method"  TEXT    NOT NULL,
    "organization_id" TEXT    NOT NULL,
    "operator_id"     TEXT    NOT NULL,
    "created_at"      INTEGER NOT NULL,
    "updated_at"      INTEGER NOT NULL,
    
    FOREIGN KEY ("customer_id")     REFERENCES "customer"("id")     ON DELETE SET NULL,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE RESTRICT,
    FOREIGN KEY ("operator_id")     REFERENCES "operator"("id")     ON DELETE RESTRICT,
    
    CONSTRAINT "ticket_status_valid"         CHECK("status" IN ('OPEN', 'PAID', 'CANCELLED', 'REFUNDED')),
    CONSTRAINT "ticket_payment_method_valid" CHECK("payment_method" IN ('CASH', 'CARD', 'TRANSFER', 'CREDIT', 'MIXED')),
    CONSTRAINT "ticket_total_amount_gte_zero" CHECK("total_amount" >= 0),
    CONSTRAINT "ticket_paid_amount_gte_zero"  CHECK("paid_amount"  >= 0),
    CONSTRAINT "ticket_paid_amount_lte_total" CHECK("paid_amount" <= "total_amount")
);

CREATE INDEX IF NOT EXISTS "ticket_org_created_idx" ON "ticket" ("organization_id", "created_at");
CREATE INDEX IF NOT EXISTS "ticket_customer_idx"    ON "ticket" ("customer_id") WHERE "customer_id" IS NOT NULL;

---

CREATE TABLE IF NOT EXISTS "ticket_detail" (
    "id"              TEXT    PRIMARY KEY NOT NULL,
    "ticket_id"       TEXT    NOT NULL,
    "product_id"      TEXT,
    "sku"             TEXT    NOT NULL,
    "name"            TEXT    NOT NULL,
    "sale_price"      INTEGER NOT NULL,
    "quantity"        INTEGER NOT NULL,
    "organization_id" TEXT    NOT NULL,
    "created_at"      INTEGER NOT NULL,
    "updated_at"      INTEGER NOT NULL,
    
    FOREIGN KEY ("ticket_id")  REFERENCES "ticket"("id")  ON DELETE CASCADE,
    FOREIGN KEY ("product_id") REFERENCES "product"("id") ON DELETE SET NULL,
    
    CONSTRAINT "detail_sale_price_gte_zero" CHECK("sale_price" >= 0),
    CONSTRAINT "detail_quantity_gt_zero"   CHECK("quantity"   > 0)
);

CREATE INDEX IF NOT EXISTS "ticket_detail_product_id_idx" ON "ticket_detail" ("product_id");
CREATE INDEX IF NOT EXISTS "ticket_detail_ticket_idx"     ON "ticket_detail" ("ticket_id");


-- ============================================================
-- SECTION: INVENTORY
-- ============================================================

CREATE TABLE IF NOT EXISTS "inventory_item" (
    "product_id"      TEXT    NOT NULL,
    "organization_id" TEXT    NOT NULL,
    "stock"           INTEGER NOT NULL DEFAULT 0,
    "min_stock"       INTEGER NOT NULL DEFAULT 0,
    "created_at"      INTEGER NOT NULL,
    "updated_at"      INTEGER NOT NULL,
    
    PRIMARY KEY ("product_id", "organization_id"),
    FOREIGN KEY ("product_id")      REFERENCES "product"("id")      ON DELETE CASCADE,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE,
    
    CONSTRAINT "stock_gte_zero"     CHECK("stock" >= 0),
    CONSTRAINT "min_stock_gte_zero" CHECK("min_stock" >= 0)
);

---

CREATE TABLE IF NOT EXISTS "stock_movement" (
    "id"               TEXT    PRIMARY KEY NOT NULL,
    "type"             TEXT    NOT NULL,
    "quantity"         INTEGER NOT NULL,
    "reason"           TEXT,
    "product_id"       TEXT    NOT NULL,
    "organization_id"  TEXT    NOT NULL,
    "ticket_detail_id" TEXT,
    "metadata"         TEXT,
    "created_at"       INTEGER NOT NULL,
    
    FOREIGN KEY ("product_id")       REFERENCES "product"("id")       ON DELETE RESTRICT,
    FOREIGN KEY ("organization_id")  REFERENCES "organization"("id")  ON DELETE RESTRICT,
    FOREIGN KEY ("ticket_detail_id") REFERENCES "ticket_detail"("id") ON DELETE SET NULL,
    
    CONSTRAINT "stock_movement_quantity_gt_zero" CHECK("quantity" > 0),
    CONSTRAINT "stock_movement_type_valid" 
        CHECK("type" IN ('IN', 'OUT', 'ADJUSTMENT_ADD', 'ADJUSTMENT_SUB', 'RETURN', 'LOSS'))
);

CREATE INDEX IF NOT EXISTS "stock_movement_product_idx" ON "stock_movement" ("product_id", "created_at");