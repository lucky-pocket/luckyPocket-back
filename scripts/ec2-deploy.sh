#!/bin/sh

CONFIG_PATH="/home/ubuntu/deploy/docker-compose.yml"

sudo docker-compose pull

if [ ! -e /home/ubuntu/.initialized ];then
  echo "First time running"
  sudo docker-compose -f $CONFIG_PATH up -d
  touch /home/ubuntu/.initialized
else 
  sudo docker-compose -f $CONFIG_PATH up app -d
  sudo docker-compose -f $CONFIG_PATH up batch -d
fi