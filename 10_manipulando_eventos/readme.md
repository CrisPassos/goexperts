# Manipulando Eventos

Aula de Manipulando Eventos no Go(lang):

- Introdução a Eventos

  - Eventos e Mensageria
  - Eventos são coisas que aconteceram no passado
  - No desenvolvimento do SW é realizar uma ação que significa um evento (registrar, inserir, pintar)
  - Podemos reagir com nossos eventos

- Elementos táticos de contexto de eventos

  - Evento (carregar dados)
  - Operações que serão executadas quando um evento é chamado
  - Gerenciador dos nossos eventos/operações - quando o cliente é criado eu envio um email, fire ou dispatch

    - registrar os eventos e suas operações
    - Despachar / Fire no evento para suas operaçóes sejam executado

- Criando Interfaces da solução
- Registrando Eventos
- Criando Suite de testes
- Testando Registro de Handlers Repetidos
- Implementando e testando método Clear
- Implementando e testando o método Has
- Implementando método Dispatch
- Revisitando slices
- Removendo Handlers
- Adicionando go routine no event dispatcher
  - trabalhando com eventos paralelos
- Utilizando go routines no Dispatcher
- Instalando RabbitMQ
- Entendendo o RabbitMQ
  - Visa facilitar os multiplos request feitas por segundo
  - Funciona com um pool, uma fila de banco
  - Intermediador entre as mensages, as mensagens ficam contidas no sistema até o momento de ser usado
  - primeira coisa que temos é o producer ou produtor que envia a mensagem para outro servidor
  - consumer: o consumidor vai receber a mensagem do producer e dizer que foi OK
  - eu tenho uma fila queue no meio do producer e consumer, ou chamadas de transações
  - Quando uma mensagem cai na transacoes ele fica aguardando o cosumer estiver livre para receber
  - No rabbitmq o producer não envia a mensagem diretamente ao producer e sim para um exchanger (roteador)
  - Eu mando mensagem para a exanger e a exanger manda mensagem para essa fila
  - direct, fanout, topic (redireciona o exchange para a mensageria correta)
  - Mensagem/Condicao/Tempo
  - Elementos Produtor - Exchange - Mensage - Consumer
- Consumindo mensages do RabbitMQ
- Criando consumidor na prática
- Produzindo e consumindo mensagens
- Acessar o RabbitMQ Create Queue - Publish Mesages
- Exchanges - amg.direct - bind (procurar a fila)

## Commands

- checa o status

```
    docker compose ps
```

- para acessar o rabbitmq

```
  http://localhost:15672/#/
```

## Links

- tryerabbitmq.com
