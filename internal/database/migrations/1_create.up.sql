-- products
CREATE TABLE products (
	"id" uuid NOT NULL,
	"name" varchar NULL,
	"price_varejo" numeric NULL,
	"price_atacado" numeric NULL,
	"atacado_amount" int8 NULL,
	CONSTRAINT products_pkey PRIMARY KEY (id)
);

-- sales

CREATE TABLE sales (
	"id" uuid NOT NULL,
	"payment_type" varchar NULL,
	"total" numeric NULL,
	"description" varchar NULL,
	"date" date NULL,
	CONSTRAINT sales_pkey PRIMARY KEY (id)
);

-- sale items

CREATE TABLE sale_items (
	"id" uuid NOT NULL,
	"sale_id" uuid NOT NULL,
	"name" varchar NOT NULL,
	"price" numeric NOT NULL,
	"amount" int8 NOT NULL,
	CONSTRAINT sale_items_pkey PRIMARY KEY (id)
);

ALTER TABLE "sale_items" ADD CONSTRAINT fk_sales_items FOREIGN KEY (sale_id) REFERENCES sales(id);

-- users

CREATE TABLE users (
	"id" uuid NOT NULL,
	"name" varchar UNIQUE NOT NULL,
	"email" varchar UNIQUE NOT NULL,
	"password_hash" varchar NOT NULL,
	"permissions" varchar NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);


-- payments

CREATE TABLE payments (
	"id" text NOT NULL,
	"value" numeric NULL,
	"description" text NULL,
	"created_at" timestamptz NULL,
	CONSTRAINT payments_pkey PRIMARY KEY (id)
);

CREATE INDEX payments_created_at_idx ON public.payments USING btree (created_at);