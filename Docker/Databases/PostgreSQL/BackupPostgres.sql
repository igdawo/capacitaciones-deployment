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
    "idDueño" integer
);


ALTER TABLE public."Perro" OWNER TO postgres;

--
-- Name: Vacuna; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Vacuna" (
    id integer NOT NULL,
    "Fecha" date,
    "NombreVacuna" text,
    "idPerro" integer
);


ALTER TABLE public."Vacuna" OWNER TO postgres;

--
-- Data for Name: Dueño; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Dueño" (id, "Nombre", "Edad", "Sexo") FROM stdin;
1	Ariel Pérez	16	Femenino
2	Mulan Rodríguez	18	Femenino
3	Eugene Fitzherbert Torres	25	Masculino
4	Elsa Lopéz	21	Femenino
5	Eric González	21	Masculino
\.


--
-- Data for Name: Perro; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Perro" (id, "Nombre", "Raza", "Color", "Edad", "idDueño") FROM stdin;
1	Flounder	Chihuahua	Amarillo con manchas azules	2	1
2	Mushu	Beagle	Rojo con manchas amarillas	15	2
3	Max	Husky	Blanco	7	3
4	Bruni	Pomerania	Celeste con manchas moradas	1	4
5	Sebastían	Corgi	Rojo	7	5
\.


--
-- Data for Name: Vacuna; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Vacuna" (id, "Fecha", "NombreVacuna", "idPerro") FROM stdin;
1	2021-04-15	Astradogseneca	1
2	2022-05-15	Sinodogvac	2
3	2023-02-01	Dogfizer	3
4	2021-02-02	Dogfizer	4
5	2022-04-18	Astradogseneca	5
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
-- Name: Perro FK_Dueño; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Perro"
    ADD CONSTRAINT "FK_Dueño" FOREIGN KEY ("idDueño") REFERENCES public."Dueño"(id) NOT VALID;


--
-- Name: Vacuna FK_Perro; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Vacuna"
    ADD CONSTRAINT "FK_Perro" FOREIGN KEY ("idPerro") REFERENCES public."Perro"(id) NOT VALID;


--
-- PostgreSQL database dump complete
--

