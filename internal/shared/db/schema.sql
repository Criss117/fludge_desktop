-- ============================================================
-- SCHEMA LOCAL MULTITENANT (CORREGIDO)
-- ============================================================

-- =========================
-- IAM
-- =========================

CREATE TABLE IF NOT EXISTS "operator" (
    "id"            text    PRIMARY KEY NOT NULL,
    "name"          text    NOT NULL,
    "email"         text    NOT NULL,
    "username"      text    NOT NULL,
    "pin"           text    NOT NULL,
    "is_root"       integer NOT NULL DEFAULT false,
    "created_at"    integer NOT NULL,
    "updated_at"    integer NOT NULL,
    "deleted_at"    integer
);

CREATE UNIQUE INDEX IF NOT EXISTS "operator_email_unique" ON "operator" ("email");
CREATE UNIQUE INDEX IF NOT EXISTS "operator_username_unique" ON "operator" ("username");


CREATE TABLE IF NOT EXISTS "organization" (
    "id"            text    PRIMARY KEY NOT NULL,
    "name"          text    NOT NULL,
    "slug"          text    NOT NULL,
    "logo"          text,
    "metadata"      blob,
    "legal_name"    text    NOT NULL,
    "address"       text    NOT NULL,
    "contact_phone" text,
    "contact_email" text,
    "created_at"    integer NOT NULL,
    "updated_at"    integer NOT NULL,
    "deleted_at"    integer
);

CREATE UNIQUE INDEX IF NOT EXISTS "organization_slug_uidx"       ON "organization" ("slug");
CREATE UNIQUE INDEX IF NOT EXISTS "organization_legal_name_uidx" ON "organization" ("legal_name");


CREATE TABLE IF NOT EXISTS "app_state" (
    "id"                     text PRIMARY KEY NOT NULL DEFAULT 'local',
    "active_organization_id" text,
    "active_operator_id"     text,
    "updated_at"             integer NOT NULL,
    FOREIGN KEY ("active_organization_id") REFERENCES "organization"("id") ON DELETE SET NULL,
    FOREIGN KEY ("active_operator_id")     REFERENCES "operator"("id")     ON DELETE SET NULL,
    CONSTRAINT "app_state_singleton" CHECK("id" = 'local')
);


CREATE TABLE IF NOT EXISTS "member" (
    "id"              text PRIMARY KEY NOT NULL,
    "organization_id" text NOT NULL,
    "operator_id"     text NOT NULL,
    "role"            text NOT NULL,
    "created_at"      integer NOT NULL,
    "updated_at"      integer NOT NULL,
    "deleted_at"      integer,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE,
    FOREIGN KEY ("operator_id")     REFERENCES "operator"("id")     ON DELETE CASCADE,
    CONSTRAINT "member_role_valid" CHECK("role" IN ('ROOT', 'MEMBER'))
);


