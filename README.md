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


# Intrucciones despliegue con docker
* Acceder al proyecto de backend de la rama bdiaz-postgres
    * Crear el archivo .env, cambiando el nombre de 'ex.env' por '.env'
    * Crear imagen del backend usando el dockerfile con el comando
    ```
    docker build -t backend_image .
    ```
* Acceder al proyecto de frontend
    * Crear el archivo .env, cambiando el nombre de 'ex.env' por '.env'
    * Crear imagen del frontend usando el dockerfile con el comando
    ```
    docker build -t frontend_image .
    ```
* Acceder a la carpeta Databases/Postgresql
    * Crear imagen de la base de datos usando el dockerfile con el comando
    ```
    docker build -t postgres_image .
    ``` 
* Acceder a la carpeta Databases/Mongodb
    * Crear imagen de la base de datos usando el dockerfile con el comando
    ```
    docker build -t mongodb_image .
    ``` 
* Acceder a carpeta Docker
    * Correr docker-compose usando
    ```
    docker-compose up
    ``` 


Integrante:
Nombre: Bastián Díaz
Correo Electrónico: bastian.diaz.p@usach.cl