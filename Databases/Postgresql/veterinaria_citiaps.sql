--
-- PostgreSQL database dump
--

-- Dumped from database version 15.4
-- Dumped by pg_dump version 15.4

-- Started on 2023-08-29 13:42:14

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
-- TOC entry 215 (class 1259 OID 16497)
-- Name: duenos; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.duenos (
    id integer NOT NULL,
    nombre character varying(50) NOT NULL,
    edad integer NOT NULL,
    sexo character varying(20) NOT NULL
);


ALTER TABLE public.duenos OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 16496)
-- Name: duenos_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.duenos_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.duenos_id_seq OWNER TO postgres;

--
-- TOC entry 3358 (class 0 OID 0)
-- Dependencies: 214
-- Name: duenos_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.duenos_id_seq OWNED BY public.duenos.id;


--
-- TOC entry 221 (class 1259 OID 16528)
-- Name: imagenes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.imagenes (
    id integer NOT NULL,
    url character varying(150) NOT NULL,
    perro_id integer NOT NULL
);


ALTER TABLE public.imagenes OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 16527)
-- Name: imagenes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.imagenes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.imagenes_id_seq OWNER TO postgres;

--
-- TOC entry 3359 (class 0 OID 0)
-- Dependencies: 220
-- Name: imagenes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.imagenes_id_seq OWNED BY public.imagenes.id;


--
-- TOC entry 217 (class 1259 OID 16504)
-- Name: perros; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.perros (
    id integer NOT NULL,
    nombre character varying(50) NOT NULL,
    raza character varying(50) NOT NULL,
    color character varying(50) NOT NULL,
    edad integer NOT NULL,
    dueno_id integer NOT NULL
);


ALTER TABLE public.perros OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 16503)
-- Name: perros_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.perros_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.perros_id_seq OWNER TO postgres;

--
-- TOC entry 3360 (class 0 OID 0)
-- Dependencies: 216
-- Name: perros_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.perros_id_seq OWNED BY public.perros.id;


--
-- TOC entry 219 (class 1259 OID 16516)
-- Name: vacunas; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.vacunas (
    id integer NOT NULL,
    nombre character varying(50) NOT NULL,
    fecha date NOT NULL,
    perro_id integer NOT NULL
);


ALTER TABLE public.vacunas OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 16515)
-- Name: vacunas_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.vacunas_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.vacunas_id_seq OWNER TO postgres;

--
-- TOC entry 3361 (class 0 OID 0)
-- Dependencies: 218
-- Name: vacunas_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.vacunas_id_seq OWNED BY public.vacunas.id;


--
-- TOC entry 3188 (class 2604 OID 16500)
-- Name: duenos id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.duenos ALTER COLUMN id SET DEFAULT nextval('public.duenos_id_seq'::regclass);


--
-- TOC entry 3191 (class 2604 OID 16531)
-- Name: imagenes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.imagenes ALTER COLUMN id SET DEFAULT nextval('public.imagenes_id_seq'::regclass);


--
-- TOC entry 3189 (class 2604 OID 16507)
-- Name: perros id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.perros ALTER COLUMN id SET DEFAULT nextval('public.perros_id_seq'::regclass);


--
-- TOC entry 3190 (class 2604 OID 16519)
-- Name: vacunas id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vacunas ALTER COLUMN id SET DEFAULT nextval('public.vacunas_id_seq'::regclass);


--
-- TOC entry 3346 (class 0 OID 16497)
-- Dependencies: 215
-- Data for Name: duenos; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.duenos (id, nombre, edad, sexo) FROM stdin;
1	Shaggy	20	Masculino
2	Tomás Pérez	31	Masculino
3	Pierre Nodoyuna	41	Masculino
4	Bart Simpson	10	Masculino
5	Juana Gomez	29	Femenino
\.


