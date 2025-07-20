# Banco de Dados

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

- Iniciando com GORM

  - ORM: library que ajuda a trabalhar com sql
  - Ajuda nos relacionamentos 1to1, 1toMany, HasBelong
  - Cenários que o GORM ajuda: has one, has many, belongs to, many to many, polymorphism, single-table inheritance
  - Eager loading with preaload -> trabalha com paginação
  - Consigo inserir informações em batch
  - Auto migrations, define as structs e ele cria as tabelas de acordo com o struct que eu defini

- Configurando e criando operações

  - Precisa ser instalado é bem documentado
  - Basicamente ele usa muito as structs com base para criar as operações

- Realizando primeiras consultas

  - Para buscar os produtos podemos usar os metodos First

- Realizando consultas com WHERE

  - Fazendo consultas com Where
  - Possibilidade de trabalhar com paginação,etc,etc

- Alterando e removendo registros

  - vide código

- Trabalhando com Soft delete

  - gorm.model: gerencia alguns status da base de dados
  - created_at, updated_at, deleted_at
  - nenhum registro é deletado, apenas arquivado

- Belongs TO

  - Trabalhando com relacionamentos
  - Relacionamento entre Tabelas
  - Eu tenho um produto que pertence a uma categoria
  - Um produto que pertence a varias categorias

- Has One

  - Relação um para um
  - um livro tem um unico isbn

- Has Many

  - Quando um registro pode ter mais de um outro tabela
  - Uma categoria pode ter muito produtos

- Pegadinha no has Many

  - Como trabalhar com Has Many e Has One juntos

- Many To Many

  - Uma Table pode estar em diferentes tabelas e vice e versa
  - precisamos de uma tabela intermediaria

- Lock otimista vs Pessimista
  - Bloqueio pessimista vs lock otimista
  - lock otimista, versiona quando alguém está fazendo uma alteração no sistemas
  - fazemos os locks quando vamos fazer muitas transações quando náo há muitas concorrência
  - Lock pessimista: lockar a tabela, não deixa mais ninguém alterar, isso é feito através da base de dados, para resolver problemas com concorrências
  - para trabalhar com grandes aplicações
  - Nesse exemplos vamos trabalhar com o cenário pessimista usando o GORM

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

- Deletar Tabela

```
    drop table products;
```

- Show tables

```
    show tables;
```

- Instalação do GORM

```
    go get -u gorm.io/gorm
```

## Links

https://gorm.io
