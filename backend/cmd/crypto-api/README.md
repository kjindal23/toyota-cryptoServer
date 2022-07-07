# crypto API

## Server

### Run locally

-   crpto-api > `LOCAL=true go run main.go`


### Generate

-   crypto-api > `go-service-doc -d docs/service -o docs/service -p /docs/service`

### View locally

-   http://localhost:8080/docs/service

### View locally

-   http://localhost:8080/docs/swagger

## Datadog

### Run locally

```
DOCKER_CONTENT_TRUST=1 \
docker run -d -v /var/run/docker.sock:/var/run/docker.sock:ro \
              -v /proc/:/host/proc/:ro \
              -v /sys/fs/cgroup/:/host/sys/fs/cgroup:ro \
              -e DD_API_KEY=<your_api_key> \
              -e DD_APM_ENABLED=True \
              datadog/agent:latest
```