--
-- TOC entry 3352 (class 0 OID 16528)
-- Dependencies: 221
-- Data for Name: imagenes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.imagenes (id, url, perro_id) FROM stdin;
1	/imagenes/4/_107435678_perro1.jpg	4
2	/imagenes/6/_107435679_perro2.jpg	6
3	/imagenes/5/4ba1557b-5a6b-4321-8521-ba6aa7673018_16-9-discover_1200x675.jpg	5
4	/imagenes/3/gettyimages-57350310-64884f6c47965.jpg	3
5	/imagenes/1/8b5f41e3-ead3-4eea-a247-ee65aa086152_16-9-aspect-r.jpg	1
6	/imagenes/2/f.elconfidencial.com_original_8ec_08c_85c_8ec08c85c866ccb70c4f1c36492d890f.jpg	2
11	/imagenes/4/carlino-a.jpg	4
13	/imagenes/4/pitufilogo.jpg	4
14	/imagenes/4/WhatsApp Image 2023-08-07 at 15.41.23.jpeg	4
\.


--
-- TOC entry 3348 (class 0 OID 16504)
-- Dependencies: 217
-- Data for Name: perros; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.perros (id, nombre, raza, color, edad, dueno_id) FROM stdin;
6	Bobby	Quiltro	Gris y blanco	1	5
5	Cachupín	Quiltro	Gris	1	5
3	Patán	Quiltro	Negro	8	3
1	Scooby Doo	Dogo Alemán	Café	4	1
2	Spike	Bulldog	Gris	2	2
4	Ayudante de Santa	Galgo	Blanco	3	4
\.


--
-- TOC entry 3350 (class 0 OID 16516)
-- Dependencies: 219
-- Data for Name: vacunas; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.vacunas (id, nombre, fecha, perro_id) FROM stdin;
9	Parvovirus	2023-08-05	6
8	Parvovirus	2023-08-05	5
6	Polivalente	2023-08-14	3
5	Polivalente	2019-03-20	3
4	Rabia	2016-03-15	3
1	Parvovirus	2020-01-01	1
2	Polivalente	2020-12-18	1
3	Polivalente	2022-09-11	2
7	Rabia	2018-05-25	4
\.


--
-- TOC entry 3362 (class 0 OID 0)
-- Dependencies: 214
-- Name: duenos_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.duenos_id_seq', 5, true);


--
-- TOC entry 3363 (class 0 OID 0)
-- Dependencies: 220
-- Name: imagenes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.imagenes_id_seq', 14, true);


--
-- TOC entry 3364 (class 0 OID 0)
-- Dependencies: 216
-- Name: perros_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.perros_id_seq', 6, true);


--
-- TOC entry 3365 (class 0 OID 0)
-- Dependencies: 218
-- Name: vacunas_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.vacunas_id_seq', 9, true);


--
-- TOC entry 3193 (class 2606 OID 16502)
-- Name: duenos duenos_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.duenos
    ADD CONSTRAINT duenos_pkey PRIMARY KEY (id);


--
-- TOC entry 3199 (class 2606 OID 16533)
-- Name: imagenes imagenes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.imagenes
    ADD CONSTRAINT imagenes_pkey PRIMARY KEY (id);


--
-- TOC entry 3195 (class 2606 OID 16509)
-- Name: perros perros_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.perros
    ADD CONSTRAINT perros_pkey PRIMARY KEY (id);


--
-- TOC entry 3197 (class 2606 OID 16521)
-- Name: vacunas vacunas_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vacunas
    ADD CONSTRAINT vacunas_pkey PRIMARY KEY (id);


--
-- TOC entry 3202 (class 2606 OID 16534)
-- Name: imagenes imagenes_perro_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.imagenes
    ADD CONSTRAINT imagenes_perro_id_fkey FOREIGN KEY (perro_id) REFERENCES public.perros(id) ON DELETE CASCADE;


--
-- TOC entry 3200 (class 2606 OID 16510)
-- Name: perros perros_dueno_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.perros
    ADD CONSTRAINT perros_dueno_id_fkey FOREIGN KEY (dueno_id) REFERENCES public.duenos(id) ON DELETE CASCADE;


--
-- TOC entry 3201 (class 2606 OID 16522)
-- Name: vacunas vacunas_perro_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vacunas
    ADD CONSTRAINT vacunas_perro_id_fkey FOREIGN KEY (perro_id) REFERENCES public.perros(id) ON DELETE CASCADE;


-- Completed on 2023-08-29 13:42:15

--
-- PostgreSQL database dump complete
--

