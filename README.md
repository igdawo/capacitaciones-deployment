Nombre: Alan Cristian Curilem Chacón
Email: alan.curilem@umag.cl

Comandos a ejecutar:
cd Docker
docker compose up -d
zcat ../Databases/Postgresql/veterinaria_citiaps.sql.gz | docker exec -i <nombre-contenedor-postgresql> psql -U alan veterinaria_citiaps
En Config/Backend/.env.docker POSTGRES_HOST colocar el nombre del contendor de la bd postgresql


# Repositorio Capacitaciones Deployment
En este repositorio se encuentran las carpetas donde podrán subir los dump de las base de datos y los archivos de configuración para el despliegue con Docker
## Estructura de carpetas
* Config - Debe contener los archivos de configuración 
    * Backend - Archivos de configuración para el backend
    * Frontend - Archivos de configuración para el frontend
* Databases
    * Mongodb - Dump de la base de datos en mongodb
    * Postgresql - Dump de la base de datos en postgresql
* Docker - Debe contener los archivos necesarios para realizar el despliegue mediante docker
