## Clean Code

Agora é a hora de botar a mão na massa. Para este desafio, você precisará criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:

- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL

Não esqueça de criar as migrações necessárias e o arquivo `api.http` com as requests para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (`Dockerfile` / `docker-compose.yaml`), com isso ao rodar o comando `docker compose up` tudo deverá subir, preparando o banco de dados.

Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.

---

## Steps

### 1. Subir o ambiente com Docker

```bash
docker compose up
```

### 2. Rodar a aplicação

```bash
go run main.go wire_gen.go
```

### 3. Testar os serviços

#### REST API

- Endpoint: `GET /order`
- Arquivo de exemplo: `api/create_order.http`

#### gRPC

- Use o Evans para testar:

```bash
evans -r repl
call CreateOrder
```

#### GraphQL

- Abra no navegador:
  [http://localhost:8080/](http://localhost:8080/)

#### RabbitMQ

- Abra no navegador:
  [https://localhost:15672](https://localhost:15672)
- Crie a fila: `amq.direct`
