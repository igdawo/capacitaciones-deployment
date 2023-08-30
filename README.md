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

## Instruccioner para deplegar proyecto

- Desde el directorio de trabajo, clonar con GIT los siguientes repositorios, usando los siguientes comandos:
```bash
git clone --branch asantana-postgresql https://github.com/citiaps/capacitaciones-back-go.git
git clone --branch asantana https://github.com/citiaps/capacitaciones-front-nuxt.git
git clone --branch asantana https://github.com/citiaps/capacitaciones-deployment.git
```

- Renombrar el siguiente archivo:
```bash
mv capacitaciones-back-go/.env.example capacitaciones-back-go/.env
```
- Posicionarse en el siguiente directorio:
```bash
cd capacitaciones-deployment/Docker/
```
- Ejectar el siguiente comando docker:
```bash
docker compose up
```
- Una vez que haya levantado todos los sistemas, abrir un navegador y colocar en la barra de direcciones:
[http://localhost:3000](http://localhost:3000)

## Mis Datos Personales
* Nombre: Ariel Santana
* Email: ariel.santana@umag.cl
