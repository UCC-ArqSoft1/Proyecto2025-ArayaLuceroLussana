# Usa una imagen de Go oficial
FROM golang:1.24.1

# Establece el directorio de trabajo
WORKDIR /app

# Copia los archivos del backend
COPY . .

# Descarga las dependencias
RUN go mod tidy

# Compila la aplicación
RUN go build -o main main.go

# Expone el puerto usado por el backend
EXPOSE 8080

# Comando por defecto
CMD ["./main"]
