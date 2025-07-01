FROM node:20

# Establece el directorio de trabajo
WORKDIR /app

# Copia los archivos del frontend
COPY . .

# Instala dependencias
RUN npm install

# Compila la app
RUN npm run build

# Instala servidor liviano para servir contenido estático
RUN npm install -g serve

# Expone el puerto de producción
EXPOSE 5173

# Comando por defecto
CMD ["serve", "-s", "dist"]
