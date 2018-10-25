dsn: "amqp://guest:guest@localhost:5672/"
exchange: uh.audit.tasks
consumer_queue: uh.audit.events
services:
  api-server:
    own_data:
      device_count:
        type: int
      deployments_count:
        type: int
  billing-service:
    matched_data:
      device_count:
        type: int
        integer_offset: 10
