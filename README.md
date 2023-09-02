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
- Renombrar el archivo .env.example a .env y agregar en la variable de entorno GITHUB_TOKEN,  el Personal Access Token de Github correspondiente al usuario con privilegios para clonar los respositorios citiaps/capacitaciones-back-go y citiaps/capacitaciones-front-nuxt. 
<sub><sup>Para obtener un token válido deberá seguir los siguientes pasos:
1.- Hace login en Github con su cuenta con privilegios para acceder a los respositorios antes mencionados. 
2.- hacer click en el icono ubicado en la esquina superior derecha.
3.- Seleccionar la opción Settings.
4.- En la columna izquierda ir a la última opción, Developer settings.
5.- En la columna izquierda, hacer click en Personal Access Token y seleccionar Tokens (classic).
6.- Seleccionar la opción Generate new token (classic).
7.- Ingresar un nombre, seleccionar su duración, y chekear la primera opción repo.
8.- Ir al final de la página, y presionar el botón Generate token.
9.- Acá se creará el token de su usuario, el cual deberá copiar y pegar en el .env</sup></sub>

- Ejectar el siguiente comando docker:
```bash
docker compose up
```
- Una vez que haya levantado todos los sistemas, abrir un navegador y colocar en la barra de direcciones:
[http://localhost:3000](http://localhost:3000)

## Mis Datos Personales
* Nombre: Ariel Santana
* Email: ariel.santana@umag.cl
