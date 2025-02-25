Criando tópico:
 ```
 kafka-topics --create --topic=teste --bootstrap-server=localhost:9092 --partitions=3
 ```
 . kafka-topics: É a ferramenta de linha de comando do Kafka usada para gerenciar tópicos.
 . --create: Indica que você deseja criar um novo tópico.
 . --topic=teste: Especifica o nome do tópico que está sendo criado, neste caso, "teste".
 . --bootstrap-server=localhost:9092: Define o endereço do servidor Kafka ao qual a ferramenta deve se conectar. Aqui, está configurado para localhost na porta 9092.
 . --partitions=3: Define o número de partições que o tópico terá. Partições são unidades de paralelismo no Kafka, permitindo que o tópico seja distribuído e escalado.

 Lista tópicos:
  - kafka-topics --list --bootstrap-server=localhost:9092

 Detalhando tópico:
  - kafka-topics --bootstrap-server=localhost:9092 --topic=teste --describe
  Topic: teste    PartitionCount: 3       ReplicationFactor: 1    Configs: 
          Topic: teste    Partition: 0    Leader: 1       Replicas: 1     Isr: 1
          Topic: teste    Partition: 1    Leader: 1       Replicas: 1     Isr: 1
          Topic: teste    Partition: 2    Leader: 1       Replicas: 1     Isr: 1

  Consumindo e produzindo mensagens:
  - kafka-console-producer --bootstrap-server=localhost:9092 --topic=teste
  - kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste --from-beginning

  Consumer Groups:
  - kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste --group=x
  - kafka-consumer-groups --bootstrap-server=localhost:9092 --group=x --describe

  Navegando pelo Confluent Control Center:
  - http://localhost:9021/