--
-- PostgreSQL database dump
--

-- Dumped from database version 15.3
-- Dumped by pg_dump version 15.3

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
-- Name: Dueño; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Dueño" (
    id integer NOT NULL,
    "Nombre" text,
    "Edad" integer,
    "Sexo" text
);


ALTER TABLE public."Dueño" OWNER TO postgres;

--
-- Name: Perro; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Perro" (
    id integer NOT NULL,
    "Nombre" text,
    "Raza" text,
    "Color" text,
    "Edad" integer,
    "Dueño_id" integer
);


ALTER TABLE public."Perro" OWNER TO postgres;

--
-- Name: Vacuna; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Vacuna" (
    id integer NOT NULL,
    "Fecha" date,
    "NombreVacuna" text,
    "Perro_id" integer
);


ALTER TABLE public."Vacuna" OWNER TO postgres;

--
-- Data for Name: Dueño; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Dueño" (id, "Nombre", "Edad", "Sexo") FROM stdin;
1	Jaime Jaramillo	59	Masculino
2	Jenna Jefferson	59	Femenino
3	Jina Jiménez	47	Femenino
4	Johana Jones	65	Femenino
5	Juismael Juárez	86	Masculino
\.


--
-- Data for Name: Perro; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Perro" (id, "Nombre", "Raza", "Color", "Edad", "Dueño_id") FROM stdin;
1	Scud	Bull terrier	Blanco con pocas manchas grises, Mancha grande en el ojo izquierdo	3	5
2	Peludito	Golden	Dorado	1	3
3	Pelao	Xoloitzcuintle	Sin pelo, Cafe muy oscuro	5	1
4	Huesillo	Terrier chileno	Blanco con la cabeza negra con detalles cafes y blancos	4	2
5	Chopin	San bernardo	Blanco con cafe	7	4
\.


--
-- Data for Name: Vacuna; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Vacuna" (id, "Fecha", "NombreVacuna", "Perro_id") FROM stdin;
1	2023-04-04	antirrábica	4
2	2023-03-03	antirrábica	3
3	2023-02-02	antirrábica	2
4	2023-05-05	antirrábica	5
5	2023-10-10	antirrábica	1
\.


--
-- Name: Dueño Dueño_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Dueño"
    ADD CONSTRAINT "Dueño_pkey" PRIMARY KEY (id);


--
-- Name: Perro Perro_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Perro"
    ADD CONSTRAINT "Perro_pkey" PRIMARY KEY (id);


--
-- Name: Vacuna Vacuna_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Vacuna"
    ADD CONSTRAINT "Vacuna_pkey" PRIMARY KEY (id);


--
-- Name: Perro fk_Perro_Dueño; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Perro"
    ADD CONSTRAINT "fk_Perro_Dueño" FOREIGN KEY ("Dueño_id") REFERENCES public."Dueño"(id) ON DELETE CASCADE NOT VALID;


--
-- Name: Vacuna pk_Vacuna_Perro; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Vacuna"
    ADD CONSTRAINT "pk_Vacuna_Perro" FOREIGN KEY ("Perro_id") REFERENCES public."Perro"(id) ON DELETE CASCADE NOT VALID;


--
-- PostgreSQL database dump complete
--

