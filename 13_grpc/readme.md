# GraphQL

Aula gRPC de no Go(lang):

- Iniciando com gRPC
  - remote procedure call (usando outros tipos de protocolos)
  - a maioria dos sistemas internos do Google usam o gRPC em vez do REST
  - mais rápida e segura para comunicação entre sistemas
- Conceitos Iniciais
  - é um framework desenvolvido pela google que tem o objetivo facilitar o processo de comunicação entre sistemas de uma forma extremamente rápida, leve, independente de linguagem.
  - Faz parte do CNCF ( Cloud Native Computing Foundation - Kubernets, Telemetry )
  - Em quais casos podemos utilizar?
    - idela para microsserviços (cominucação entre o sistemas o tempo todo)
    - mobile, browsers e backend
    - geração das bibliotecas de forma automática de código
    - streaming bidirecional utilizando HTTP/2 (comunicação binária, manter o canal de atualização aberto, enviar diversas responses ou requests)
  - Linguagens (Suporte oficial)
    - gRPC-GO
    - gRPC-JAVA
    - gRPC-C (C#, C++, Python, Ruby, PHP, Objective C, Node.js, Dart)
- gRPC HTTP2 e Protocol Buffers
  - RPC (Remote Preocedure Call)
    - Client (server.soma(a,b)) -> server (func soma(int a, int b){})
    - Fazer uma chamada remota que está em outro sistema
  - Protocol Buffers
    - o gRPC trafega no protocolo http2
    - Protocol buggers are Google's language-neutral, platform-neutral, extensible, mechanism for serializeing structured data - think XML, but smaller, faster and simplers"
    - trabalha com schemas, binários e protocolos existentes
    - necessário fazer a serialização e a desarrializacão dos binários
  - Protocol Buffers vs JSON
    - Arquivos binários x JSON
    - Processo de serializacão é mais leve (CPU) do que JSON
    - Gasta menos recursos de rede
    - Processo é mais veloz
  - HTTP/2
    - Nome original criado pela Google era SPDY
    - Lançado em 2015
    - Dados trafegados são binários e não texto como no HTTP 1.1
    - Utiliza a mesma conexão TCP para enviar e receber dados do cliente e do servidor (Multiplex)
    - Server Push (assests)
    - Headers são comprimidos
    - Gasta menos recurso de red
    - Processo é mais veloz
- Formatos de comunição
  - gRPC - API "unary"
    - Client -> Request <-> Reponse -> Server
  - gRPC - Api "Server Streagming"
    - API - Application Program Interface
    - Cliente -> Request <> Response(devolve muitas repostas) -> Server
    - respostas parciais, ele vai enviando as informações antes de finalizar completamente
  - gRPC - API "Client Streaming
    - faz o oposto
    - Client -> Request(envia várias requests) <> Response (volta uma única) -> Server
  - gRPC - API "Bi directional streaming"
    - tanto o cliente quanto o serve envia multiplas requests ou responses
    - cliente -> request(varios) <> response (varios) -> server
- REST vs gRPC
  - Texto/json
  - unidirecional
  - alta latência
  - sem contrato (maior chances de erros)
  - sem suporte a streaming (Request/Response)
  - Design pré-definido (POST, PUT, DELETE ...)
  - Bibliotecas de terceiros
  - VERSUS
  - Protocol bufers
  - Bidirecional e Assincrono
  - Baixa latência
  - Contrato definito (.proto)
  - Suporte a streaming
  - Design é livre
  - Geração de código
- gRPC vs Protocol buffers
- Fazendo setup do projeto
- Fazendo geração de código com PROTOC
- Implementando Create Category
- Criando servidor gRPC
- Interagind com evans
- Criando CategoryList no ProtoFile
- Listando categories
- Buscando uma categoria
- Trabalhando com stream
- Trabalhando com streams bidirecionais

## commands

- gerar o protoc

```
     protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```

```
    evans -r repl
```

## Links

- https://grpc.io
- https://protobuf.dev/
- github.com/ktr0731/evans
