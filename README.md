# Introduction

The Secret Manager is a Go-based agent designed to manage secrets in any environment, similar to the External Secrets Operator but for non-Kubernetes environments.

It runs as a daemon or a CLI tool that:
*   Fetches secrets from external stores (HashiCorp Vault, Azure Key Vault).
*   Templates configuration files with those secrets.
*   Writes the resulting files to the local filesystem with correct permissions.
*   Rotates secrets based on configured policies.

## Key Features

*   **Secret Fetching:** Retrieve secrets from multiple backends.
*   **Templating:** Inject secrets into config files.
*   **Rotation:** Automatically rotate secrets and restart services if needed.
*   **Security:**
    *   **mTLS** for secure communication.
    *   **SPIRE** integration for workload attestation.
    *   **JWT** authentication.
*   **Audit Logging:** Comprehensive audit trails for all operations.

## Status

Active development. Supports file-level secret resolution.

## Getting Started

Check out the [Installation](getting-started/installation.md) guide to get started.
