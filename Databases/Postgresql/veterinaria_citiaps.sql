/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : PostgreSQL
 Source Server Version : 150004 (150004)
 Source Host           : localhost:5432
 Source Catalog        : veterinaria_citiaps
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 150004 (150004)
 File Encoding         : 65001

 Date: 17/08/2023 09:23:40
*/


-- ----------------------------
-- Table structure for Dueno
-- ----------------------------
DROP TABLE IF EXISTS "public"."Dueno";
CREATE TABLE "public"."Dueno" (
  "id" int4 NOT NULL,
  "Nombre" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "Edad" int4 NOT NULL,
  "Sexo" varchar(255) COLLATE "pg_catalog"."default" NOT NULL
)
;

-- ----------------------------
-- Records of Dueno
-- ----------------------------
INSERT INTO "public"."Dueno" VALUES (1, 'Franco', 30, 'Masculino');
INSERT INTO "public"."Dueno" VALUES (2, 'Laura', 24, 'Femenino');
INSERT INTO "public"."Dueno" VALUES (3, 'Jaime', 50, 'Masculino');
INSERT INTO "public"."Dueno" VALUES (4, 'Javier', 18, 'Masculino');
INSERT INTO "public"."Dueno" VALUES (5, 'Eduardo', 23, 'Femenino');

-- ----------------------------
-- Table structure for Dueno_Perros
-- ----------------------------
DROP TABLE IF EXISTS "public"."Dueno_Perros";
CREATE TABLE "public"."Dueno_Perros" (
  "id" int4 NOT NULL,
  "Dueno_id" int4 NOT NULL,
  "Perro_id" int4 NOT NULL
)
;

-- ----------------------------
-- Records of Dueno_Perros
-- ----------------------------

-- ----------------------------
-- Table structure for Perro
-- ----------------------------
DROP TABLE IF EXISTS "public"."Perro";
CREATE TABLE "public"."Perro" (
  "id" int4 NOT NULL,
  "Nombre" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "Raza" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "Color" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "Edad" int2 NOT NULL
)
;

-- ----------------------------
-- Records of Perro
-- ----------------------------
INSERT INTO "public"."Perro" VALUES (1, 'Buddy', 'Labrador', 'Amarillo', 3);
INSERT INTO "public"."Perro" VALUES (2, 'Bella', 'Golden Retriever', 'Dorado', 4);
INSERT INTO "public"."Perro" VALUES (3, 'Charlie', 'Bulldog', 'Blanco', 2);
INSERT INTO "public"."Perro" VALUES (4, 'Lucy', 'Poodle', 'Negro', 5);
INSERT INTO "public"."Perro" VALUES (5, 'Max', 'Beagle', 'Tricolor', 1);

-- ----------------------------
-- Table structure for Perro_Vacunas
-- ----------------------------
DROP TABLE IF EXISTS "public"."Perro_Vacunas";
CREATE TABLE "public"."Perro_Vacunas" (
  "id" int4 NOT NULL,
  "Perro_id" int4 NOT NULL,
  "Vacuna_id" int4 NOT NULL
)
;

-- ----------------------------
-- Records of Perro_Vacunas
-- ----------------------------

-- ----------------------------
-- Table structure for Vacuna
-- ----------------------------
DROP TABLE IF EXISTS "public"."Vacuna";
CREATE TABLE "public"."Vacuna" (
  "id" int4 NOT NULL,
  "Fecha" date NOT NULL,
  "NombreVacuna" varchar COLLATE "pg_catalog"."default" NOT NULL
)
;

-- ----------------------------
-- Records of Vacuna
-- ----------------------------
INSERT INTO "public"."Vacuna" VALUES (1, '2018-03-08', 'Vacuna 1');
INSERT INTO "public"."Vacuna" VALUES (2, '2005-03-02', 'Vacuna 2');
INSERT INTO "public"."Vacuna" VALUES (3, '2009-12-09', 'Vacuna 3');
INSERT INTO "public"."Vacuna" VALUES (4, '2022-05-24', 'Vacuna 4');
INSERT INTO "public"."Vacuna" VALUES (5, '2012-07-08', 'Vacuna 5');

-- ----------------------------
-- Primary Key structure for table Dueno
-- ----------------------------
ALTER TABLE "public"."Dueno" ADD CONSTRAINT "Dueno_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table Perro
-- ----------------------------
ALTER TABLE "public"."Perro" ADD CONSTRAINT "Perro_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table Perro_Vacunas
-- ----------------------------
CREATE INDEX "fk_vacunas" ON "public"."Perro_Vacunas" USING btree (
  "Vacuna_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table Perro_Vacunas
-- ----------------------------
ALTER TABLE "public"."Perro_Vacunas" ADD CONSTRAINT "Perro_Vacunas_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table Vacuna
-- ----------------------------
ALTER TABLE "public"."Vacuna" ADD CONSTRAINT "Vacuna_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table Dueno_Perros
-- ----------------------------
ALTER TABLE "public"."Dueno_Perros" ADD CONSTRAINT "dueno_id_foreign" FOREIGN KEY ("Dueno_id") REFERENCES "public"."Dueno" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."Dueno_Perros" ADD CONSTRAINT "perro_id_foreign" FOREIGN KEY ("Perro_id") REFERENCES "public"."Perro" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table Perro_Vacunas
-- ----------------------------
ALTER TABLE "public"."Perro_Vacunas" ADD CONSTRAINT "perro_id_foreign" FOREIGN KEY ("Perro_id") REFERENCES "public"."Perro" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."Perro_Vacunas" ADD CONSTRAINT "vacuna_id_foreign" FOREIGN KEY ("Vacuna_id") REFERENCES "public"."Vacuna" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
