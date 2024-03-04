if ! buf lint; then
  echo "buf lint failed"
  exit 1
fi

buf generate --template ./server.gen.yaml
npx buf generate --template ./client.gen.yaml
