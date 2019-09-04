#!/usr/bin/env bash

mkdir certs

rm -rf certs/*

openssl req -new -nodes -x509 -out certs/server.pem -keyout certs/server.key -days 3650 \
    -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=www.random.com/emailAddress=random@email.com"

mv certs ../