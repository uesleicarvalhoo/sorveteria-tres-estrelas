-- popsicle
CREATE TABLE popsicles (
	"id" uuid NOT NULL,
	"flavor" varchar UNIQUE NULL,
	"price" numeric NULL,
	CONSTRAINT popsicles_pkey PRIMARY KEY (id)
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