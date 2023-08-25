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
* Integrante
	* Nombre: Héctor Ballesteros
	* E-mail: hector.ballesteros@usach.cl
Requisitos:
+ Docker instalado
+ Docker-compose instalado

Para realizar despliegue, abrir la consola y dirigirse 
al directorio donde se encuentre el archivo docker-compose.yml

Ejecutar el siguiente comando:
sudo docker-compose up --build -d

Por defecto, el frontend accede al backend GO con base de datos MongoDB,
el cual se sitúa en http://localhost:8000.
Para cambiar el backend al que tiene base de datos PostgreSQL,
dirigirse al directorio Config/Frontend/.env, y modificar la 
enviroment "BACK-URL" a http://localhost:8080

*Puerto Go+MongoDB=8000
*Puerto Go+PostgreSQL=8080

Para visualizar la aplicación, ingresar http://localhost:3000 en el navegador

+Consideraciones: 
- las imágenes de cada servicio fueron buildeadas
mediante sus Dockerfile y pusheadas a Docker-hub, por lo que el 
docker-compose.yml obtiene las imágenes desde allí.
- las bases de datos (tanto postgresql como mongodb) están previamente
cargadas con datos, ya que la aplicación al inicializarse,
realiza automáticamente el restore de los dump

