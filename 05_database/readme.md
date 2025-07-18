# Contexto

Aula de como usar Banco de dados no Go(lang):

- Introdução a banco de dados

  - Comunicação com a base de dados
  - Nesse Módulo vamos fazer um CRUD completo
  - O Go tem uma comunicão nativa com a base de dados

- Preparando Base do Código

  - Os devs Go tenta usar o máximo de coisas nativas do GO, para não usar Library
  - Usamos mais o SQL puro do que ORM ou GORM (Go RM)
  - Criamos um docker compose para nos ajudar na criação

- Inserindo Dados no Banco

  - A comunicacão com a base de dados é simples
  - Abrimos a conexão, fazemos a lógica e depois fechamos a base de dados
  - "user:senha@tcp(localhost:3306)/goexperts"

- Alterando Dados no Banco de dados

  - Momento de atualizar o registro que alteramos
  - Usando Sanitização
  - Usando query select

- Trabalhando com QueryRow

  - Selecionar os dados atrabés do QueryRow
  - Começamos com a selecão de um produto
  - Função scan, vai rodar cada coluna e comparar com o produto

- Selecionando múltiplos registros

  - Selecionar vários produtos de uma vez

- Removendo Registro
  - removendo o registro da base de dados

## Commands

- Cria/levanta o compose

```
    docker compose up -d
```

- Para acessar o bash ou a base de dados usamos

```
    docker compose exec mysql bash
```

- Para acessar a base de dados

```
    mysql -uroot -p goexpert [nome_base_de_dados]
```

- Para criar uma tabela básica

```
    create table products (id varchar(255), name varchar(80), price decimal(10,2), primary key(id));
```

- Ver conteúdo na base de dados

```
    select * from products;
```

## Links
