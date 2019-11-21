# Status Check


## Instalation


### Install from source

```
go build -o status main.go
```

### Run docker image

```
docker run -itd -p 8008:8008 .
```

### Build your own image

1. Build this Dockerimage
```
docker build -t kununu/statuscheck .
```
2. Create your `config.yaml` file

3. Create your final image
```
FROM kununu/statuscheck

COPY config.yaml /sc/config.yaml

CMD [ "/sc/status", "-c", "/sc/config.yaml" ]
```

## Configuration

Create a `config.yaml` file on the system.
If you want to place it anywhere on the system's path, just pass the `-c path/to/config.yaml` parameter on the startup.

```
port: 8008
endpoint: /status
okMessage: "ok"
errorMessage: "fail"
```
If no configuration file is detected, the default values will be applied.


| Option       | Type   | Description                                               |
|--------------|--------|-----------------------------------------------------------|
| port         | string | The port to listen on. (default: 8008)                    |
| endpoint     | string | The endpoint to listen on. (default: /status)             | 
| okMessage    | string | The message to return in case of success. (default: ok)   |
| errorMessage | string | The message to return in case of failure. (default: fail) |

## Types of status check 

### HTTP (default)

The HTTP status check allows you to check for an HTTP endpoint: 

```
port: 8008
endpoint: /status
okMessage: "ok"
errorMessage: "fail"
check:
    type: "http"
    url: "https://www.kununu.com"
    followRedirects: false
    statusCode: 200
```

| Option          | Type   | Description                                               |
|-----------------|--------|-----------------------------------------------------------|
| type            | string | The HTTP type sets a request to an HTTP endpoint.         |
| url             | string | The url to make the HTTP request to. (default: localhost) | 
| followRedirects | bool   | Sets if it should follow HTTP redirects. (default: false) |
| statusCode      | int    | The status code to compare to. (default: 200)             |


