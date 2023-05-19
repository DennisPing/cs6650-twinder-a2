# RabbitMQ Host

The RabbitMQ host that receives and sends messages

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

## How to Deploy on Arm64 Instance
```bash
docker pull arm64v8/rabbitmq:3-management

docker run -d --hostname rabbitmq-1 --name rabbitmq-1 -p 5672:5672 -p 15672:15672 arm64v8/rabbitmq:3-management
```

## What does that do?
1. Runs the container in detached mode
2. Sets the internal container's hostname to `rabbitmq-1` in case we ever want to run a cluster of rabbitmq's
3. Names the container `rabbitmq-1`
4. Exposes the ports 5672 (main connection) and 15672 (management page)
