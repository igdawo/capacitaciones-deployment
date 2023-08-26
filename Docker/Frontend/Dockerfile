# Usa una imagen base de node para construir la aplicación
FROM node:16
# Se crea directorio de destino
WORKDIR /app
COPY package*.json ./
# Instala las dependencias
RUN npm install
# Exponemos el puerto en el que escucha la aplicación
EXPOSE 3000
# Se copia la aplicación
COPY . .
# Se ejecuta la aplicación cuando el contenedor se inicie
CMD ["npm","run","start"]