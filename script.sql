-- public.clientes definition

-- Drop table

-- DROP TABLE public.clientes;

CREATE TABLE public.clientes (
	id bigserial NOT NULL,
	limite int8 NULL,
	saldo int8 NULL,
	CONSTRAINT clientes_pkey PRIMARY KEY (id)
);

-- public.transacaos definition

-- Drop table

-- DROP TABLE public.transacaos;

CREATE TABLE public.transacaos (
	id bigserial NOT NULL,
	cliente_id int8 NULL,
	valor int8 NULL,
	tipo varchar(1) NULL,
	descricao text NULL,
	realizado_em timestamptz NULL,
	CONSTRAINT transacaos_pkey PRIMARY KEY (id)
);

-- public.transacaos foreign keys

ALTER TABLE public.transacaos ADD CONSTRAINT fk_clientes_transacoes FOREIGN KEY (cliente_id) REFERENCES public.clientes(id);

-- insert data

INSERT INTO public.clientes (limite,saldo) VALUES
	 (8000000,0),
	 (100000000,0),
	 (1000000000,0),
	 (50000000,0),
	 (10000000,0);
