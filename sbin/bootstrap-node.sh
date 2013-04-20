#!/bin/bash
sudo su -
apt-get update
DEBIAN_FRONTEND=noninteractive apt-get -y -o Dpkg::Options::="--force-confdef" -o Dpkg::Options::="--force-confold" dist-upgrade
apt-get -y install vim git mercurial golang
apt-get -y install rabbitmq-server redis-server
APPS_DIR="/u/apps"
APP_NAME="go.io"
sudo mkdir -p $APPS_DIR
sudo chown vagrant $APPS_DIR
ln -s /vagrant $APPS_DIR/$APP_NAME