# MC Hello World Go App

A basic Go hello world application built on docker

## Debug locally

```
# create local docker register
docker run -d -p 5000:5000 --restart=always --name registry registry:2
# build the docker image with tag
docker build -t localhost:5000/mc-hello .
# push to local registry
docker push localhost:5000/mc-hello
# run the docker image forwarding the ports
docker run -p 8080:8080 localhost:5000/mc-hello
```

Connect to the app on `http://localhost:8080`

## Endpoints

### /

| Property      | Detail
| --------------| --------------
| Description   | Hello world
| URL           | `/`
| Method        | GET
| Parameters    | None
| Returns       | 200

### /healthcheck

| Property      | Detail
| --------------| --------------
| Description   | Checks health of the system
| URL           | `/healthcheck`
| Method        | GET
| Parameters    | None
| Returns       | 200 HTTP response code if healthy

### /readiness

| Property      | Detail
| --------------| --------------
| Description   | Checks readiness of the system
| URL           | `/readiness`
| Method        | GET
| Parameters    | None
| Returns       | 200 HTTP response code if ready

