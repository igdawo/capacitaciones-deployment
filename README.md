# Tutor de commit
Nombre: Hernán Pinochet
Correo Electrónico: hernan.pinochet@usach.cl

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
    

# Pasos para la instalación de dockers
* Paso 1: Creación de las imagenes por medio de los dockerfiles
    * En la carpeta de Backend: docker build . -t backend_image
    * En la carpeta de Frontend: docker build . -t frontend_image
    * En la carpeta de Postgresql: docker build . -t postgresql_image
    * En la carpeta de Mongodb: docker build . -t mongodb_image
* Paso 2: Se ejecuta Docker-compose
    * En la carpeta de Docker-compose: docker-compose up -d
* Paso 3: Realizar la carga de datos correspondientes a cada base de datos (Restore):
    * Para Postgresql realizar lo siguiente en su terminal de su contendedor correspondiente:
        * Dirigirse a la carpeta docker-entrypoint-initdb.d con: cd docker-entrypoint-initdb.d
        * Realizar el restore con: pg_restore -U postgres -d veterinaria_citiaps_v2 veterinaria_citiaps_v2.sql
    * Para Mongodb realizar lo siguiente en su terminal de su contendedor correspondiente:
        * Dirigirse a la carpeta docker-entrypoint-initdb.d con: cd docker-entrypoint-initdb.d
        * Se crea la carpeta dump: mkdir dump
        * Se ingresa a la carpeta dump la carpeta veterinaria_citiaps: mv veterinaria_citiaps dump
        * Se realiza el restore: mongorestore --username=admin --password=admin --authenticationDatabase=admin dump
        * Como podrán ver no se subió la carpeta dump, sino lo que contenía, por lo que se debe crear la carpeta manualmente y ingresarle veterinaria_citiaps. Si hace directamente el mongorestore en veterinaria_citiaps por alguna razón no funcionara.

Dicho lo anterior ya se puede navegar en la pagina web, PD1: Solo funciona la pagina web con la base de datos postgresql ya que esta no esta adaptada para el tipo de dato de mongodb. PD2: En la carpeta de front revisar la rama hpinochet_v2, la rama hpinochet ignorarla.