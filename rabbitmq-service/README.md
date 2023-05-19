# RabbitMQ Service

The RabbitMQ service that hosts the message queues


## How to Deploy on Arm64 Instance
```
docker pull arm64v8/rabbitmq

docker run -d --hostname rabbitmq-1 --name rabbitmq-1 -p 5672:5672 -p 15672:15672 arm64v8/rabbitmq:3-management
```

## What does that do?
1. Runs the container in detached mode
2. Sets the internal container's hostname to `rabbitmq-1` in case we ever want to run a cluster of rabbitmq's
3. Names the container `rabbitmq-1`
4. Exposes the ports 5672 (main connection) and 15672 (management page)
