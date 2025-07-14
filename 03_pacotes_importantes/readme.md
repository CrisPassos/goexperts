# Pacotes Importantes

Aula de Pacotes Importantes do Go(lang):

- Manipulação de Arquivos

  - [os](https://pkg.go.dev/os)
  - Criação de arquivos .txt
  - para ler o arquivo e um o ReadFile (ele ler como byte e depois eu preciso converter para string)
  - Quando precisamos ler o arquivo em partes usamos a library [bufio](https://pkg.go.dev/bufio)

- Realizando chamada Http

  - UMA DAS COISAS MAIS IMPORTANTES DO GO
  - [net/http](https://pkg.go.dev/net/http)

- Defer

  - [defer](https://pkg.go.dev/builtin#defer)
  - Executa uma função quando a função que o chamou termina
  - Útil para fechar arquivos, conexões, etc.
  - atrasar a execução de uma função até que a função que o chamou retorne

- Trabalhando com Json

  - [encoding/json](https://pkg.go.dev/encoding/json)
  - Serialização e desserialização de dados JSON
  - Converte structs para JSON e vice-versa
  - Manipulação de dados JSON de forma fácil
  - Exemplo de uso: `json.Marshal` e `json.Unmarshal`
  - Tags JSON em structs para personalizar a serialização

- Busca CEP

  - https://viacep.com.br/
  - converter o JSON retornado em uma struct https://transform.tools/json-to-go
  - Chamada HTTP + JSON + Struct + Defer + Manipulação de Arquivos
  - gerei o build com o comando `go build -o cep main.go`
  - executei o comando `./cep 01001000` para buscar o CEP

- Iniciando com HTTP

  - Criar um servidor HTTP simples e buscar o CEP em vez de usar o terminal
  - vamos implementar a interface `http.Handler` para criar um servidor HTTP
  - Vamos usar o ResponseWriter para escrever a resposta HTTP
  - HandleFunc para definir rotas e métodos HTTP

-

## Commands

```
    curl localhost:8080
```

## Links

- converter o JSON retornado em uma struct https://transform.tools/json-to-go
