# Wiselink Test Golang y React

por Franco Montini

## CONFIG BACKEND

### Docker, Postgress y Golang(Gorm, GorillaMux)

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

## CONFIG FRONTEND

### Docker, React y Redux-Toolkit
