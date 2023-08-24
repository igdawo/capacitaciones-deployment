Nombre: Alan Cristian Curilem Chacón
Email: alan.curilem@umag.cl

Comandos a ejecutar:
cd Docker
docker compose up -d
zcat ../Databases/Postgresql/veterinaria_citiaps.sql.gz | docker exec -i <nombre-contenedor-bd-postgresql> psql -U alan veterinaria_citiaps
En Config/Backend/.env.docker POSTGRES_HOST colocar el nombre del contendor de la bd postgresql
En Config/Frontend/.env.docker backUrl se debe colocar la url del backend