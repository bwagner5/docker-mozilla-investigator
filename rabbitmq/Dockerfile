FROM rabbitmq

COPY ca.crt /etc/rabbitmq/
COPY rabbitmq.key /etc/rabbitmq/
COPY rabbitmq.crt /etc/rabbitmq/

COPY rabbitmq.config /etc/rabbitmq/rabbitmq.config

COPY create_rabbitmq_config.sh /
RUN service rabbitmq-server start && sh /create_rabbitmq_config.sh && service rabbitmq-server stop

#CMD ["rabbitmq-server", "start"]
