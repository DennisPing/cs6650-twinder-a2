# AWS Linux 2 on Arm64

## If using an AWS Graviton Instance (arm64)

You cannot build a Docker image on an amd64 computer (Intel or AMD) and use it on arm64 instance. Either use `buildx` to do cross platform build, or just clone the repo and build it natively.

### Setup dependencies
```
sudo yum update
sudo yum -y install docker
sudo usermod -a -G docker ec2-user
newgrp docker
sudo systemctl enable docker.service
sudo systemctl start docker.service
```

### Build Docker image natively
```
git clone https://github.com/DennisPing/cs6650-twinder-a2.git
cd rmqconsumer
docker build -t mushufeels/rmqconsumer .
docker run -d --name rmqconsumer --env-file ~/config.env mushufeels/rmqconsumer
```