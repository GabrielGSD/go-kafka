# Kafka Connect

Kafka Connect é uma ferramenta robusta e escalável para integrar o Apache Kafka com outros sistemas. Ele facilita a ingestão de dados de diversas fontes para o Kafka e a exportação de dados do Kafka para outros sistemas. Kafka Connect oferece:

- **Conectores**: Plugins que permitem a integração com diferentes sistemas de dados.
- **Escalabilidade**: Pode ser executado em uma única instância ou em um cluster distribuído.
- **Facilidade de uso**: Configuração simples através de arquivos JSON ou REST API.
- **Transformações**: Permite transformar dados em trânsito com Single Message Transforms (SMTs).

Kafka Connect é ideal para pipelines de dados em tempo real, ETL (Extract, Transform, Load) e integração de sistemas heterogêneos.

### Worker:
No contexto do Kafka Connect, um worker é uma instância do processo do Kafka Connect que executa conectores e tarefas. Os workers podem ser configurados para operar em modo standalone (única instância) ou em modo distribuído (várias instâncias em um cluster). Aqui estão alguns pontos importantes sobre os workers:

- **Standalone**: Em modo standalone, um único worker executa todos os conectores e tarefas. Este modo é mais simples de configurar, mas não oferece alta disponibilidade ou escalabilidade.
- **Distribuído**: Em modo distribuído, múltiplos workers podem ser executados em um cluster. Isso permite a distribuição de carga de trabalho entre várias instâncias, proporcionando alta disponibilidade e escalabilidade.

Os workers são responsáveis por:
- Gerenciar a execução dos conectores e suas tarefas.
- Coordenar a distribuição de tarefas entre diferentes instâncias (no modo distribuído).
- Monitorar e reiniciar tarefas em caso de falhas.
- Aplicar transformações nos dados em trânsito, se configurado.

### Converter:
É responsável por transformar dados de um formato para outro. Ele é usado para garantir que os dados estejam no formato correto antes de serem processados ou armazenados. Por exemplo, um converter pode transformar dados JSON em um formato de objeto Java ou vice-versa. Isso é essencial para a integração de diferentes sistemas e para garantir a consistência dos dados ao longo do pipeline de processamento.
- Avro 
- Protobuf
- JsonSchema
- Json
- String
- ByteArray

### Dead Letter Queue (DLQ)
Uma Dead Letter Queue (DLQ) é um mecanismo de tratamento de erros usado no Kafka Connect para lidar com mensagens que não podem ser processadas corretamente. Quando uma mensagem falha no processamento, em vez de simplesmente descartá-la ou bloquear o pipeline, ela é movida para uma fila especial (DLQ) para:

- **Isolamento**: Separar mensagens problemáticas do fluxo principal
- **Análise**: Permitir investigação posterior dos erros
- **Reprocessamento**: Possibilitar tentativas futuras de processamento
- **Monitoramento**: Facilitar o acompanhamento de falhas

Quando há um registro inválido, independente da razão, o erro pode ser tratado nas configurações do conector através da propridade "errors.tolerance". Esse tipo de configuração pode ser realizado apenas para conectores do tipo "Sink".
- none: Faz a tarefa de falhar imediatamente. (Sem tolerancia a erros) (ESSE É O PADRÃO)
- all: Erros são ignorados e o processo continua normalmente. (Passa tudo, mas não sei que o erro aconteceu)
- errors.deadletterqueue.topic.name = <nome_do_topico> -> Diz que sempre que der um erro, a mensagem será jogada em um tópico, dessa forma eu sei qual mensagem deu erro e não será perdida. (MELHOR DOS MUNDOS)

