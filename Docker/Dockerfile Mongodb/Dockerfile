# Usar la imagen base de MongoDB
FROM mongo:latest

# Variables de entorno para configurar MongoDB
ENV MONGO_INITDB_ROOT_USERNAME=admin
ENV MONGO_INITDB_ROOT_PASSWORD=admin
ENV MONGO_INITDB_DATABASE=veterinaria_citiaps

# Copia el archivo de volcado MongoDB a la imagen
COPY dump /docker-entrypoint-initdb.d/

# Exponer el puerto por defecto de MongoDB
EXPOSE 27017

# Ejecutar un script personalizado si es necesario al iniciar el contenedor
CMD ["mongod"]