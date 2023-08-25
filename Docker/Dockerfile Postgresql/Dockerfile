# Utiliza la imagen base de PostgreSQL
FROM postgres:latest

# Variables de entorno opcionales para configurar PostgreSQL
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD admin
ENV POSTGRES_DB veterinaria_citiaps_v2

# Copia el archivo de volcado SQL a la imagen
COPY veterinaria_citiaps_v2.sql /docker-entrypoint-initdb.d/

# Puerto en el que PostgreSQL escuchará dentro del contenedor
EXPOSE 5432