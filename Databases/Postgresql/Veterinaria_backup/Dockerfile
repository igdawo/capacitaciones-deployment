# Usar la imagen base oficial de PostgreSQL 15
FROM postgres:15

# Establecer una variable de entorno para deshabilitar las advertencias de interacción con el usuario
ENV DEBIAN_FRONTEND=noninteractive

# Copiar los archivos SQL a la imagen
COPY Backup_veterinaria_citiaps.sql /docker-entrypoint-initdb.d/

# Establecer las variables de entorno para la creación de la base de datos
ENV POSTGRES_DB veterinaria_citiaps
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD admin

# Puerto por defecto de PostgreSQL
EXPOSE 5433