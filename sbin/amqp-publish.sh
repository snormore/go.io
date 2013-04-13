#!/bin/bash
EXCHANGE="test-exchange"
BINDING_KEY="test-key"
amqp-publish -e $EXCHANGE -r $BINDING_KEY -C "text/json" -b "{\"SentAt\":`date +%s`, \"Body\":\"$1\"}"
