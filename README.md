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

## Instrucciones para desplegar proyecto

- Desde el directorio de trabajo, clonar con GIT los siguientes repositorios, usando los siguientes comandos:
```bash
git clone --branch asantana https://github.com/citiaps/capacitaciones-deployment.git
```
- Posicionarse en el siguiente directorio:
```bash
cd capacitaciones-deployment/Docker/
```
- Renombrar el archivo .env.example a .env y agregar en la variable de entorno GITHUB_TOKEN,  el Personal Access Token de Github correspondiente al usuario con privilegios para clonar los respositorios citiaps/capacitaciones-back-go y citiaps/capacitaciones-front-nuxt. Para ver instrucciones de como generar un Personal Access Token en Github, haga click [aquí](https://docs.github.com/es/enterprise-server@3.7/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens)

- Ejectar el siguiente comando docker:
```bash
docker compose up
```
- Una vez que haya levantado todos los sistemas, abrir un navegador y colocar en la barra de direcciones:
[http://localhost:3000](http://localhost:3000)

## Mis Datos Personales
* Nombre: Ariel Santana
* Email: ariel.santana@umag.cl
