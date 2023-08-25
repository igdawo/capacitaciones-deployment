# Usa una imagen base de Go para construir la aplicación
FROM golang:alpine

# Se crea directorio de destino
WORKDIR /app

# Se copian las dependencias de go.mod
COPY go.mod .
RUN go mod download

# Se copia la aplicación
COPY . .

# Se compila la aplicación
RUN go build -o main.go


# Exponemos el puerto en el que escucha la aplicación
EXPOSE 8000

# Se ejecuta la aplicación cuando el contenedor se inicie
CMD ["./main"]