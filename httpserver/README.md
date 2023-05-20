# HTTP Server

Load balanced via Railway's horizontal scaling

## Build

```bash
docker built -t mushufeels/httpserver .
```

## Environment Variables

Make sure to add the environment variables either in a .env file or into your PaaS console

1. AXIOM_API_TOKEN (api key)
2. AXIOM_DATASET (dataset name)
3. LOG_LEVEL (debug, info, warn, etc)
4. RABBITMQ_HOST (ip address)