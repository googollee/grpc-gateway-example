#!/bin/sh

curl "http://localhost:8080/v1/service/greet?name=world&ticker_second=1" -H "Authorization: basic token"
