# Wiselink Test Golang y React

Testing de endpoints:
Ir al la web de [Postman Endpoints](https://documenter.getpostman.com/view/22472853/VUr1GYKJ) y luego ir a probar con postman(arriba a la derecha)

## BACKEND

Tecnologias utilizadas:
Docker, Golang, GORM, gorilla/mux, Postgress, AIR y Postman(TEST y ENDPOINTS).
Host API: [http://localhost:8080/](http://localhost:8080/)

### Pasos para iniciar postgres con docker y el backend.

- Desde la consola (BASH) ir a la carpeta api-golang
  `cd api-golang/`

- Crear docker virtual server de Postgres

`docker run --name docker_db -e POSTGRES_USER=montini -e POSTGRES_PASSWORD=123123 -p 5433:5432 -d postgres`

- Controlar si esta ON

`docker ps`

- Ingresar al virtual server de Postgres

`docker exec -it docker_db bash`

- Conectarse (Nos va a pedir la pass es: 123123)

`psql -U montini --password`

- Crear la base de datos con la que vamos a trabajar

`CREATE DATABASE montini_db;`

- Entramos a la base de datos (Nos va a pedir la pass es: 123123)

`\c montini_db`

- Volvemos a la consola (BASH) ubicados de la carpeta: /api-golang e iniciamos el server

`go run main.go` o `go run .` o `air`

Test de Backend creado por Franco Montini para Wiselink.
