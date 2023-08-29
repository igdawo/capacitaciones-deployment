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

# Franco Miranda
franco.miranda@umag.cl


# Pasos para desplegar docker
Construir y Levantar los Servicios: Navega hasta la carpeta donde se encuentra el archivo docker-compose.yml y ejecuta el siguiente comando para construir y levantar todos los servicios:

```
docker-compose up --build
```

Verificar los Servicios: Puedes verificar el estado de los servicios ejecutando:
```
docker-compose ps
```

Detener y Eliminar los Servicios: Cuando hayas terminado, puedes detener y eliminar todos los servicios con:
```
docker-compose down 
```