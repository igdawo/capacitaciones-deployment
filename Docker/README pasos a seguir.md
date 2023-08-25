# Aplicación veterinaria Citiaps
Repositorio contiene el backend, frontend y bases de datos para montar en contenedores docker de cada componenete.

El repositorio contiene:
 * Archivos para el Fronend 
 * Archivos para el Backend conectado a MongoDB
 * Archivos para el Backend conectado a Postgresql
 * Archivos para realizar backup de la base de datos de MongoDB
 * Archivos para realizar backup de la base de datos de Postgresql
 * Dockerfile para cada uno de los compenentes anteriormente mencionados
 * Para la BD de mongo se cuenta con un docker-compose en vez de un dockefile
 * Archivo Docker-compose.yaml para montar los contenedores

## Ejecución

Antes de ejecutar la aplicación, es necesario:
1) Para correr el proyecto, clonar en cualquier directorio.
2) Verificar la existencia de los archivos .env, en caso de no existir, dDefinir los parámetros en el archivo .env: Copiar [.env.example](.env.example), renombrar a .env y cambiar los valores (Los valores predeterminados que se contienen deben funcionar siempre que dichos puertos esten desocupados).
3) Se debe modificar el Docker-compose.yaml: Por defecto se levantan los contederores con el Backend y base de datos de MongoDB. Para correr el Backend y base de datos de Postgresql se deben comentar los servicios de postgres y backend en el Docker-compose.yaml, y descomentar los servicios de backend y mongo (se indican mas claramente con comentarios estas secciones en el codigo). Guardar los cambios.
) Utilizar el comando:

```bash
docker-compose up -d
```
Se montan los contenedores y quedan corriendo en segundo plano

## Requisitos
* Golang v1.13 o superior.
* MongoDB 3.6 o superior.
* Postgresql 15 o superior.
* Node 18 o superior.

El proyecto fue creado con:
* Golang v1.21.0
* MongoDB 3.6.23
* Postgresql 15.3 
* Node 18.16.0

## Notas
La carpeta ya cuenta con los componenetes del proyecto, estos son iguales a los subidos a sus respectivos repositorios, la unica diferencia son los .env que han sido modificados para adaptarse a las necesidades de docker. Ademas se suma la existencia de mongorestore.sh, que permite cargar la BD de mongo.

## Estructura del proyecto

El proyecto está estructurado de la siguiente forma:

```bash
├── README.md
├── Docker-compose.yaml
├── Frontend
│   ├── Dockerfile
│   └── *archivos varios del Frontend*
├── Backend
│   ├── Back-mongodb
│   │   ├── Dockerfile
│   │   └── *archivos varios del Backend*
│   └── Back-psql
│      ├── Dockerfile
│      └── *archivos varios del Backend*
└── Databases
    ├── Mongodb
    │   ├── Docker-compose.yaml
    │   └── veterinaria_citiaps
    │       ├── mongorestore.sh
    │       └── *archivos varios de para restaurar la BD*   
    └── Postgresql
      └── veterinaria_citiaps
            ├── Dockerfile
            └── Backup_veterinaria_citiaps.sql 
```

