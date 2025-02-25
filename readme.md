Ao subir o Kafka, é importante subir pelo menos 3 brokers

zookeeper -> health check

### Tópicos:
 - Canal de comunicação responsavel por receber e disponibilizar os dados enviados para o Kafka
 - Offset: id da posição de cada mensagem
    - É possivel ler novamente uma mensagem através do Offset, pois as mensagens ficam salvas em disco e não em memória igual o Rabbitmq
 - O registro é composto pelos seguintes argumentos:
    - Header
    - Key
    - Value
    - Timestamp
 - Partições: Cada tópico pode ter uma ou mais partições para conseguir garantir a distribuição e resiliencia dos dados. Basicamente as mensagens irão ser distribuidas nas diferentes partições, isso torna mais rapido a leitura das mensagens além de evitar que um consumidor leia a mesma mensagem.
 - Keys: Define a partição que a mensagem irá ser enviada
 - Partições Distribuidas:
    - Replicator Factor -> define quantas cópias de partições deverá existir em cada broker, dessa forma, caso um broker cair eu tenho a réplica em outro broker. Isso garante a RESILIENCIA dos dados. Normalmente utiliza-se 2 ou 3 para o Replicator Factor, pois é custoso manter todas as partições.
        - Ex: Tópico Vendas | Replicator Factor = 2 
            - Broker A -> Partição 1 | Partição 3
            - Broker B -> Partição 2 | Partição 1
            - Broker C -> Partição 3 | Partição 2
 - Partition Leadership: Define qual broker é o líder de uma partição específica. O líder é responsável por todas as operações de leitura e escrita para essa partição. Caso um broker líder falhe, o Zookeeper detecta a falha e o Kafka promove uma das réplicas dessa partição em outro broker para ser o novo líder. Isso garante que as operações possam continuar sem interrupção. Quando um consumidor precisa ler uma partição, ele sempre se conecta ao broker líder dessa partição.
 - Producer: Garantia de entrega
    - Ack 0 (Sem garantia): O producer manda a mensagem e fodac se foi entregue ou não, é bom quando precisa mandar uma grande quantidade de dados sem se importar caso perca algum
    - Ack 1 (Garantia parcial): Manda a mensagem e o Leader do Broker confirma que recebeu a mensagem. Porém, não garante que a mensagem foi replicada nos Followers (Brokers backup) para garantir a resiliencia
    - Ack -1 (Garantia total - lento): Manda a mensagem > O Leander recebe a mensagem e envia para os Followers > Os Followers confirmam que receberam a mensagem > O Leander confirma que está tudo OK para o Producer
 - Producer: Indepotencia
    - O Kafka identifica que o producer mandou uma mensagem repetida, trata para descartar uma e que a mensagem tenha o index correto. Porém isso tem o custo de ser mais lento para publicar as mensagens.
 - Consumers e Consumer groups
    - Um consumidor pode consumir de várias partições
    - Consumer group, cria um grupo de consumidores para que as partições sejam distribuidas para os diferentes Consumers, caso não defina o Group, o Kafka vai funcionar como se eu tivesse um único consumer. 
    - Não tem como mais de um consumer do mesmo grupo ler a mesma partição
    - MELHOR DOS MUNDOS É TER A MESMA QUANTIDADE DE PARTIÇÕES E CONSUMERS.
 