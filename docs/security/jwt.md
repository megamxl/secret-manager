# JWT / Keycloak

The Secret Manager supports JWT (JSON Web Token) based authentication.
This allows integration with identity providers like Keycloak.

## Configuration

To enable JWT authentication:

1.  Set `security.auth.method` to `jwt`.
2.  Configure the `security.auth.jwt` section.

```yaml
security:
  auth:
    method: "jwt"
    jwt:
      secret: ""            # HMAC shared secret (optional if using JWKS)
      jwksUrl: "https://auth.example.com/.well-known/jwks.json" # e.g. Keycloak
      issuer: "https://auth.example.com/realms/master"
      audience: "secret-manager"
```

## How it works

The Secret Manager validates the JWT signature and claims (`iss`, `aud`) provided in the `Authorization: Bearer <token>` header.
The `sub` claim is used to identify the caller in audit logs.
