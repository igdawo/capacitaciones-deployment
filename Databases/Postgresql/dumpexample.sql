DROP TABLE IF EXISTS Vacuna;
DROP TABLE IF EXISTS Perro;
DROP TABLE IF EXISTS Dueño;

CREATE TABLE Dueño(
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	Nombre VARCHAR(250) NOT NULL,
	Edad INT NOT NULL,
	Sexo VARCHAR(250) NOT NULL
);

CREATE TABLE Perro(
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	Nombre VARCHAR(250) NOT NULL,
	Raza VARCHAR(250) NOT NULL,
	Color VARCHAR(250) NOT NULL,
	Edad INT NOT NULL,
	id_dueño INT NOT NULL references Dueño(id)
);

CREATE TABLE Vacuna(
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	fecha DATE NOT NULL,
	NombreVacuna varchar(250) NOT NULL,
	id_perro INT NOT NULL references Perro(id)
);

insert into dueño(nombre, edad, sexo) values('Pedro', 25, 'Masculino');
insert into dueño(nombre, edad, sexo) values('Juan', 30, 'Masculino');
insert into dueño(nombre, edad, sexo) values('Daniella', 19, 'Femenino');
insert into dueño(nombre, edad, sexo) values('Sofia', 23, 'Femenino');
insert into dueño(nombre, edad, sexo) values('Jesus', 32, 'Otro');

insert into perro (nombre, raza, color, edad, id_dueño) values ('Firulais', 'Quiltro', 'Cafe claro', 9, 1);
insert into perro (nombre, raza, color, edad, id_dueño) values ('Mopa', 'Poodle', 'Blanco', 7, 2);
insert into perro (nombre, raza, color, edad, id_dueño) values ('Muñeca', 'Quiltro', 'Blanco', 12, 3);
insert into perro (nombre, raza, color, edad, id_dueño) values ('Micky', 'Yorkshaire Terrier', 'Cafe oscuro', 5, 4);
insert into perro (nombre, raza, color, edad, id_dueño) values ('Javi', 'Pastor aleman', 'Cafe claro', 2, 5);

insert into vacuna (fecha, nombrevacuna, id_perro) values ('2021-02-10','Octuple',1);
insert into vacuna (fecha, nombrevacuna, id_perro) values ('2012-01-09','Rabica',2);
insert into vacuna (fecha, nombrevacuna, id_perro) values ('2005-07-18','Sextuple',3);
insert into vacuna (fecha, nombrevacuna, id_perro) values ('2023-11-18','Influenza',4);
insert into vacuna (fecha, nombrevacuna, id_perro) values ('2012-12-12','Octuple',5);