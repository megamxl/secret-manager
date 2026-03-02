# Installation

Currently, the Secret Manager is built from source.

## Prerequisites

*   Go 1.21 or higher
*   Git

## Build

1.  Clone the repository:
    ```bash
    git clone https://github.com/your-org/secret-manager.git
    cd secret-manager
    ```

2.  Build the binary:
    ```bash
    go build -o secret-manager ./cmd/agent
    ```

3.  Run the agent:
    ```bash
    ./secret-manager --config config.yaml
    ```
