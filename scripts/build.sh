if ! buf lint; then
  echo "buf lint failed"
  exit 1
fi

npx buf generate --template ./proto/server.gen.yaml --path ./proto/gateway
npx buf generate --template ./proto/client.gen.yaml --path ./proto/gateway
