# Información Personal
*Nombre: María Andrea Salinas Pittet  Email: andrea.salinas@umag.cl

# Estructura de carpetas
 Config - Contiene los archivos de configuración tanto del backend como el frontend
    Backend - Archivo Dockerfile y .env para cargar el backend
    Frontend - Archivo dockerfile y .env para cargar el frontend
 Databases
    Postgresql - base de datos en postgresql
 Docker - Contiene los archivos para el despliegue por docker

 # Información Personal
*Nombre: María Andrea Salinas Pittet  Email: andrea.salinas@umag.cl

# Pasos
Se debe copiar el capacitaciones-deployment  para ello se debe realizar un 
git clone -b asalinas https://github.com/citiaps/capacitaciones-deployment.git veterinaria_citiaps

-Posteriomente se debe configurar las varibles de los .env  del backend frontend y base de datos

CONFIG /BACKEND / .env se utilizan las siguientes variables
 GO_REST_ENV=dev
 GIN_MODE=debug
 DB_USER=usuario base de datos
 DB_PASS=clave base de datos
 DB_DB=nombre de base de datos
 DB_URL=db:5432
 ADDR=0.0.0.0:3001
 JWT_KEY=string_largo_unico_por_proyecto
CONFIG / FRONTEND / .env se debe colcoar la url del backurl y browserurl
 BACK_URL=0.0.0.0:3001
 HOST=0.0.0.0
 BROWSER_BASE_URL="http://localhost:3001"

DOCKER /.env debe colocar la variable
 GITHUB_TOKEN=token de github

Luego debe ir a la carpeta de docker por terminal y construir con docker-compose build --no-cache
y luego realizar docker-compose up -d 
Para acceder al sitio puede realizarlo desde el navegador en http://localhost:3000
