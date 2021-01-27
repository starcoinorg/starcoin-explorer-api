## docker starcoin-explorer-api

1. stop container，delete old container

`./stop.sh`

2. build docker image

`./build_image.sh`

3. start container

`./run.sh`

4. check log

`docker logs -f starcoin-explorer-api`

5. One-click for all above

`./rebuild.sh`

6. inspect a running container.
`docker exec -it <CONTAINER_ID> /bin/bash`

## Publish docker image to the hub

`docker push starcoin-explorer-api:latest`