CREATE TABLE IF NOT EXISTS "team" (
    "id"              text PRIMARY KEY NOT NULL,
    "name"            text NOT NULL,
    "organization_id" text NOT NULL,
    "permissions"     blob NOT NULL,
    "description"     text,
    "created_at"      integer NOT NULL,
    "updated_at"      integer NOT NULL,
    "deleted_at"      integer,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "team_member" (
    "id"              text PRIMARY KEY NOT NULL,
    "team_id"         text NOT NULL,
    "operator_id"     text NOT NULL,
    "organization_id" text NOT NULL,
    "created_at"      integer NOT NULL,
    "updated_at"      integer NOT NULL,
    "deleted_at"      integer,
    FOREIGN KEY ("team_id")         REFERENCES "team"("id")         ON DELETE CASCADE,
    FOREIGN KEY ("operator_id")     REFERENCES "operator"("id")     ON DELETE CASCADE,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE
);


-- =========================
-- CATALOG
-- =========================

CREATE TABLE IF NOT EXISTS "category" (
    "id"              text PRIMARY KEY NOT NULL,
    "name"            text NOT NULL,
    "description"     text,
    "organization_id" text NOT NULL,
    "created_at"      integer NOT NULL,
    "updated_at"      integer NOT NULL,
    "deleted_at"      integer,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE,
    CONSTRAINT "category_name_min_length" CHECK(length(trim("name")) >= 2)
);

CREATE UNIQUE INDEX IF NOT EXISTS "category_name_organization_idx"
ON "category" ("name", "organization_id");


CREATE TABLE IF NOT EXISTS "supplier" (
    "id"              text PRIMARY KEY NOT NULL,
    "name"            text NOT NULL,
    "contact_phone"   text,
    "contact_email"   text,
    "organization_id" text NOT NULL,
    "metadata"        text,
    "created_at"      integer NOT NULL,
    "updated_at"      integer NOT NULL,
    "deleted_at"      integer,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS "supplier_name_org_unique"
ON "supplier" ("name", "organization_id");


CREATE TABLE IF NOT EXISTS "product" (
    "id"              text PRIMARY KEY NOT NULL,
    "sku"             text NOT NULL,
    "name"            text NOT NULL,
    "description"     text,
    "wholesale_price" integer NOT NULL,
    "sale_price"      integer NOT NULL,
    "cost_price"      integer NOT NULL,
    "category_id"     text,
    "supplier_id"     text,
    "organization_id" text NOT NULL,
    "created_at"      integer NOT NULL,
    "updated_at"      integer NOT NULL,
    "deleted_at"      integer,
    FOREIGN KEY ("category_id")     REFERENCES "category"("id")     ON DELETE NO ACTION,
    FOREIGN KEY ("supplier_id")     REFERENCES "supplier"("id")     ON DELETE NO ACTION,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE,
    CONSTRAINT "sale_price_covers_cost" CHECK("sale_price" >= "cost_price")
);

CREATE UNIQUE INDEX IF NOT EXISTS "product_sku_org_uidx"
ON "product" ("sku", "organization_id");

CREATE UNIQUE INDEX IF NOT EXISTS "product_name_org_uidx"
ON "product" ("name", "organization_id");


-- =========================
-- INVENTORY
-- =========================

CREATE TABLE IF NOT EXISTS "inventory_item" (
    "product_id"      text NOT NULL,
    "organization_id" text NOT NULL,
    "quantity"        integer NOT NULL DEFAULT 0,
    "min_stock"       integer NOT NULL DEFAULT 0,
    "created_at"      integer NOT NULL,
    "updated_at"      integer NOT NULL,
    PRIMARY KEY ("product_id", "organization_id"),
    FOREIGN KEY ("product_id")      REFERENCES "product"("id")      ON DELETE CASCADE,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE,
    CONSTRAINT "quantity_gte_zero"  CHECK("quantity"  >= 0),
    CONSTRAINT "min_stock_gte_zero" CHECK("min_stock" >= 0)
);


CREATE TABLE IF NOT EXISTS "stock_movement" (
    "id"              text PRIMARY KEY NOT NULL,
    "product_id"      text NOT NULL,
    "type"            text NOT NULL,
    "quantity"        integer NOT NULL,
    "reason"          text,
    "organization_id" text NOT NULL,
    "metadata"        text,
    "created_at"      integer NOT NULL,
    FOREIGN KEY ("product_id")      REFERENCES "product"("id"),
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id"),
    CONSTRAINT "quantity_not_zero" CHECK("quantity" != 0)
);


-- =========================
-- SALES (POS)
-- =========================

CREATE TABLE IF NOT EXISTS "ticket" (
    "id"              text PRIMARY KEY NOT NULL,
    "customer_id"     text,
    "total_amount"    integer NOT NULL,
    "paid_amount"     integer NOT NULL,
    "status"          text NOT NULL,
    "payment_method"  text NOT NULL,
    "organization_id" text NOT NULL,
    "operator_id"     text NOT NULL,
    "created_at"      integer NOT NULL,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id"),
    FOREIGN KEY ("operator_id")     REFERENCES "operator"("id")
);


CREATE TABLE IF NOT EXISTS "ticket_detail" (
    "id"              text PRIMARY KEY NOT NULL,
    "ticket_id"       text NOT NULL,
    "product_id"      text,
    "sku"             text NOT NULL,
    "name"            text NOT NULL,
    "sale_price"      integer NOT NULL,
    "quantity"        integer NOT NULL,
    "organization_id" text NOT NULL,
    FOREIGN KEY ("ticket_id") REFERENCES "ticket"("id") ON DELETE CASCADE
);


-- =========================
-- CREDIT
-- =========================

CREATE TABLE IF NOT EXISTS "customer" (
    "id"              text PRIMARY KEY NOT NULL,
    "name"            text NOT NULL,
    "email"           text,
    "phone"           text,
    "address"         text,
    "credit_limit"    integer NOT NULL DEFAULT 0,
    "current_balance" integer NOT NULL DEFAULT 0,
    "organization_id" text NOT NULL,
    "created_at"      integer NOT NULL,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id"),
    CONSTRAINT "balance_within_limit" CHECK("current_balance" <= "credit_limit")
);


CREATE TABLE IF NOT EXISTS "credit_payment" (
    "id"              text PRIMARY KEY NOT NULL,
    "customer_id"     text NOT NULL,
    "amount"          integer NOT NULL,
    "organization_id" text NOT NULL,
    "created_at"      integer NOT NULL,
    FOREIGN KEY ("customer_id")     REFERENCES "customer"("id"),
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id"),
    CONSTRAINT "amount_gt_zero" CHECK("amount" > 0)
);