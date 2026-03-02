# Security Overview

The Secret Manager supports multiple authentication methods to secure access to the API.

## Authentication Methods

The following methods are supported:

*   **None (Default):** No authentication. This is intended for development only.
*   **mTLS:** Mutual TLS authentication.
*   **JWT:** JSON Web Token authentication (compatible with Keycloak).
*   **SPIRE:** SPIFFE Workload Attestation.

Configuration is handled in the `security.auth` section of the `config.yaml` file.
