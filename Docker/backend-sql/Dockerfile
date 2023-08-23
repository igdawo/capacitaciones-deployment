# Utiliza una imagen base de Go para compilar
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos de código fuente de tu aplicación
COPY . .


# Compila tu aplicación Go y genera el ejecutable
RUN go build -o main .

# Usa una imagen más ligera para ejecutar tu aplicación
#FROM alpine:latest
# Copia el archivo .env a la imagen
COPY .env /.env

# Comando para ejecutar tu aplicación
CMD ["./main"]