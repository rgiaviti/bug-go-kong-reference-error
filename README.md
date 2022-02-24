# Go-Kong Bug 
This is a sample Go project to simulate a detected bug in go-kong v0.28.0. Issue opened [here](https://github.com/Kong/go-kong/issues/137)

## When the bug happens
When we try to validate a plugin schema using [go-kong](https://github.com/Kong/go-kong) and the target kong is offline 
(or even if the Admin API was disabled), we get a reference error because the response is `nil`.

## Steps to reproduce
### Deploy a Kong with AdminAPI Enabled
```
$ docker run -d --name kong \
    -e "KONG_DATABASE=off" \
    -e "KONG_PROXY_ACCESS_LOG=/dev/stdout" \
    -e "KONG_ADMIN_ACCESS_LOG=/dev/stdout" \
    -e "KONG_PROXY_ERROR_LOG=/dev/stderr" \
    -e "KONG_ADMIN_ERROR_LOG=/dev/stderr" \
    -e "KONG_ADMIN_LISTEN=0.0.0.0:8001, 0.0.0.0:8444 ssl" \
    -e "KONG_PLUGINS=bundled" \
    -p 8000:8000 \
    -p 8443:8443 \
    -p 8001:8001 \
    -p 8444:8444 \
    kong
```

### Check if AdminAPI is responding in Port 8001
```
curl --location --request GET 'http://localhost:8001/status'
```
Status Code must be `200 OK`

### Run the Program
```
$ go run main.go
```

### Result
One function will validate a plugin with success. The other function will crash go-kong.
```
Reproducing bug...
Calling AdminAPI in Port 8001
validation of correlation-id
result: true
message: 
Calling AdminAPI in Port 9000
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x76792c]

goroutine 1 [running]:
github.com/kong/go-kong/kong.(*PluginService).Validate(0xc0000b0040, {0x81bdb0, 0xc0000240b8}, 0xc0000b4000)
        /home/rgiaviti/go/pkg/mod/github.com/kong/go-kong@v0.28.1/kong/plugin_service.go:186 +0x26c
main.CallingIncorrectKongPort()
        /home/rgiaviti/Desktop/bug-go-kong-reference-error/main.go:41 +0x123
main.main()
        /home/rgiaviti/Desktop/bug-go-kong-reference-error/main.go:12 +0x70
```
