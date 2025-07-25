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

## Commands

- Commandos

```
    go test .
```

## Links

- https://github.com/golang-standards/project-layout
