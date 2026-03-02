# Mutual TLS (mTLS)

mTLS authentication ensures that both the client and the server are authenticated using certificates.

## Configuration

To enable mTLS authentication:

1.  Enable TLS (`security.tls.enabled: true`).
2.  Set `security.auth.method` to `mtls`.
3.  Set `security.tls.clientCAFile` to the path of the CA certificate that signed the client certificates.

```yaml
security:
  tls:
    enabled: true
    certFile: "/path/to/server.crt"
    keyFile: "/path/to/server.key"
    clientCAFile: "/path/to/client-ca.crt" # Required for mTLS
  auth:
    method: "mtls"
```

## How it works

When mTLS is enabled, the Secret Manager will verify that the client presents a valid certificate signed by the CA specified in `clientCAFile`.
The subject of the client certificate is used to identify the caller in audit logs.
