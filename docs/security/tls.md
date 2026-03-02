# Transport Layer Security (TLS)

The Secret Manager can be configured to serve HTTPS traffic.

## Configuration

To enable TLS, configure the `security.tls` section in your `config.yaml`:

```yaml
security:
  tls:
    enabled: true
    certFile: "/path/to/server.crt"
    keyFile: "/path/to/server.key"
```

If `enabled` is set to `true`, the `certFile` and `keyFile` must be provided.
This ensures that all communication with the Secret Manager API is encrypted.
