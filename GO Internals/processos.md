# Go Internals - Processo

Aulas mais teoricas sobre o Go(lang):

## Processos:

- Instância de um programa em execucão
- Componentes para o processo funcionar é preciso de componentes organizados
  - Endereçamento (região da memória dedicada a um processo)
  - Contextos
    - Conjunto de dados que o SO "salva"para gerenciar um processo (para um processo rodar em determinado aspecto eu preciso executar esse Contexto)
      - EX.: Programa Counter (PC) ou Instruction Pointer (IP)
      - Possui o enderco da próxima instrução que o precessador irá executar
      - Auxília no Contexto Switch
  - Registro do Processador
    - Áreas que temporariamente armazenam no CPU dados e endereços para realizar a execução
    - Dados
      - Ex: Realiza operações aritméticas e lógicas
    - Registro de Endereços
      - Armazenamento em memória, incluindo stacks pointers
      - Ex: Ao acessar uma variável o CPU possui um registro na memória para guardar seu valor

## HEAP

      - Utilizado para alocação de memória dinamica. Cresce e encolhe em tempo de execução conforme a necessidade de mais ou mennos espaços
      - conforme eu vou lendo mais os dados vão indo para a memória heap, como ela é dinamica, ela vai crescendo e encolhendo conforme a necessidade
      - Ex: Alocação de memória para estruturas de dados como listas, árvores, etc
      - Tipo um garbage collector, que vai limpando a memória quando não é mais utilizada
      - O custo da HEAP é maior que o custo da stack, pois a HEAP precisa de mais gerenciamento

## Stack

      - Utilizada para armazenar variáveis locais e chamadas de função
      - Armazena informaçõees de controle de chamadas de função, como endereços de retorno, e parametros de função. Segue uma estrtutura LIFO (Last In, First Out - Ultimo a entrar, primeiro a sair)
      - Trabalhar com stack é mais eficiente, quando eu paro de usar a função ela é automaticamente liberada
      - Stack Overflow
        - Ocorre quando a pilha de chamadas de função excede o limite de memória alocada para ela
        - Ex: Recursão infinita ou uso excessivo de memória em chamadas de função
      - Ex: Variáveis locais, parâmetros de função e endereços de retorno

## Registros de Status/Flags

- Tudo que estamos falando agora é sobre processos, um dos componentes muitos importantes, são as flags
- Fornecem os status recentes das operações realizadas pela CPU
- Trabalha através de bits específicos(flags)
  - Ex:
    - Flag Zero (Z): Resultado de uma operação o qual o resultado é zero. Decide o fluxo do programa baseado nesse valor
    - Flag Signal (S) ou Negative (N): Indica o resultado de uma operação positiva ou negativa
    - Flag Overflow: Produz resultados além da capacidade

## Ciclo de vida de um processo

- Creation:
  - um novo processo é criado quando um programa solicita a execução de um processo, por meio de chamadas de sistema como fork() no UNIX/LINUX ou CreateProcess() no Windows
- Executtion:
  - O processo está ativamente sendo executado pela CPU. Pode alterar entre os estados de "executando" e "pronto" (para ser executado)
  - Waiting / Blocked:
    - O processo é suspenso e colocado em espera até que um evento externo ocorra. Comum em operações de I/O, onde o processo aguarda pelo término de uma leitura de disco ou recebimento de uma entrada de rede.
  - Termination: O processo completa sua execução ou é forçadamente terminado.
    - Exit: Conclusão bem-sucedida do processo após completar suas instruções
    - Killed: Interrupção por um erro de execução ou por ser terminado por outro processo (por exemplo, através do comando "kill")

## Criação de um novo Processo no SO

- Unix/Linux:
  - fork()
  - Clona o processo atual
  - Gerado um processo filho
  - fork() retorna um valor diferente para o processo pai (PID)
  - Processo pai e filho sãp quase idËnticos, porém os valores na mem'proa são copiados para outro endereçamento separados e independentes
  - O processo pai recebe um PID(valor inteiro positivo) do filho quando o fork() é chamado
  - Processo filho retorna o PID 0 - Indicando que é um processo filho.
  - Agora que o for() foi realizado, cada processo pode seguir para caminhos diferentes

## Gerenciamento de Processos

- Como o SO gerencia os processos?
- Scheduler (agendador): manda o CPU processar o processo que queremos
  - Decide qual processo será executado
  - Alternar entre processos
  - Possui diversos algoritmos para tentar maximizar o uso da CPU
  - O GO tem seu próprio Scheduler
  - Scheduler pode:
    - Selecionar processos de uma fila que estão "ready queue"
    - Escolhe os processos prontos para rodar
    - Alocar CPU: Mudança de estado: Ready to Running
    - Retirar CPU: I/O, etc
  - Tipos de Scheduler:
    - Colaborativo / Cooperativo
      - Processos que estão sendo executados tem controle quando liberam a CPU para outros processos
      - O scheduler vai esperar o processo terminar até o fim para iniciar outro
    - Preemptivo
      - SO tem a capacidade de interromper um processo em execução e ceder o uso da CPU para outro processo. Trabalha de forma mais "justa"
      - Conseguimos interromper o processo ou então voltar o processo da onde ele parou
  - Pontos a serem levados em consideração:
    - Processos cooperativos: Processos podem monopolizar a CPU
    - Processos preemptivos: Muitas mudanças de contexto (context switching)
