FROM postgres:15 
ENV DEBIAN_FRONTEND=noninteractive
COPY Backup_veterinaria_citiaps.sql /docker-entrypoint-initdb.d/
ENV POSTGRES_DB veterinaria_citiaps
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD admin
EXPOSE 5433