CREATE TABLE public.duenos (
    id serial PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL,
    edad INT NOT NULL,
    sexo VARCHAR(255) NOT NULL
);

CREATE TABLE public.perros (
    id serial PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL,
    raza VARCHAR(255) NOT NULL,
    color VARCHAR(255) NOT NULL,
    edad INT NOT NULL,
    dueno_id INT NOT NULL,
	FOREIGN KEY (dueno_id) REFERENCES duenos (id) ON DELETE CASCADE
);

CREATE TABLE public.vacunas (
    id serial PRIMARY KEY,
    fecha date,
    nombre_vacuna VARCHAR(255),
    perro_id INT NOT NULL,
	FOREIGN KEY (perro_id) REFERENCES perros (id) ON DELETE CASCADE
);

CREATE TABLE imagenes (
	id serial PRIMARY KEY,
	url VARCHAR (255) NOT NULL,	
	perro_id INT NOT NULL,
	FOREIGN KEY (perro_id) REFERENCES perros (id) ON DELETE CASCADE
);

INSERT INTO public.duenos (nombre, edad, sexo) VALUES('Pedro',	21,	'M');
INSERT INTO public.duenos (nombre, edad, sexo) VALUES('Camila',	26,	'M');
INSERT INTO public.duenos (nombre, edad, sexo) VALUES('Juan',	31,	'M');
INSERT INTO public.duenos (nombre, edad, sexo) VALUES('Esteban',	53,	'M');
INSERT INTO public.duenos (nombre, edad, sexo) VALUES('Paola',	66,	'F');


INSERT INTO public.perros (nombre, raza, color, edad, dueno_id) VALUES('Snoopy',	'Terranova',	'Cafe',	2,	1);
INSERT INTO public.perros (nombre, raza, color, edad, dueno_id) VALUES('Pelos',	'Labrador',	'Negro-Cafe',	4,	2);
INSERT INTO public.perros (nombre, raza, color, edad, dueno_id) VALUES('Feliz',	'Boxer',	'Negro',	5,	3);
INSERT INTO public.perros (nombre, raza, color, edad, dueno_id) VALUES('Max',	'Schnauzer',	'Gris',	1,	4);
INSERT INTO public.perros (nombre, raza, color, edad, dueno_id) VALUES('Snop',	'Basset',	'Cafe-Blanco',	2,	5);


INSERT INTO public.vacunas (fecha, nombre_vacuna, perro_id) VALUES('2023-03-22',	'Sextuple',	1);
INSERT INTO public.vacunas (fecha, nombre_vacuna, perro_id) VALUES('2019-08-01',	'Sextuple',	2);
INSERT INTO public.vacunas (fecha, nombre_vacuna, perro_id) VALUES('2020-12-30',	'Antirrabica',	3);
INSERT INTO public.vacunas (fecha, nombre_vacuna, perro_id) VALUES('2022-03-22',	'Octuple2019',	4);
INSERT INTO public.vacunas (fecha, nombre_vacuna, perro_id) VALUES('2023-12-16',	'Quintuple',	5);

INSERT INTO public.imagenes (url, perro_id) VALUES('/imagenes/imagen_no_disponible.png',	1);
INSERT INTO public.imagenes (url, perro_id) VALUES('/imagenes/imagen_no_disponible.png',	2);
