# Web/Sockets Example

Web server gwith web socket middleware and TLS support.

[Web Server](https://github.com/gin-gonic)
[Web Sockets](github.com/gorilla/websocket)

## Local Development

### TLS Setup

If you want to HTTPS locally you can generate a cert for localhost/custom host name using mkcert:

https://github.com/FiloSottile/mkcert

Use mkcert to install a local CA (certificate authority) and generate a cert+key.

Example using `dev.local` as a local development url:

Update hosts file:

```txt
127.0.0.1 dev.local
```

```bash
mkcert -install
mkcert wsl.local
```

Update configs/local.json TLSKeyFile and TLSCertFile paths to point at the generated files