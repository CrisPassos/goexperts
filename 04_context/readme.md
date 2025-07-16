# Contexto

Aula de Contexto do Go(lang):

- Introdução aos Contextos

  - um dos pacotes mais importantes do GO
  - realizado para fazer chamadas http
  - Go Routes
  - Cenário: se algo exterior finalizar sua tarefa vc para o que estava fazendo
  - Contexto controla o que a aplicação está vivendo e se precisar para algo ele para
  - Guardar informações para que possa ser resgatada em outras áreas da aplicação
  - Permite trabalhar com key/value
  - usados desde chamadas http, consultas em base de dados, etc, etc

- Entendendo conceitos básicos

  - Começaremos com um exemplo simples de booking de hotel, temos um tempo limite para fazer o book

- Contexto utilizando server HTTP

  - Agora vamos trabalhar com requisições HTTP
  - Para isso vamos clicar um servidor HTTP CLient

- Context no lado Client

  - vamos trabalhar com um client e o server, já temos nosso server e agora vamos criar nosso cliente
  - é possível controlar o contexto do lado do cliente e do server

- Context WithValue
  - Trabalharemos com Key/Value
  - O contexto pode passar valores de um lado para o outro
  - Informações de metadados, em vez de passar por funções posso passar as coisas via contexto

## Commands

```
    curl localhost:8080
```

## Links
