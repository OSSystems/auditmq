# AuditMQ

* AuditMQ it's a simple 'auditor' of consistency, made for microservices architecture.

## Configuration

The configuration is pretty simple, basically you defines your 'data' and RabbitMQ configuration.
The idea behind this data, is a shared information between your services, in the example below, it's a
sales_count (quantity of sales) from sales-service (owner of this data), supposedly this data is shared
with your billing-service (that controls all communication with payment gateway) and needs to be always
synchronized.

```
dsn: "amqp://guest:guest@localhost:5672/" # RabbitMQ connection string
exchange: auditmq                         # RabbitMQ exchange
consumer_queue: audit_in                  # RabbitMQ consumer queue
data:                                     # Data configuration section
  sales_count:                            # Field in the report
    type: int                             # Is a type of this field,
    owner: sales-service                  # Service owner identification
    samples: 3                            # Quantity of samples considered safe for comparison
    replicas:                             # Replicas configuration section
      - billing-service:                  # Dependant service identification
          offset: 10                      # Offset is thought for expected data variations*
```
> Imagine you're deployed the sales-service and it works as well independently of billing-service, when billing-service
will have been deployed, the service can be a variation of sales-service.

## Running
```
make && ./auditmq
```
> If you needs more information, you can use the DEBUG environment variable to see more informations

## Service Report Data

The service report data contains the service identification and data in the json.

```
{
  "service": "sales-service",
  "data": {
    "sales_count": 10
  }
}
```
> Schema is available inside the package called **pkg**

## Comparison

The comparison runs only with the samples list are **filled** and considered **stable** without variance in integer sampling
for example. After that, the AuditMQ send to RabbitMQ a consistency report data for each service configurated and
global consistency status with hardcoded named with 'all' in the report.

## Consistency Status Report

Including the service identification and yours status, the **status** can be 0 for not synchronized, and 2 for synchronized.

```
{
  "service": "sales-service",
  "status": 2
}
```
> Schema is available inside the package called **pkg**

## Integration

AuditMQ can be integrated with log processing stack, in our environment is used with Fluentd and InfluxDB to build a simple
semaphore and store consistency log along the time.

![Grafana](docs/integration.jpg)
