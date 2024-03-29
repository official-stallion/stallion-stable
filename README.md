# Stable

![GitHub top language](https://img.shields.io/github/languages/top/official-stallion/stable)
![GitHub release (with filter)](https://img.shields.io/github/v/release/official-stallion/stable)

Release ```Stallion``` docker images based on versions.
You can create a ```Stallion``` server on docker from the images below.

## ENV

```Stallion``` cluster environmental variables are as follow:

- ```ST_SERVER_PORT``` : cluster server port
- ```ST_METRICS_PORT``` : cluster metrics server port (available on ```/metrics```)
- ```ST_USER``` : cluster root user
- ```ST_PASSWORD``` : cluster root password

## Image

```shell
amirhossein21/stallion:v1.4.1
```

### Docker run
To run image on a single container:

```shell
docker run -p 7025:7025 -d stallion amirhossein21/stallion:v1.4.1
```

### Docker compose
Docker compose for stallion server:

```yaml
version: "3.9"
services:
  stallion-server:
    image: amirhossein21/stallion:v1.4.1
    ports:
      - "7025:7025"
    environment:
      ST_SERVER_PORT: 7025
      ST_METRICS_PORT: 7026
```

## Docker hub

You can check docker hub [repository](https://hub.docker.com/repository/docker/amirhossein21/stallion) for more information and image lists.
