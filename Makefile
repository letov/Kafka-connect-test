docker_compose_bin := $(shell command -v docker-compose 2> /dev/null)

up:
	$(docker_compose_bin) up -d

down:
	$(docker_compose_bin) down -v

plugins:
	curl localhost:8083/connector-plugins | jq

file:
	curl -X PUT \
    -H "Content-Type: application/json" \
    --data @infra/connectors/file.json \
    http://localhost:8083/connectors/file-sink/config

filejson:
	curl -X PUT \
    -H "Content-Type: application/json" \
    --data @infra/connectors/file-json.json \
    http://localhost:8083/connectors/file-json-sink/config

jdbc:
	curl -X PUT \
    -H "Content-Type: application/json" \
    --data @infra/connectors/jdbc.json \
    http://localhost:8083/connectors/jdbc-source/config

debezium:
	curl -X PUT \
    -H "Content-Type: application/json" \
    --data @infra/connectors/debezium.json \
    http://localhost:8083/connectors/debezium-source/config

debezium-status:
	curl localhost:8083/connectors/debezium-source/status | jq