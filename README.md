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


Felipe Gómez Obreque
felipe.gomez.o@usach.cl

Pasos para el despliegue:

1) Abrir una consola en la carpeta /Docker
2) Escribir el comando docker-compose up -d

Con eso correra el contenedor de back, front y bd con las imagenes que se encuentran en mi repositorio de dockerhub.

* Las imagenes de las bases de datos contienen los datos del respaldo.

Luego se puede interactuar desde el navegador con la app usando el link:
localhost:3000

* Si se usa chrome se debe agregar la extensión: Moesif Orign & CORS Changer y asegurar que esté en estado ON para que backend y frontend puedan interactuar bien.