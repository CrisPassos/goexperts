# APIs

Aula de Apis no Go(lang):

- Falando sobre APIs

  - Criando Apis através do Go(lang)
  - Disponibilizar apis Rest, Webservers etc
  - Roteadores, Muxs, garantindo as melhores práticas de desenvolvimento
  - Nosso desafio vai ser criar uma autenticação de user e um CRUD de produtos
  - Tudo desenvolvido nas boas práticas do go, inclusive na criação dos dicionários

- Estruturas diretórios

  - Seguir o exemplo de organização do repo do github chamado golang-standars/project-layout
  - é um guia, mas pode ser modificado de projeto para projetos
  - /api - guarda a documentação das apis tipo swagger
  - /internal - onde roda a aplicação, o código da aplicação, o codigo aqui não pode ser reutilizado em outras aplicações, é privado e especifico da aplicação, sem importação
  - /pkg - pasta do package, oposto do internal, tudo que for utilizado para externos tipo libraries que podem ser utilizadas em diversos projetos
  - /cmd - contém os arquivos do projeto, o executavel do projeto fica nessa pasta, main.go, dentro da pasta cmd eu coloco o nome da aplicação
  - com essas pastas conseguimos fazer o código, gerar o build e criar códigos genericos, baixados por outras APPS
  - /configs - configurações padrões de inicialização, padrão de variaveis de ambientes, como subir a aplicação etc... booting da aplicação
  - /test - pasta que contém aquivos que auzilia os testes de unidades, documentação, mock, stubs, declaração do postman, não necessáriamente precisam ser aquivos .go ou de testes
  - vamos trabalhar com a openAPI

- Criando arquivo de configuração

  - Primeiro passo executar o go mod init
  - vamos iniciar as configs do sistemas, comicação com a base de dados, apis autenticação etc

- Finalizando configuação

  - vide code
  - vamos usar o package chamado viper

- Outras Possibilidades de configuracão

  - Poderiamos deixar as configuracões no modo privado para que o user não possa alterar as informações diretamente
  - nesse caso criamos getters and setters

- Criando entidade User

  - camos trabalhar na pasta internal
  - usamos o bcript que gera um hash para salvar a senha, jamais salvamos a senha pura e sim através de hash

- Testando User
- Criando entidade Product
- Testando Product
- Criando UserDB
- Testando criação do usuário
- Finalizando teste de UserDB
- Criando Principais métodos de PRODUCTDB
- Criando FindAll com paginação
- Testando ProductDB
- Finalizando testes de ProductDB
- Criando Handlers para criar produto
  - Para esse exemplo vamos criar uma camada diretamente para conexão com os dados
- Testando endpoint de criação de Product
- Ajustando package para os handlers
- Web frameoworks vc Frameworks
  - O mux padrão do Go(lang) não permite trabalhar com variaveis dentro das URLs
  - Roteadores implementações que permite criar URLs mais dinamicas, ou adicionar server
  - Frameoworks: echoes, fiber (parecido com o node), gin gonic
  - Frameworks: preocupados com manter na natividade e ajudar no padrão de desenvolvimento, tipo laravel: exemplos buffalo
- Roteadores
  - para esse projeto vamos usar roteadores, o roteador implementa interfaces para nosso mux server
  - apenas com o roteador conseguimos usar todas as assinaturas do GO, o mais tradicional é o GorilaMux
  - O roteador trabalha com as rotas e handler e depois passamos para o roteador
  - vamos usar o go chi

## Commands

- Acessar o SQLit3

```
    sqlite3 test.db
```

## Links

- https://github.com/golang-standards/project-layout
- https://echo.labstack.com
- https://gofiber.io
- https://gin-gonic.com
- https:/gobuffalo.io
- https://www.gorillatooklkit.go
- http://github.com/go-chi/chi
