# Usa la imagen oficial de Go
FROM golang:1.20

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el código al contenedor
COPY . .

# Descarga las dependencias
RUN go mod tidy

# Comando predeterminado para ejecutar la app
CMD ["go", "run", "main.go"]
