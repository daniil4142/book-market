#!/bin/sh

GRPC_HOST="localhost:8002"
GRPC_METHOD="daniil4142.book_market.book_service.v1.BookService/CreateBook"

payload=$(
  cat <<EOF
{
  "categoryId": "3",
  "name": "Hello"
}
EOF
)

grpcurl -plaintext -emit-defaults \
  -d "${payload}" ${GRPC_HOST} ${GRPC_METHOD}