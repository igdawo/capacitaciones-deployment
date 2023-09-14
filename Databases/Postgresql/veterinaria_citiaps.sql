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

--
-- Name: make_serial(text, text, text); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.make_serial(p_table_schema text, p_table_name text, p_column_name text) RETURNS void
    LANGUAGE plpgsql
    AS $$
declare
  l_sql text;
  l_seq_name text;
  l_full_name text;
begin
  l_seq_name := concat(p_table_name, '_', p_column_name, '_seq');
  l_full_name := quote_ident(p_table_schema)||'.'||quote_ident(l_seq_name);
  
  execute format('create sequence %I.%I', p_table_schema, l_seq_name);
  execute format('alter table %I.%I alter column %I set default nextval(%L)', p_table_schema, p_table_name, p_column_name, l_seq_name);
  execute format('alter table %I.%I alter column %I set not null', p_table_schema, p_table_name, p_column_name);
  execute format('alter sequence %I.%I owned by %I.%I', p_table_schema, l_seq_name, p_table_name, p_column_name);
  execute format('select setval(%L, coalesce(max(%I),0)) from %I', l_full_name, p_column_name, p_table_name);
end;
$$;


ALTER FUNCTION public.make_serial(p_table_schema text, p_table_name text, p_column_name text) OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: Dueno_Perros; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Dueno_Perros" (
    id integer NOT NULL,
    "Dueno_id" integer NOT NULL,
    "Perro_id" integer NOT NULL
);


ALTER TABLE public."Dueno_Perros" OWNER TO postgres;

--
-- Name: Perro_Vacunas; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Perro_Vacunas" (
    id integer NOT NULL,
    "Perro_id" integer NOT NULL,
    "Vacuna_id" integer NOT NULL
);


ALTER TABLE public."Perro_Vacunas" OWNER TO postgres;

