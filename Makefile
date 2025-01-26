docker_compose_bin := $(shell command -v docker-compose 2> /dev/null)

up:
	$(docker_compose_bin) up -d

down:
	$(docker_compose_bin) down -v

file:
	curl -X PUT \
    -H "Content-Type: application/json" \
    --data @connectors/file-connector.json \
    http://localhost:8083/connectors/file-stream-sink/config

filejson:
	curl -X PUT \
    -H "Content-Type: application/json" \
    --data @connectors/file-json-connector.json \
    http://localhost:8083/connectors/file-stream-sink-json/config

jdbc:
	curl -X PUT \
    -H "Content-Type: application/json" \
    --data @connectors/jdbc.json \
    http://localhost:8083/connectors/jdbc-sink/config

jdbc-status:
	curl localhost:8083/connectors/jdbc-sink/status | jq