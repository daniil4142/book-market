#!/bin/sh

GRPC_HOST="localhost:7002"
GRPC_METHOD="daniil4142.book_market.category_service.v1.CategoryService/GetCategoryById"

payload=$(
  cat <<EOF
{
  "id": 3
}
EOF
)

grpcurl -plaintext -emit-defaults \
  -rpc-header 'x-app-name:dev' \
  -rpc-header 'x-app-version:1' \
  -d "${payload}" ${GRPC_HOST} ${GRPC_METHOD}