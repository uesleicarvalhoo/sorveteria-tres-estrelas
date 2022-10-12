-- popsicle
CREATE TABLE popsicles (
	"id" uuid NOT NULL,
	"flavor" varchar NULL,
	"price" numeric NULL,
	CONSTRAINT popsicles_pkey PRIMARY KEY (id)
);

-- sale items

CREATE TABLE sale_items (
	"id" uuid NOT NULL,
	"sale_id" uuid NULL,
	"name" varchar NULL,
	"price" numeric NULL,
	"amount" int8 NULL,
	CONSTRAINT sale_items_pkey PRIMARY KEY (id)
);

ALTER TABLE "sale_items" ADD CONSTRAINT fk_sales_items FOREIGN KEY (sale_id) REFERENCES sales(id);


-- sales

CREATE TABLE sales (
	"id" uuid NOT NULL,
	"payment_type" varchar NULL,
	"total" numeric NULL,
	"description" varchar NULL,
	"date" date NULL,
	CONSTRAINT sales_pkey PRIMARY KEY (id)
);
