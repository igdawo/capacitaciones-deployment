# capacitaciones-front-nuxt
Tarea capacitación go, nuxt, docker

## Información personal Alumno
Nombre: Alan Cristian Curilem Chacón
Email: alan.curilem@umag.cl

## Estructura de carpetas
* Config - Contiene los archivos de configuración 
    * Backend - Archivos de configuración para el backend
    * Frontend - Archivos de configuración para el frontend
* Databases
    * Mongodb - Dump de la base de datos en mongodb
    * Postgresql - Dump de la base de datos en postgresql
* Docker - Contiene los archivos necesarios para realizar el despliegue mediante docker

## Despliegue
```bash
- Se debe clonar el repositorio con los archivos necesarios para el despliegue
git clone -b acurilem https://github.com/citiaps/capacitaciones-deployment.git veterinaria_citiaps

- Luego se debe configurar las variables de entorno del backend, frontend y .env del directorio raiz.
En Config/Backend/.env.docker se deben setear las siguiente variables de entorno
	POSTGRES_HOST: nombre del servidor de base de datos postgresql que se creará.
	POSTGRES_USER: nombre de usuario de la base de datos postgresql.
	POSTGRES_PASSWORD: contraseña de la base de datos postgresql.
	POSTGRES_DB: nombre de la base de datos postgresql a crear.
	POSTGRES_PORT: puerto del servidor de base de datos postgresql.
	DB_USER: nombre de usuario base de datos mongodb.
	DB_PASS: contraseña del DB_USER de la base de datos mongodb.
	DB_DB: nombre de la base de datos mongodb.
	DB_URL: direccion ip de base de datos mongodb
	JWT_KEY: token unico para el consumo seguro de apis desde el backend.
	CORS_URLS: dominios permitidos para el consumo de apis, debe agregarse el dominio del frontend

En Config/Frontend/.env.docker se deben setear las siguiente variables de entorno
	BACK_URL se debe colocar la url del backend (por ejemplo http://localhost:3001/)

En .env
	BRANCH_BACK: nombre de la rama del repositorio a clonar desde https://github.com/citiaps/capacitaciones-back-go.git
	BRANCH_FRONT: nombre de la rama del repositorio a clonar desde https://github.com/citiaps/capacitaciones-front-nuxt.git
	GITHUB_TOKEN: token con permisos necesarios para clonar los repositorios desde https://github.com/citiaps/capacitaciones-back-go.git y https://github.com/citiaps/capacitaciones-front-nuxt.git

- Creación e inicio de los contenedores bd postgresql, backend golang y frontend nuxt.
docker compose -f veterinaria_citiaps/Docker/docker-compose.yml up -d

## Acceso al sitio

Para acceder al sitio desde el navegador debe ingresar la url http://localhost:3000
```
