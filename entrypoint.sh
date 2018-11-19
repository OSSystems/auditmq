#!/bin/sh

echo $CONFIG | base64 -d > config.yaml

/usr/local/bin/auditmq
