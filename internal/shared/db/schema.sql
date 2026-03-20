-- ============================================================
-- SCHEMA LOCAL MULTITENANT (sin auth)
-- ============================================================

CREATE TABLE IF NOT EXISTS "operator" (
    "id"            text    PRIMARY KEY NOT NULL,
    "name"          text    NOT NULL,
    "email"         text    NOT NULL,
    "username"      text    NOT NULL,
    "pin"           text    NOT NULL,
    "is_root"       integer NOT NULL DEFAULT false,
    "created_at"    integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "updated_at"    integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
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
    "created_at"    integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "updated_at"    integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "deleted_at"    integer
);

CREATE UNIQUE INDEX IF NOT EXISTS "organization_slug_uidx"       ON "organization" ("slug");
CREATE UNIQUE INDEX IF NOT EXISTS "organization_legal_name_uidx" ON "organization" ("legal_name");


CREATE TABLE IF NOT EXISTS "app_state" (
    "id"                     text PRIMARY KEY NOT NULL DEFAULT 'local',
    "active_organization_id" text,
    "active_operator_id"     text,
    "updated_at"             integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    FOREIGN KEY ("active_organization_id") REFERENCES "organization"("id") ON DELETE SET NULL,
    FOREIGN KEY ("active_operator_id")     REFERENCES "operator"("id")     ON DELETE SET NULL,
    CONSTRAINT "app_state_singleton" CHECK("id" = 'local')
);

INSERT OR IGNORE INTO "app_state" ("id", "updated_at")
VALUES ('local', cast(unixepoch('subsecond') * 1000 as integer));


CREATE TABLE IF NOT EXISTS "member" (
    "id"              text    PRIMARY KEY NOT NULL,
    "organization_id" text    NOT NULL,
    "operator_id"     text    NOT NULL,
    "role"            text    NOT NULL,
    "created_at"      integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "updated_at"      integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "deleted_at"      integer,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE,
    FOREIGN KEY ("operator_id")     REFERENCES "operator"("id")     ON DELETE CASCADE,
    CONSTRAINT "member_role_valid" CHECK("role" IN ('ROOT', 'MEMBER'))
);



