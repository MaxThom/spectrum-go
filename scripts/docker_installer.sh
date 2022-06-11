#!/bin/sh

sudo apt-get update -y
sudo apt-get upgrade -y

# To install the local domain
# sudo apt-get install avahi-daemon

# Installing Docker
curl -sSL https://get.docker.com | sh
sudo sh get-docker.sh
sudo usermod -aG docker ${USER}
groups ${USER}
sudo su - ${USER}
docker version
docker info

# Installing Docker-Compose
sudo apt-get install libffi-dev libssl-dev -y
sudo apt install python3-dev -y
sudo apt-get install -y python3 python3-pip
sudo pip3 install docker-compose
sudo systemctl enable docker