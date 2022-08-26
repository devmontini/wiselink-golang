# Wiselink Test Golang y React

API para MVP de Wiselink con Golang.

Testing de endpoints:

1. Ir al la web de [Postman Endpoints](https://documenter.getpostman.com/view/22472853/VUr1GYKJ)
2. luego ir a probar los tests con postman(Para abrir con postman client clickear arriba a la derecha donde dice "Run in Postman")

A tener en cuenta:

1. Dos de los endpoints no los pude lograr, intente usar el ORMs GROM y se me complico con el tema de agregar eventos a los usuarios (muchos a muchos), por ende la de ver las inscripciones de cada usuario no anda.
2. El filtro del tiempo no anda porque al pasarlo a string y no quedaba igual a los tiempos de los eventos.
3. Mi fuerte es Node.js poero queria intentarlo con Golang, no me salio como esperaba.

## BACKEND

Tecnologias utilizadas:
Docker, Golang, GORM, gorilla/mux, Postgress, AIR y Postman(TEST y ENDPOINTS).
Host API: [http://localhost:8080/](http://localhost:8080/)

### Pasos para iniciar postgres con docker y el backend con golang sin docker.

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
