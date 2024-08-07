
# Fase de construcción
FROM node:16 as build-stage
COPY frontend/package*.json ./

RUN npm cache clean --force
RUN npm install
COPY . .
RUN npm run serve



# Fase de producción
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
