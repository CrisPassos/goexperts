# GraphQL

Aula SQLC no Go(lang):

- Mais sobre banco de dados
- Trabalhando com migrations
  - mudar o schema do banco de dados de uma forma dinamica
  - arquivo up todos os sqls que vai rodar na primeira migration
  - arquivo down todos os sqls que vai destruir a migration
- Falando sobre SQLX
- Iniciando com SQLC
- Criando CRUD
- Finalizando CRUD
- Iniciando com transações
- Implementando Método para transação
- Vendo a transação funcionar
- Ajustando tipo de Campo
- Exibindo dados com Join

## commands

- começando com a migrate

```
  migrate create -ext=sql -dir=sql/migrations -seq init
```

- rodar o migration up

```
  migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up
```

- rodar o migration down

```
  migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down
```

- verificar base de dados

```
  docker compose exec mysql bash
```

```
  mysql -uroot -p courses
```

## Links

- https://sqlc.dev
- https://github.com/golang-migrate/migrate
- https://github.com/jmoiron/sqlx
