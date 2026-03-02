# Configuration

The Secret Manager is configured via a YAML file. The default configuration file is `config.yaml`.

## Global Structure

```yaml
agent:
  nodeName: "my-node"       # Defaults to hostname
  trustedRoots: []          # List of trusted paths from which root folder secrets are allowed to be placed 

server:
  port: 8090

logging:
  level: "info"             # debug, info, warn, error
  format: "json"            # json or console
  auditFile: "/var/log/secret-manager/audit.log"

database:
  path: "./secrets.db"

security:
  tls:
    enabled: false
    certFile: ""
    keyFile: ""
    clientCAFile: ""        # Used for mTLS client CA verification
  
  auth:
    method: "none"          # "none", "mtls", "jwt", "spire"
    
    jwt:
      secret: ""            # HMAC shared secret
      jwksUrl: ""           # e.g. https://issuer/.well-known/jwks.json
      issuer: ""            # Validated if non-empty
      audience: ""          # Validated if non-empty
      
    spire:
      socketPath: "/tmp/spire-agent/public/api.sock"
      trustDomain: "example.org"
      allowedSpiffeIds: []  # Allowlist of SPIFFE IDs that may call the API
```

## Environment Variables

Configuration can also be set via environment variables prefixed with `SECRET_MANAGER`.
For example, `agent.workDir` can be set via `SECRET_MANAGER_AGENT_WORKDIR`.
