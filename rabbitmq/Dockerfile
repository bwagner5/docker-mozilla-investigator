FROM rabbitmq:3.5.6-management

COPY ca.crt /etc/rabbitmq/
COPY rabbitmq.key /etc/rabbitmq/
COPY rabbitmq.crt /etc/rabbitmq/

COPY rabbitmq.config /etc/rabbitmq/rabbitmq.config

COPY create_rabbitmq_config.sh /
COPY rabbitmq-entrypoint.sh /
#RUN service rabbitmq-server start && sh /create_rabbitmq_config.sh && service rabbitmq-server stop

ENTRYPOINT ["/rabbitmq-entrypoint.sh"]

CMD ["rabbitmq-server"]
