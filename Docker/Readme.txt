
#Consideraciones:
1) Verificar la existencia de los archivos .env en las carpetas:
	../Backend/BackendMongo
	../Backend/BackendPostgres
	./Frontend
2) En caso de no existir, cada carpeta posee un .env.example los cuales pueden ser utilizados para el despliegue.
3) Para utilizar un .env.example, modificar su nombre a .env

#Despliegue

#Despliegue para MongoDB
#El archivo Docker-compose.yaml trae por defecto la elaboración del proyecto utilizando la Base de datos MongoDB
#Para ejecutarlo, seguir los siguientes pasos:
1) Abrir una terminar dentro de la carpeta Docker la cual cuenta con todas las carpetas del proyecto.
2) En la terminal, escribir el comando docker-compose up -d
3) Al finalizar, es posible acceder al Frontend de la aplicación en http://localhost:3000/

#Despliegue para PostgreSQL:
1) Comentar todo el bloque de código para ejecutar el proyecto de MongoDB (Líneas 6 a 48)
2) Descomentar todo el bloque de codigo para ejecutar el proyecto de PostgreSQL (Líneas 50 a 93)
3) Abrir una terminar dentro de la carpeta Docker la cual cuenta con todas las carpetas del proyecto.
4) En la terminal, escribir el comando docker-compose up -d
5) Al finalizar, es posible acceder al Frontend de la aplicación en http://localhost:3000/
