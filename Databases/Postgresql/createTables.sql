CREATE TABLE IF NOT EXISTS "Perro" (
                                          "id" bigint PRIMARY KEY,
                                          "Nombre" varchar NOT NULL,
                                          "Raza" varchar NOT NULL,
                                          "Edad" bigint NOT NULL,
                                          "idDueño" bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS "Dueño" (
                                       "id" bigint PRIMARY KEY,
                                       "Nombre" varchar NOT NULL,
                                       "Edad" bigint NOT NULL,
                                       "Sexo" varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS "Vacuna" (
                                               "id" bigint PRIMARY KEY,
                                               "Fecha" boolean NOT NULL,
                                               "NombreVacuna" Date NOT NULL,
                                               "idPerro" bigint NOT NULL
);


ALTER TABLE "Perro" ADD FOREIGN KEY ("idDueño") REFERENCES "Dueño" ("id");

ALTER TABLE "Vacuna" ADD FOREIGN KEY ("idPerro") REFERENCES "Perro" ("id");



