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
docker build -t enkodr/statuscheck .
```
2. Create your `config.yaml` file

3. Create your final image
```
FROM enkodr/statuscheck

COPY config.yaml /sc/config.yaml

CMD [ "/sc/status", "-c", "/sc/config.yaml" ]
```

## Configuration

Create a `config.yaml` file on the same path as the status binary file.

If you want to place it anywhere on the system's path, just pass the `-c path/to/config.yaml` parameter on the startup.

```
port: 8008
endpoint: /status
json: false
okMessage: "ok"
errorMessage: "fail"
```
NOTE: If no configuration file is detected, the default values will be applied.


| Option       | Type   | Default | Description                                                                            |
|--------------|--------|---------|----------------------------------------------------------------------------------------|
| port         | string | 8008    | The port to listen on.                                                                 |
| endpoint     | string | /status | The endpoint to listen on.                                                             | 
| json         | bool   | false   | Returns the response in JSON. If true, the endpoint name will be used as the JSON key. |
| okMessage    | string | ok      | The message to return in case of success.                                              |
| errorMessage | string | fail    | The message to return in case of failure.                                              |

## Types of status check 

### HTTP (default)

The HTTP status check allows you to check for an HTTP endpoint: 

```
port: 8008
endpoint: /status
json: true
okMessage: "ok"
errorMessage: "fail"
check:
    type: "http"
    url: "https://www.kununu.com"
    followRedirects: false
    statusCode: 200
```

| Option          | Type   | Default          | Description                                       |
|-----------------|--------|------------------|---------------------------------------------------|
| type            | string | http             | The HTTP type sets a request to an HTTP endpoint. |
| url             | string | http://localhost | The url to make the HTTP request to.              | 
| followRedirects | bool   | false            | Sets if it should follow HTTP redirects.          |
| statusCode      | int    | 200              | The status code to compare to.                    |

### Port

The Port status check verifies if a specific port is responding:

```
port: 8008
endpoint: /status
json: true
okMessage: "ok"
errorMessage: "fail"
check:
    type: "port"
    address: "127.0.0.1"
    port: 80
    protocol: "tcp"
```

| Option   | Type   | Default   | Description                                       |
|----------|--------|-----------|---------------------------------------------------|
| type     | string | http      | The HTTP type sets a request to an HTTP endpoint. |
| address  | string | 127.0.0.1 | The address to check.                             |
| port     | int    | 80        | The port to check.                                |
| protocol | string | tcp       | The protocol to use.                              |


# TODO

- Check for specific services connection
- Check for file presence
- Check for ready service
