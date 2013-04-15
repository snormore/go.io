#!/bin/bash
sudo apt-get update
sudo apt-get install -y vim git mercurial golang
sudo apt-get install -y rabbitmq-server redis-server
APPS_DIR="/u/apps"
APP_NAME="go.io"
sudo mkdir -p $APPS_DIR
sudo chown vagrant $APPS_DIR
ln -s /vagrant $APPS_DIR/$APP_NAME