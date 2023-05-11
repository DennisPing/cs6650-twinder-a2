# RabbitMQ Consumer

## Build and Push up to Docker Hub
```
docker build -t mushufeels/rmqconsumer .
docker push mushufeels/rmqconsumer:latest
```

## Pull down from Docker Hub into VM
```
docker pull mushufeels/rmqconsumer:latest
```

## Create .env file in VM
```
touch ~/rmqconsumer.env
echo "RABBITMQ_HOST={ip_address}" >> ~/rmqconsumer.env
echo "LOG_LEVEL=warn" >> ~/rmqconsumer.env
```

## Run container
```
docker run -d --name rmqconsumer --env-file ~/rmqconsumer.env mushufeels/rmqconsumer
```

## Stop containers
```
docker stop {container_name}
```

## Restart stopped containers
```
docker start {container_name}
```

## View logs of containers
```
docker logs {container_name}
```

## Remove containers
```
docker rm {container_name}
```

## Remove images
```
docker image rm {image_id}
```