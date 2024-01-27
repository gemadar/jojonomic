CREATE TABLE IF NOT EXISTS public.harga
(
    id serial,
	admin_id varchar,
	harga_topup numeric,
	harga_buyback numeric
);

CREATE TABLE IF NOT EXISTS public.topup
(
    id serial,
	gram varchar, 
	harga varchar
);

CREATE TABLE IF NOT EXISTS public.rekening
(
    id serial,
	norek varchar,
	saldo float
);

CREATE TABLE IF NOT EXISTS public.transaksi
(
    id serial,
	date timestamp without time zone,
	type varchar,
	gram float,
	harga_topup numeric,
	harga_buyback numeric,
	norek varchar
);

INSERT INTO public.harga (admin_id, harga_buyback, harga_topup) VALUES (null, null, null);