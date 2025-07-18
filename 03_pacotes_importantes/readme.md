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

- Manipulando Headers

  - Todas as requisições HTTP possuem headers
  - Headers são metadados que descrevem a requisição ou resposta
  - Podemos manipular headers usando o ResponseWriter

- Criando funções para BuscaCep

  - Vamos criar uma função que recebe o CEP como parâmetro
  - A função irá buscar o CEP e retornar os dados em formato JSON
  - Vamos usar a biblioteca `encoding/json` para serializar os dados

- Finalizando resposta para o servidor HTTP

  - Vamos finalizar a resposta HTTP com o status code 200 OK
  - Usar o ResponseWriter para escrever os dados JSON na resposta
  - Garantir que o Content-Type seja application/json

- ServeMux

  - ServeMux é um multiplexer de requisições HTTP
  - Permite definir rotas e métodos HTTP de forma organizada
  - Usar ServeMux para gerenciar diferentes endpoints da API
  - Toda vez que eu dou um ListenAndServe, eu subo um multiplexer padrão do GO
  - O Multiplexer padrão do GO é simples, mas podemos criar o nosso próprio multiplexer atraves do ServeMux
  - Usamos o ServerMux quando queremos criar novas rotas e endpoints para a nossa API, adicionar um novo handler, etc.

- FileServer

  - FileServer é um handler que serve arquivos estáticos
  - Podemos usar o FileServer para servir arquivos de um diretório específico
  - Útil para servir arquivos como HTML, CSS, JS, imagens, etc.
  - Vamos usar o FileServer para servir arquivos estáticos da nossa aplicação
  - Trabalhamos com arquivos estáticos, como HTML, CSS, JS, imagens, etc.

- Iniciando com templates

  - Vamos iniciar com templates para renderizar HTML dinâmico
  - Usar a biblioteca `html/template` para trabalhar com templates
  - Criar templates HTML e renderizar dados dinâmicos neles
  - HTML dinamico, com dados dinâmicos da base de dados
  - Por exemplo, emails personalizados, páginas web dinâmicas, etc.

- Template Must

  - Template permite que o erro e o sucesso volte no mesmo local
  - Facilita na parte de criação de erros
  - É uma função auxiliar da biblioteca text/template usada para encapsular erros de parsing de template

- Templates com arquivos externo

  - Criação de um arquivo externo para mostrar as informações dinamicas
  - Comunicação entre o ficheiro go e html

- Templates com Webserver

  - Vamos fazer tudo o que estamos fazendo, mas em um webserver
  - Utilizado para quando queremos trabalhar com arquivos e não somente webservers

- Compondo Templates

  - Criando diversas páginas onde só muda o conteúdo
  - Matém cria um header, body, conteudo, etc...

- Mapeando funções nos templates

  - usando html/template
  - usar o HTML sabemos que a informação vai ser parseada do tipo html
  - isso evita ataques, além de ser o padrão
  - serve para fazer transformações ou executar algumas ações
  - podemos usar as funções
  - Funciona como os pipes do angular

- HTTPClient com Timeout

  - Fazendo requisições externas com o HTTP (GET, POST, DELETE, PUT)
  - Quando trabalhamos com sistemas precisamos pensar em como deixar o sistema performática
  - Devemos estabelecer os limites de chamadas externas

- Trabalhanco com POST

  - Usamos o método Post
  - Para passar o body usamos um slice de bytes, essa é a grande diferença

- Customizando Request

  - usamos o metódo NewRequest para fazer isso
  - podemos colocar headers e mais customização
  - usamos o metódo DO também
  - usar baer token

- Trabalhando com HTTP usando Contextos
  - Pacote de contexto, permite passar as informações para diversas chamadas no sistemas
  - Podemos cancelar contextos
  - Usamos os contextos principalmente quando estamos trabalhando com contexto

## Commands

```
    curl localhost:8080
```

## Links

- converter o JSON retornado em uma struct https://transform.tools/json-to-go
- extensão Thunder Client para testar APIs (https://marketplace.visualstudio.com/items?itemName=rangav.vscode-thunder-client)
- documentação de templates: https://pkg.go.dev/html/template
