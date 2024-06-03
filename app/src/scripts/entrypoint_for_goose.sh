#!/bin/sh

/usr/local/bin/wait
/usr/local/bin/goose \
  -dir /migrations \
  mysql "${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" \
  $@
