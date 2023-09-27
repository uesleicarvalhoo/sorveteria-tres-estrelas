-- transactions

CREATE TABLE transactions (
	"id" text NOT NULL,
	"value" numeric NULL,
	"type" text NULL,
	"description" text NULL,
	"created_at" timestamptz NULL,
	CONSTRAINT transactions_pkey PRIMARY KEY (id)
);

CREATE INDEX transactions_created_at_idx ON public.transactions USING btree (created_at);