CREATE TABLE IF NOT EXISTS "team" (
    "id"              text  PRIMARY KEY NOT NULL,
    "name"            text  NOT NULL,
    "organization_id" text  NOT NULL,
    "permissions"     blob  NOT NULL,
    "description"     text,
    "created_at"      integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "updated_at"      integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "deleted_at"      integer,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS "team_organization_id_idx" ON "team" ("organization_id");


CREATE TABLE IF NOT EXISTS "team_member" (
    "id"              text    PRIMARY KEY NOT NULL,
    "team_id"         text    NOT NULL,
    "operator_id"     text    NOT NULL,
    "organization_id" text    NOT NULL,
    "created_at"      integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "updated_at"      integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "deleted_at"      integer,
    FOREIGN KEY ("team_id")         REFERENCES "team"("id")         ON DELETE CASCADE,
    FOREIGN KEY ("operator_id")     REFERENCES "operator"("id")     ON DELETE CASCADE,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS "customer" (
    "id"              text    PRIMARY KEY NOT NULL,
    "name"            text    NOT NULL,
    "email"           text,
    "phone"           text,
    "address"         text,
    "credit_limit"    integer DEFAULT 0 NOT NULL,
    "current_balance" integer DEFAULT 0 NOT NULL,
    "organization_id" text    NOT NULL,
    "created_at"      integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "updated_at"      integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "deleted_at"      integer,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE,
    CONSTRAINT "credit_limit_gte_zero"    CHECK("credit_limit"    >= 0),
    CONSTRAINT "current_balance_gte_zero" CHECK("current_balance" >= 0),
    CONSTRAINT "balance_within_limit"     CHECK("current_balance" <= "credit_limit"),
    CONSTRAINT "phone_format"             CHECK("phone" IS NULL OR (length("phone") >= 7 AND "phone" GLOB '[0-9+-]*'))
);

CREATE INDEX IF NOT EXISTS "customer_org_idx"  ON "customer" ("organization_id");
CREATE INDEX IF NOT EXISTS "customer_name_idx" ON "customer" ("name");


CREATE TABLE IF NOT EXISTS "credit_payment" (
    "id"              text    PRIMARY KEY NOT NULL,
    "customer_id"     text    NOT NULL,
    "amount"          integer NOT NULL,
    "payment_method"  text,
    "organization_id" text    NOT NULL,
    "metadata"        text,
    "created_at"      integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "updated_at"      integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "deleted_at"      integer,
    FOREIGN KEY ("customer_id")     REFERENCES "customer"("id")     ON DELETE CASCADE,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE,
    CONSTRAINT "amount_gt_zero"       CHECK("amount" > 0),
    CONSTRAINT "metadata_valid_json"  CHECK(json_valid("metadata"))
);

CREATE INDEX IF NOT EXISTS "payment_customer_idx" ON "credit_payment" ("customer_id");
CREATE INDEX IF NOT EXISTS "payment_org_idx"      ON "credit_payment" ("organization_id");


CREATE TABLE IF NOT EXISTS "category" (
    "id"              text    PRIMARY KEY NOT NULL,
    "name"            text    NOT NULL,
    "description"     text,
    "organization_id" text    NOT NULL,
    "created_at"      integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "updated_at"      integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "deleted_at"      integer,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE,
    CONSTRAINT "category_name_min_length" CHECK(length(trim("name")) >= 2)
);

CREATE INDEX        IF NOT EXISTS "category_name_idx"              ON "category" ("name");
CREATE INDEX        IF NOT EXISTS "category_organizationId_idx"    ON "category" ("organization_id");
CREATE UNIQUE INDEX IF NOT EXISTS "category_name_organization_idx" ON "category" ("name", "organization_id");


CREATE TABLE IF NOT EXISTS "supplier" (
    "id"              text    PRIMARY KEY NOT NULL,
    "name"            text    NOT NULL,
    "contact_phone"   text,
    "contact_email"   text,
    "organization_id" text    NOT NULL,
    "metadata"        text,
    "created_at"      integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "updated_at"      integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "deleted_at"      integer,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE,
    CONSTRAINT "metadata_valid_json"    CHECK(json_valid("metadata")),
    CONSTRAINT "contact_phone_format"   CHECK("contact_phone" IS NULL OR (length("contact_phone") >= 7 AND "contact_phone" GLOB '[0-9+-]*')),
    CONSTRAINT "contact_email_format"   CHECK("contact_email" IS NULL OR ("contact_email" GLOB '*@*.*' AND "contact_email" NOT LIKE '% %'))
);

CREATE INDEX        IF NOT EXISTS "supplier_org_idx"                    ON "supplier" ("organization_id");
CREATE INDEX        IF NOT EXISTS "supplier_name_idx"                   ON "supplier" ("name");
CREATE UNIQUE INDEX IF NOT EXISTS "supplier_name_organizationId_unique" ON "supplier" ("name", "organization_id");


CREATE TABLE IF NOT EXISTS "product" (
    "id"              text    PRIMARY KEY NOT NULL,
    "sku"             text    NOT NULL,
    "name"            text    NOT NULL,
    "description"     text,
    "wholesale_price" integer NOT NULL,
    "sale_price"      integer NOT NULL,
    "cost_price"      integer NOT NULL,
    "stock"           integer NOT NULL,
    "min_stock"       integer NOT NULL,
    "category_id"     text,
    "organization_id" text    NOT NULL,
    "supplier_id"     text,
    "created_at"      integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "updated_at"      integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "deleted_at"      integer,
    FOREIGN KEY ("category_id")     REFERENCES "category"("id")     ON DELETE NO ACTION,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE,
    FOREIGN KEY ("supplier_id")     REFERENCES "supplier"("id")     ON DELETE NO ACTION,
    CONSTRAINT "wholesale_price_gte_zero" CHECK("wholesale_price" >= 0),
    CONSTRAINT "sale_price_gte_zero"      CHECK("sale_price"      >= 0),
    CONSTRAINT "cost_price_gte_zero"      CHECK("cost_price"      >= 0),
    CONSTRAINT "reorder_level_gte_zero"   CHECK("min_stock"       >= 0),
    CONSTRAINT "sale_price_covers_cost"   CHECK("sale_price"      >= "cost_price"),
    CONSTRAINT "stock_gte_zero"           CHECK("stock"           >= 0)
);

CREATE INDEX        IF NOT EXISTS "sku_idx"                  ON "product" ("sku");
CREATE INDEX        IF NOT EXISTS "product_name_idx"         ON "product" ("name");
CREATE INDEX        IF NOT EXISTS "product_category_idx"     ON "product" ("category_id");
CREATE INDEX        IF NOT EXISTS "product_supplier_idx"     ON "product" ("supplier_id");
CREATE INDEX        IF NOT EXISTS "product_organization_idx" ON "product" ("organization_id");
CREATE INDEX        IF NOT EXISTS "product_stock_idx"        ON "product" ("stock");
CREATE UNIQUE INDEX IF NOT EXISTS "sku_organization_id_idx"  ON "product" ("sku", "organization_id");
CREATE UNIQUE INDEX IF NOT EXISTS "name_organization_id_idx" ON "product" ("name", "organization_id");


CREATE TABLE IF NOT EXISTS "stock_movement" (
    "id"              text    PRIMARY KEY NOT NULL,
    "product_id"      text    NOT NULL,
    "type"            text    NOT NULL,
    "quantity"        integer NOT NULL,
    "reason"          text,
    "organization_id" text    NOT NULL,
    "metadata"        text,
    "created_at"      integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "updated_at"      integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "deleted_at"      integer,
    FOREIGN KEY ("product_id")      REFERENCES "product"("id")      ON DELETE NO ACTION,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE,
    CONSTRAINT "quantity_not_zero"   CHECK("quantity" != 0),
    CONSTRAINT "metadata_valid_json" CHECK(json_valid("metadata"))
);

CREATE INDEX IF NOT EXISTS "stock_movement_product_idx" ON "stock_movement" ("product_id");
CREATE INDEX IF NOT EXISTS "stock_movement_org_idx"     ON "stock_movement" ("organization_id");
CREATE INDEX IF NOT EXISTS "stock_movement_type_idx"    ON "stock_movement" ("type");


CREATE TABLE IF NOT EXISTS "ticket" (
    "id"              text    PRIMARY KEY NOT NULL,
    "customer_id"     text,
    "total_amount"    integer DEFAULT 0 NOT NULL,
    "paid_amount"     integer DEFAULT 0 NOT NULL,
    "status"          text    DEFAULT 'PENDING' NOT NULL,
    "payment_method"  text    NOT NULL,
    "organization_id" text    NOT NULL,
    "operator_id"     text    NOT NULL,
    "created_at"      integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "updated_at"      integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "deleted_at"      integer,
    FOREIGN KEY ("customer_id")     REFERENCES "customer"("id")     ON DELETE SET NULL,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE,
    FOREIGN KEY ("operator_id")     REFERENCES "operator"("id")     ON DELETE SET NULL,
    CONSTRAINT "total_amount_gt_zero"         CHECK("total_amount" > 0),
    CONSTRAINT "paid_amount_gte_zero"         CHECK("paid_amount"  >= 0),
    CONSTRAINT "paid_amount_lte_total_amount" CHECK("paid_amount"  <= "total_amount")
);

CREATE INDEX IF NOT EXISTS "ticket_customer_idx" ON "ticket" ("customer_id");
CREATE INDEX IF NOT EXISTS "ticket_status_idx"   ON "ticket" ("status");
CREATE INDEX IF NOT EXISTS "ticket_org_idx"      ON "ticket" ("organization_id");
CREATE INDEX IF NOT EXISTS "ticket_operator_idx" ON "ticket" ("operator_id");


CREATE TABLE IF NOT EXISTS "ticket_detail" (
    "id"              text    PRIMARY KEY NOT NULL,
    "ticket_id"       text    NOT NULL,
    "sku"             text    NOT NULL,
    "name"            text    NOT NULL,
    "description"     text,
    "sale_price"      integer NOT NULL,
    "quantity"        integer NOT NULL,
    "sub_total"       integer GENERATED ALWAYS AS ("sale_price" * "quantity") VIRTUAL NOT NULL,
    "product_id"      text,
    "organization_id" text    NOT NULL,
    "created_at"      integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "updated_at"      integer DEFAULT (cast(unixepoch('subsecond') * 1000 as integer)) NOT NULL,
    "deleted_at"      integer,
    FOREIGN KEY ("ticket_id")       REFERENCES "ticket"("id")       ON DELETE CASCADE,
    FOREIGN KEY ("product_id")      REFERENCES "product"("id")      ON DELETE SET NULL,
    FOREIGN KEY ("organization_id") REFERENCES "organization"("id") ON DELETE CASCADE,
    CONSTRAINT "quantity_gt_zero"    CHECK("quantity"   > 0),
    CONSTRAINT "sub_total_gte_zero"  CHECK("sub_total"  >= 0),
    CONSTRAINT "sale_price_gte_zero" CHECK("sale_price" >= 0),
    CONSTRAINT "sub_total_correct"   CHECK("sub_total"  = "sale_price" * "quantity")
);

CREATE INDEX IF NOT EXISTS "ticket_detail_ticket_idx"  ON "ticket_detail" ("ticket_id");
CREATE INDEX IF NOT EXISTS "ticket_detail_org_idx"     ON "ticket_detail" ("organization_id");
CREATE INDEX IF NOT EXISTS "ticket_detail_product_idx" ON "ticket_detail" ("product_id");