--
-- Name: cats; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cats (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    breed character varying(255) NOT NULL,
    age integer NOT NULL,
    dueno_id integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.cats OWNER TO postgres;

--
-- Name: cats_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cats_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cats_id_seq OWNER TO postgres;

--
-- Name: cats_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cats_id_seq OWNED BY public.cats.id;


--
-- Name: duenos; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.duenos (
    id integer NOT NULL,
    nombre character varying(255) NOT NULL,
    edad integer NOT NULL,
    sexo character varying(255) NOT NULL
);


ALTER TABLE public.duenos OWNER TO postgres;

--
-- Name: duenos_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.duenos ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.duenos_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: imagenes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.imagenes (
    id integer NOT NULL,
    perro_id integer,
    url text
);


ALTER TABLE public.imagenes OWNER TO postgres;

--
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
-- Name: imagenes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.imagenes_id_seq OWNED BY public.imagenes.id;


--
-- Name: perros; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.perros (
    id integer NOT NULL,
    nombre character varying(255) NOT NULL,
    raza character varying(255) NOT NULL,
    color character varying(255) NOT NULL,
    edad smallint NOT NULL,
    dueno_id integer
);


ALTER TABLE public.perros OWNER TO postgres;

--
-- Name: perros_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.perros ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.perros_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    email character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    rol character varying(255),
    hash character varying(255),
    password character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT users_name_check CHECK ((length((name)::text) >= 2)),
    CONSTRAINT users_password_check CHECK ((length((password)::text) >= 5))
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: vacunas; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.vacunas (
    id integer NOT NULL,
    fecha date NOT NULL,
    nombre_vacuna character varying NOT NULL,
    perro_id integer
);


ALTER TABLE public.vacunas OWNER TO postgres;

--
-- Name: vacunas_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.vacunas ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.vacunas_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: cats id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cats ALTER COLUMN id SET DEFAULT nextval('public.cats_id_seq'::regclass);


--
-- Name: imagenes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.imagenes ALTER COLUMN id SET DEFAULT nextval('public.imagenes_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: Dueno_Perros; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Dueno_Perros" (id, "Dueno_id", "Perro_id") FROM stdin;
\.


--
-- Data for Name: Perro_Vacunas; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Perro_Vacunas" (id, "Perro_id", "Vacuna_id") FROM stdin;
\.


--
-- Data for Name: cats; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cats (id, name, breed, age, dueno_id, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: duenos; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.duenos (id, nombre, edad, sexo) FROM stdin;
1	Franco	30	Masculino
2	Laura	24	Femenino
3	Jaime	50	Masculino
\.


--
-- Data for Name: imagenes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.imagenes (id, perro_id, url) FROM stdin;
\.


--
-- Data for Name: perros; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.perros (id, nombre, raza, color, edad, dueno_id) FROM stdin;
9	Kiara	Poodle	Blanco	3	1
1	Buddy	Labrador	Amarillo	3	1
2	Bella	Golden Retriever	Dorado	4	1
3	Charlie	Bulldog	Blanco	2	1
4	Lucy	Poodle	Negro	5	1
5	Max	Beagle	Tricolor	1	1
10	Kiara	Poodle	Blanco	3	1
13	Kiara	Poodle	Blanco	3	1
11	Hermosa	Poodleee	Blanco	3	1
14	Kiara	Poodle	Blanco	3	1
26	Copito	Poodle	Blanco	10	3
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, email, name, rol, hash, password, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: vacunas; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.vacunas (id, fecha, nombre_vacuna, perro_id) FROM stdin;
1	2018-03-08	Vacuna 1	1
2	2005-03-02	Vacuna 2	1
3	2009-12-09	Vacuna 3	1
4	2023-08-06	Vacuna D	1
6	2023-08-05	Vacuna 10	1
12	2023-08-05	Vacuna 10	1
14	1992-11-15	Test	2
15	1992-11-15	Tesst 2	2
16	2023-08-24	Vacuna 100	1
17	1111-11-11	o	1
\.


--
-- Name: cats_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cats_id_seq', 1, true);


--
-- Name: duenos_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.duenos_id_seq', 3, true);


--
-- Name: imagenes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.imagenes_id_seq', 1, false);


--
-- Name: perros_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.perros_id_seq', 26, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, true);


--
-- Name: vacunas_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.vacunas_id_seq', 17, true);


--
-- Name: duenos Dueno_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.duenos
    ADD CONSTRAINT "Dueno_pkey" PRIMARY KEY (id);


--
-- Name: Perro_Vacunas Perro_Vacunas_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Perro_Vacunas"
    ADD CONSTRAINT "Perro_Vacunas_pkey" PRIMARY KEY (id);


--
-- Name: perros Perro_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.perros
    ADD CONSTRAINT "Perro_pkey" PRIMARY KEY (id);


--
-- Name: vacunas Vacuna_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vacunas
    ADD CONSTRAINT "Vacuna_pkey" PRIMARY KEY (id);


--
-- Name: cats cats_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cats
    ADD CONSTRAINT cats_pkey PRIMARY KEY (id);


--
-- Name: imagenes imagenes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.imagenes
    ADD CONSTRAINT imagenes_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: fk_vacunas; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fk_vacunas ON public."Perro_Vacunas" USING btree ("Vacuna_id");


--
-- Name: cats cats_dueno_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cats
    ADD CONSTRAINT cats_dueno_id_fkey FOREIGN KEY (dueno_id) REFERENCES public.duenos(id);


--
-- Name: perros dueno_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.perros
    ADD CONSTRAINT dueno_id_fk FOREIGN KEY (dueno_id) REFERENCES public.duenos(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: Dueno_Perros dueno_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Dueno_Perros"
    ADD CONSTRAINT dueno_id_foreign FOREIGN KEY ("Dueno_id") REFERENCES public.duenos(id) NOT VALID;


--
-- Name: imagenes imagenes_perro_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.imagenes
    ADD CONSTRAINT imagenes_perro_id_fkey FOREIGN KEY (perro_id) REFERENCES public.perros(id);


--
-- Name: vacunas perro_id_FK; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vacunas
    ADD CONSTRAINT "perro_id_FK" FOREIGN KEY (perro_id) REFERENCES public.perros(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: Perro_Vacunas perro_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Perro_Vacunas"
    ADD CONSTRAINT perro_id_foreign FOREIGN KEY ("Perro_id") REFERENCES public.perros(id);


--
-- Name: Dueno_Perros perro_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Dueno_Perros"
    ADD CONSTRAINT perro_id_foreign FOREIGN KEY ("Perro_id") REFERENCES public.perros(id) NOT VALID;


--
-- Name: Perro_Vacunas vacuna_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Perro_Vacunas"
    ADD CONSTRAINT vacuna_id_foreign FOREIGN KEY ("Vacuna_id") REFERENCES public.vacunas(id) NOT VALID;


--
-- PostgreSQL database dump complete
--

