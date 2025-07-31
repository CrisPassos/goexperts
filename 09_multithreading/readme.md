# Multthreading

Aula de Multthreading no Go(lang):

- Iniciando com Processos
  - Concorrencia, parelelismo, go routes
  - Quando temos aplicações, temos processos rodandos por de trás
  - Os processos são contidos pelo sistema operacional
  - Um processo não pode atingir outro
  - O processo A não consegue ler o que está no processo B, pq eles são independentes
  - Thread são linhas (fios), outros momentos de processamentos principal
  - Threads: estão relacionados aos processos, como as thears são dos mesmo processos elas estão contidas e elas tem permissão de alterar os valores, pq elas compartilham memoria
  - Exemplo Thread A e Thread B podem alterar a varíavel A = 10
- Introdução a concorrência e Mutex
  - compartilhar memória é uma dor de cabeça pq caimos uma condição de Race Conditional (isso significa uma Thread A pode estar alterando ao mesmo tempo uma Thread B)
  - geramos conflito de memória, pq acessamos valores do mesmo local de memória e por isso dá muito conflitos
  - existem diversas formas de trabalhar com multh é através de Mutex (mutua exclusion), cria um LOCK na primeira thread que começou o trabalho, sempre trabalhamos com LOCK e UNLOCK
- Concorrência vc Paralelismo vc GO
  - A concorrência ativa o parelismo, trabalhar de forma concorrente deixa as coisas mais rapidas
  - Concorrência é quando você tem várias tarefas que podem ser executadas ao mesmo tempo, mas não necessariamente estão sendo executadas ao mesmo tempo.
  - Paralelismo é quando você tem várias tarefas que estão sendo executadas ao mesmo tempo
  - Eu posso ter varias CPU rodando concorrencia e isso vira um paralelo
  - O Go permite mudar quantas CPUs
- Multithreading
  - Todas as vezes que meu sistema precisar de uma thread ele vai fazer uma chamada ao SO e vai gerar uma nova Thread
  - Toda chamada do processo ao so e o so ao threads tem um custo
  - 1 mega par a Thread é muiiiitoo caro
- Schedulers
  - Mudanças de Contexto tem um custo
  - o Scheduler gerencia a mudança de contexto, gerencia as tarefas no nosso SO
  - Preemptivo: Tempo determinado por cada tarefa, interrompe a tarefa que esta sendo executada para que a outra seja executada no lugar
  - Cooperativo: espera uma tarefa terminar para começar a outra
- Go e suas green threads

  - quando as liguagens foram criadas normalmente tinhamos comente uma CPU1
  - O Go surge num contexto diferente, ela tem um RUNTIME (motor para a linguagem funcionar)
  - No final do dia ela é bem inteligente, ao inves de ir ao SO criar um Thread e ele cria um próprio esquema de gerenciamente de threads
  - Green Threads: Threads in Userland (cria pequenos processos ou subprocesso), faz com que a linguagem trabalhasse como efeito
  - O Go tem seu próprio Scheduler, ele precisa organizar a execução, eles trabalham em forma Cooperativa (ativa um, desativa outro)
  - Eu posso trabalhar com milhares de threads, nao há syscal, não gero outra thread no SO

- Iniciando com Go Routines

  - para gerar uma thread colocamos um prefix go antes da função
  - permite que usamos os recursos da threads
  - chocada que só precisa do Go na frente da func
  - rodar simultaneamente diversas tarefas
  - é possivel trabalhar com funções anonimas

- Trabalhando com Wait Groups

  - Wait groups: separado em 3 partes:
    - 1. adicionar quantidade de tarefas/operaçoes
    - 2. Informar que você terminou uma operação
    - 3. Esperar até que as operações sejam finalizadas (uma vez que eu sei quantas operações vão ser executada eu consigo manter o sistema parado)

- Problema simples de concorrência

  - trabalhando com LOCK e UNLOCK
  - para simular concorrencia usar o sistema apache AB(BenchMark)
  - Entendendo Mutex E Operações Atômicas

- Iniciando com Channels

  - Channels ou Canais: Nos ajudam fazer comunicação entre threads, seguranção para uma Thread saber o momento em que ela pode trabalhar com um determinado dado, como a Thread A passa um valor para a Thread B
  - Pontos de comunicação para que as threads se comuniquem entre ela
  - O canal vai ter Valor (vazio ou cheio)
  - Nunca vamos ter uma thread trabalhando com a mesma informação pq ela aguarda o canal publicar algo

- Forever

  - Aplicações comum em GO
  - Gerar deadlocks
  - Tipo loop infinito

- Iterando com range

  - for e range, cada vez que o for for executar ele preenche o canal e depois tem um range que vai esvaziar
  - enquando o range não liberar o canal o for não pode ser executado

- Range with WaitGroup
- Channel Directions

  - O channel ou recebe o dados ou ele joga para uma variavel ou para alguma coisa
  - Só temos duas direções receber ou entregar (receiver or deliver)

- Criando um load Balancer

  - balanceador de carga, receber a carga e equilibrar ela, criamos workers rotinas que le em paralelo para que possamos processar rapido
  - por exemplo, servidores web
  - meu balance joga a carga para os workers

- Trabalhando com Select

  - tipo um switch case para canais

- Canais com Buffer
  - é possível que o canal receba mais de uma informação ao mesmo tempo, por exemplo usando os buffers
  - se eu tiver o buffer de 5 eu posso colocar 5 informações até o canal ficar cheio

## Commands

- Verificando se estamos com algum Data Race

```
    go run -race main.go
```
