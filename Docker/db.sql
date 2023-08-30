-- Creación de la base de datos
CREATE DATABASE veterinaria_citiaps;

-- Conexión a la base de datos
\c veterinaria_citiaps;

CREATE TABLE duenos (
	id serial PRIMARY KEY,
	nombre VARCHAR ( 50 ) NOT NULL,	
	edad INT  NOT NULL,
	sexo VARCHAR ( 20) NOT NULL
);

CREATE TABLE perros (
	id serial PRIMARY KEY,
	nombre VARCHAR ( 50 ) NOT NULL,
	raza VARCHAR ( 50 ) NOT NULL,
	color VARCHAR ( 50 ) NOT NULL,	
	edad INT NOT NULL,
	dueno_id INT NOT NULL,
	FOREIGN KEY (dueno_id) REFERENCES duenos (id) ON DELETE CASCADE
);

CREATE TABLE vacunas (
	id serial PRIMARY KEY,
	nombre VARCHAR ( 50 ) NOT NULL,
	fecha DATE NOT NULL,	
	perro_id INT NOT NULL,
	FOREIGN KEY (perro_id) REFERENCES perros (id) ON DELETE CASCADE
);

CREATE TABLE imagenes (
	id serial PRIMARY KEY,
	url VARCHAR ( 150 ) NOT NULL,	
	perro_id INT NOT NULL,
	FOREIGN KEY (perro_id) REFERENCES perros (id) ON DELETE CASCADE
);

insert into duenos (nombre, edad, sexo) values ('Shaggy', 20,'Masculino');
insert into duenos (nombre, edad, sexo) values ('Tomás Pérez', 31,'Masculino');
insert into duenos (nombre, edad, sexo) values ('Pierre Nodoyuna', 41,'Masculino');
insert into duenos (nombre, edad, sexo) values ('Bart Simpson', 10,'Masculino');
insert into duenos (nombre, edad, sexo) values ('Juana Gomez', 29,'Femenino');

insert into perros (nombre, raza, color, edad, dueno_id) values ('Scooby Doo','Dogo Alemán','Café', 4, 1);
insert into perros (nombre, raza, color, edad, dueno_id) values ('Spike','Bulldog','Gris', 2, 2);
insert into perros (nombre, raza, color, edad, dueno_id) values ('Patán','Quiltro','Negro', 8, 3);
insert into perros (nombre, raza, color, edad, dueno_id) values ('Ayudante','Galgo','Blanco', 3, 4);
insert into perros (nombre, raza, color, edad, dueno_id) values ('Cachupín','Quiltro','Gris', 1, 5);
insert into perros (nombre, raza, color, edad, dueno_id) values ('Bobby','Quiltro','Gris y blanco', 1, 5);

insert into vacunas (nombre, fecha, perro_id) values ('Parvovirus', '2020-01-01',1);
insert into vacunas (nombre, fecha, perro_id) values ('Polivalente', '2020-12-18',1);
insert into vacunas (nombre, fecha, perro_id) values ('Polivalente', '2022-09-11',2);
insert into vacunas (nombre, fecha, perro_id) values ('Rabia', '2016-03-15',3);
insert into vacunas (nombre, fecha, perro_id) values ('Polivalente', '2019-03-20',3);
insert into vacunas (nombre, fecha, perro_id) values ('Polivalente', '2023-08-14',3);
insert into vacunas (nombre, fecha, perro_id) values ('Rabia', '2018-05-25',4);
insert into vacunas (nombre, fecha, perro_id) values ('Parvovirus', '2023-08-05',5);
insert into vacunas (nombre, fecha, perro_id) values ('Parvovirus', '2023-08-05',6);