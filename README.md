# Gobrax Challenge API

Este repositório está relacionado a solução de um desafio idealizado pela empresa Gobrax.

O desafio consiste no desenvolvimento das seguintes funcionalidades:

- CRUD de motoristas (Criação, Listagem, Atualização e Remoção)

- CRUD de veículos (Criação, Listagem, Atualização e Remoção)

- Vinculação de motorista a um veículo



## Glossário

1. Domain: Entidades, Modelos e Casos de uso
2. Infraestrutura: Builds, Conexão com banco de dados, Handler e Repositório
3. Interfaces: Todas as interfaces do projeto
4. Utils: Pacotes utilitários que podem ser utilizados em várias partes do projeto

## Arquitetura

1. Segue os principios da Arquitetura Limpa, com o uso extensivo de interfaces

2. Contém padrões de design, junto disso trazendo um código limpo

## Requisitos para instalar a aplicação localmente

**Golang**
```
$ Download: https://go.dev/learn/
```

**Docker**
```
$ Download: https://www.docker.com/get-started/
```

**PostgresSQL**
```
$ Download: https://www.postgresql.org
```


## Inicio

Abra o arquivo `.env`, utilize-o como referência para configurar as suas variaveis de ambiente.

Caso opte por não utilizar o docker, será necessário fazer a instalação do postgresSQL localmente e rodar o script `docker_db_init` para fazer a criação das tabelas.

O script pode ser rodado através de um SGDB como PGADMIN,DBVEAVER ou pelo client PSQL

Toda a parte executavel está em  `cmd/main.go`.

```
$ go run cmd/main.go
```

O servidor deve iniciar no seguinte endereço: [`http:localhost:8080`](http:localhost:8080)

## Utilizando Docker-Compose

Para iniciar a aplicação com o docker-compose, utilize o comando abaixo. O docker deve iniciar um build da aplicação e inicializar um container com uma instância do PostgreSQL.

```
$ docker-compose up --build
```

## Endpoints Disponiveis

### `POST /trucker/create`

- Cria um novo caminhoneiro.

```
curl --location --request POST 'http://localhost:8080/trucker/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Murilo",
    "age": 23,
    "nationatily": "Brazil"
}'
```

### `GET /trucker`

- Lista todos os caminhoneiros.

```
curl --location --request GET 'http://localhost:8080/trucker'
```


### `GET /trucker/:id`

- Lista um caminhoneiro pelo ID


```
curl --location --request GET 'http://localhost:8080/trucker/1'
```

### `PUT /trucker/:id`

- Edita um caminhoneiro pelo ID


```
curl --location --request PUT 'http://localhost:8080/trucker/1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Joao Carlos",
    "age": 70,
    "nationatily": "Argentina"
}'
```

### `DELETE /trucker/:id`

- Deleta um caminhoneiro pelo ID


```
curl --location --request DELETE 'http://localhost:8080/trucker/1'
```

### `POST /vehicle/create`

- Cria um novo veículo.

```
curl --location --request POST 'http://localhost:8080/vehicle/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Torpedo",
    "brand": "Mercedes",
    "color": "vermelho",
    "year": 2015,
    "plate": "2S4G98B"
}'
```

### `GET /vehicle`

- Lista todos os veículos.

```
curl --location --request GET 'http://localhost:8080/vehicle' \
--data-raw ''
```

### `GET /vehicle/:id`

- Lista um veículo pelo ID.

```
curl --location --request GET 'http://localhost:8080/vehicle/1' \
--data-raw ''
```

### `PUT /vehicle/:id`

- Edita um veículo pelo ID.

```
curl --location --request PUT 'http://localhost:8080/vehicle/1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Torpedo",
    "brand": "Mercedes",
    "color": "rosa",
    "year": 2016,
    "plate": "2S4G98B"
}'
```

### `DELETE /vehicle/:id`

- Deleta um veículo pelo ID.

```
curl --location --request DELETE 'http://localhost:8080/vehicle/1' \
--data-raw ''
```

### `POST /assign/vehicle/:id/trucker/:id`

- Cria uma vinculação entre um motorista e um veículo pelo ID de ambos.

```
curl --location --request POST 'http://localhost:8080/assign/vehicle/3/trucker/2' \
--data-raw ''
```
