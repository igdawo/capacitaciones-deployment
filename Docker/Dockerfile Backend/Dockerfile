# Usamos una imagen base de Go para construir la aplicación
FROM golang:latest AS builder

# Actualizamos el linux del contenedor
RUN apt-get update

# Se ajustan variables de entorno
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Establecemos el directorio de trabajo dentro del contenedor
WORKDIR /app

# Se copian las dependencias de go.mod
COPY go.mod .
RUN go mod download

# Copia el archivo .env en la imagen
COPY .env .

# Copiamos los archivos de la aplicación al contenedor
COPY . .

# Compilamos la aplicación
RUN go build main.go

# Creamos una imagen más pequeña para la aplicación final
#FROM scratch

# Copiamos el binario compilado desde la etapa anterior
#COPY --from=builder /app .

# Exponemos el puerto en el que escucha la aplicación
EXPOSE 8000

# Comando para ejecutar la aplicación cuando el contenedor se inicie
CMD ["./main"]