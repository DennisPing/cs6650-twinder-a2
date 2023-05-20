# RabbitMQ Consumer

## Build and Push up to Docker Hub
```
docker build -t mushufeels/consumer .
docker push mushufeels/consumer:latest
```

## Pull down from Docker Hub into VM
```
docker pull mushufeels/consumer:latest
```

## Create .env file in VM
```
touch ~/consumer.env
echo "RABBITMQ_HOST={ip_address}" >> ~/consumer.env
echo "LOG_LEVEL=warn" >> ~/consumer.env
```

## Run container
```
docker run -d --name consumer --env-file ~/consumer.env -p 8080:8080 mushufeels/consumer
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