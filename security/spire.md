# SPIRE Integration

The Secret Manager can integrate with SPIRE for SPIFFE-based authentication.
This is useful for securing communication between microservices.

## Configuration

To enable SPIRE authentication:

1.  Set `security.auth.method` to `spire`.
2.  Configure the `security.auth.spire` section.

```yaml
security:
  auth:
    method: "spire"
    spire:
      socketPath: "/tmp/spire-agent/public/api.sock"
      trustDomain: "example.org"
      allowedSpiffeIds: 
        - "spiffe://example.org/ns/default/sa/my-service"
```

## How it works

The Secret Manager uses the SPIRE workload API to validate the SVID presented by the client.
If `allowedSpiffeIds` is non-empty, the caller's SPIFFE ID must match one of the allowed IDs.
Otherwise, any valid SVID in the trust domain is accepted.
The SPIFFE ID is used to identify the caller in audit logs.
