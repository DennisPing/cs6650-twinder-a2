# AWS Linux 2 on Arm64

## If using an AWS Graviton Instance (arm64)

You cannot build a Docker image on an amd64 computer (Intel or AMD) and use it on arm64 instance. Either use `buildx` to do cross platform build, or just clone the repo and build it natively.

### Setup dependencies
```bash
sudo yum update -y
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo usermod -aG docker ec2-user
# Log out and log back in here to ensure this takes effect
sudo systemctl enable docker
sudo systemctl start docker
```

### Build Docker image natively
```bash
git clone https://github.com/DennisPing/cs6650-twinder-a2.git
cd consumer
docker build -t mushufeels/consumer .
docker run -d --name consumer --env-file ~/consumer.env -p 8080:8080 mushufeels/consumer
```