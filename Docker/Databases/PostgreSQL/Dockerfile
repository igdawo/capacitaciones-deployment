FROM postgres:15 
ENV DEBIAN_FRONTEND=noninteractive
COPY BackupPostgres.sql /docker-entrypoint-initdb.d/
ENV POSTGRES_DB veterinaria_citiaps
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD admin
EXPOSE 5433