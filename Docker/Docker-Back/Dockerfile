# Utiliza una imagen base de Go
FROM golang:1.21-alpine

# Establece el directorio de trabajo en la aplicación
WORKDIR /app

# Copia los archivos de la aplicación al directorio de trabajo
COPY . .

# Compila la aplicación Go
RUN go build -o backend

# Expone el puerto en el que la aplicación se ejecutará (si es necesario)
EXPOSE 8000

# Comando para ejecutar la aplicación
CMD ["./backend"]
