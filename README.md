# Status Check

## Instalation

## Configuration

Create a `config.yaml` file on the system.
If you want to place it anywhere on the system's path, just pass the `-c path/to/config.yaml` parameter on the startup.

```
port: 8008
endpoint: /status
okMessage: "ok"
errorMessage: "fail"
```

| Option       | Description                                               |
|--------------|-----------------------------------------------------------|
| port         | The port to listen on. (default: 8008)                    |
| endpoint     | The endpoint to listen on. (default: /status)             | 
| okMessage    | The message to return in case of success. (default: ok)   |
| errorMessage | The message to return in case of failure. (default: fail) |

## Types of status check 

### HTTP (default)

The HTTP health check allows you to check for an HTTP endpoint: 

```
port: 8008
endpoint: /status
okMessage: "ok"
errorMessage: "fail"
check:
    type: "http""
    url: "https://www.kununu.com"
    followRedirects: false
    statusCode: 200
```

| Option          | Description                                               |
|-----------------|-----------------------------------------------------------|
| type            | The HTTP type sets a request to an HTTP endpoint.         |
| url             | The url to make the HTTP request to. (default: localhost) | 
| followRedirects | Sets if it should follow HTTP redirects. (default: false) |
| statusCode      | The status code to compare to. (default: 200)             |


