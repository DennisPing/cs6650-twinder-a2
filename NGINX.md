# For historical purposes

## Create internal Docker network
```
docker network create my_network
```

## Run Container #1
```
docker run -d --name consumer-1 --network=my_network \
--env-file rabbitmq.env \
-e LOG_LEVEL=warn \
mushufeels/a2-rmqconsumer
```

## Run Container #2
```
docker run -d --name consumer-2 --network=my_network \
--env-file rabbitmq.env \
-e LOG_LEVEL=warn \
mushufeels/a2-rmqconsumer
```

## Run Nginx
```
docker run --name nginx-proxy -v ~/nginx.conf:/etc/nginx/nginx.conf:ro -d -p 80:80 --network=my_network nginx
```