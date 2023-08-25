## MER
La aplicacion fue construida considerando las especificaciones de la actividad de BD, siguiendo el mer entregado.
El .env esta ligeramente cambiado para tener el puerto y el host por separado, se incluye en el .env.example.
## Ejecución

Antes de ejecutar la aplicación, es necesario:
1) Definir los parámetros en el archivo .env: Copiar [.env.example](.env.example), renombrar a .env y cambiar los valores.

Para correr el proyecto, clonar en cualquier directorio y utilizar el comando:

```bash
go run main.go
```
La descarga de dependencias se realiza de forma automática.


## Postman
Dentro de la capeta "postman" se encuentra una coleccion con las request para todos los CRUD y las otras funcionalidades.


## Requisitos
* Golang v1.13 o superior.
* MongoDB 3.6 o superior.
