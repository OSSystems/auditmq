dsn: "amqp://guest:guest@localhost:5672/"
exchange: uh.audit.tasks
consumer_queue: uh.audit.events
data:
  device_count:
    type: int
    owner: api-server
    replicas:
      - billing-service:
          offset: 10
      - usage-monitor:
          offset: 0
