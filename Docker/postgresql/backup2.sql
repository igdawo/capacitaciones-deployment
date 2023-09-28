--
-- PostgreSQL database dump
--

-- Dumped from database version 15.4
-- Dumped by pg_dump version 15.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: dueno; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.dueno (
    id integer NOT NULL,
    "Nombre" character varying(80),
    "Edad" integer,
    "Sexo" character varying(80)
);


ALTER TABLE public.dueno OWNER TO postgres;

--
-- Name: dueno_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.dueno_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.dueno_id_seq OWNER TO postgres;

--
-- Name: dueno_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.dueno_id_seq OWNED BY public.dueno.id;


--
-- Name: perro; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.perro (
    id integer NOT NULL,
    id_dueno integer,
    "Nombre" character varying(80),
    "Raza" character varying(80),
    "Color" character varying(80),
    "Edad" integer
);


ALTER TABLE public.perro OWNER TO postgres;

--
-- Name: perro_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.perro_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.perro_id_seq OWNER TO postgres;

--
-- Name: perro_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.perro_id_seq OWNED BY public.perro.id;


--
-- Name: vacuna; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.vacuna (
    id integer NOT NULL,
    id_perro integer,
    "Fecha" date,
    "NombreVacuna" character varying(80)
);


ALTER TABLE public.vacuna OWNER TO postgres;

--
-- Name: vacuna_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.vacuna_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.vacuna_id_seq OWNER TO postgres;

--
-- Name: vacuna_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.vacuna_id_seq OWNED BY public.vacuna.id;


--
-- Name: dueno id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dueno ALTER COLUMN id SET DEFAULT nextval('public.dueno_id_seq'::regclass);


--
-- Name: perro id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.perro ALTER COLUMN id SET DEFAULT nextval('public.perro_id_seq'::regclass);


--
-- Name: vacuna id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vacuna ALTER COLUMN id SET DEFAULT nextval('public.vacuna_id_seq'::regclass);


--
-- Data for Name: dueno; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.dueno (id, "Nombre", "Edad", "Sexo") FROM stdin;
1	Miguel Castro	23	Masculino
2	Antonio Salinas	34	Masculino
3	Muriel Cerda	12	Femenino
4	Bastian Huerta	43	Masculino
5	Armando Casas	61	No binario
8	Diego Casas	61	Masculino
7	Esteban Cruces	16	Masculino
9	Marcelo Cruces	12	Masculino
\.


--
-- Data for Name: perro; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.perro (id, id_dueno, "Nombre", "Raza", "Color", "Edad") FROM stdin;
1	1	Satanas	Puddle	Negro	2
2	1	Princeso	Pitbull	Blanco	4
3	4	Axel	Kiltro	Rojo	6
5	3	Axel	Gran danes	Cafe	9
4	2	Invader y destroyer	Yorkshire	Rojo carmesi	10
7	2	cobadonga	Mestizo	Azul	5
8	7	Pepino	Sharpei ingles	Negro	1
9	9	Tomate	Pastor aleman	Verde	12
10	3	kimono	astudillo	negro	10
11	8	camarino	sharpei	azul	12
\.


--
-- Data for Name: vacuna; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.vacuna (id, id_perro, "Fecha", "NombreVacuna") FROM stdin;
1	4	2022-04-15	Antirabica
2	3	2011-06-23	rabia
3	2	2010-10-03	Pulgas
5	5	2003-04-04	Antirabica
6	3	2011-03-12	Covidsar22
4	2	2006-03-08	Antirabica
8	3	2006-05-08	Antirabica v2
11	1	2023-09-12	moquito
12	1	2020-06-23	heptavalente
13	9	2023-09-04	Antirabica
14	1	2023-09-05	hepatitis b
\.


--
-- Name: dueno_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.dueno_id_seq', 9, true);


--
-- Name: perro_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.perro_id_seq', 11, true);


--
-- Name: vacuna_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.vacuna_id_seq', 14, true);


--
-- Name: dueno dueno_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dueno
    ADD CONSTRAINT dueno_pkey PRIMARY KEY (id);


--
-- Name: perro perro_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.perro
    ADD CONSTRAINT perro_pkey PRIMARY KEY (id);


--
-- Name: vacuna vacuna_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vacuna
    ADD CONSTRAINT vacuna_pkey PRIMARY KEY (id);


--
-- Name: perro perro_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.perro
    ADD CONSTRAINT perro_fk FOREIGN KEY (id_dueno) REFERENCES public.dueno(id) ON DELETE CASCADE;


--
-- Name: vacuna vacuna_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vacuna
    ADD CONSTRAINT vacuna_fk FOREIGN KEY (id_perro) REFERENCES public.perro(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

