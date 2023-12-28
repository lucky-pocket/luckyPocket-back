#!/bin/sh

CONFIG_PATH="/home/ubuntu/deploy/docker-compose.yml"

sudo docker-compose pull

if [ ! -e /home/ubuntu/.initialized ];then
  echo "First time running"
  sudo docker-compose -f $CONFIG_PATH up redis mysql -d --wait
  touch /home/ubuntu/.initialized
fi
 
sudo docker-compose -f $CONFIG_PATH up app batch -d