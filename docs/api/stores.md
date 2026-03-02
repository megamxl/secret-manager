# 🗄️ Stores API

The **Stores API** manages backend secret stores used by the Secret Manager  
(e.g. HashiCorp Vault, Azure Key Vault).

A store configuration defines **where** secrets are fetched from and **how** authentication is performed.

---

## 📌 Rest Endpoint

    /store

---


## 🧱 Data Model

### Field Description

| Field           | Type   | Required | Description                           |
|-----------------|--------|----------|---------------------------------------|
| `type`          | string | ✅        | Backend type (`hc-vault`, `az-vault`) |
| `url`           | string | ✅        | Base URL of the backend store         |
| `mountPath`     | string | ❌        | Base path inside the backend store    |
| `referenceName` | string | ✅        | Unique identifier for this store      |
| `auth`          | object | ✅        | Authentication configuration          |

---

### Store Schema (JSON)

    {
      "type": "hc-vault",
      "url": "http://vault:8200",
      "mountPath": "secret",
      "referenceName": "my-store",
      "auth": {
        "token": "my-token"
      }
    }

### Store Schema (YAML)

    type: hc-vault
    url: http://vault:8200
    mountPath: secret
    referenceName: my-store
    auth:
      token: my-token

---


### Authentication (`auth`)

#### HashiCorp Vault (`hc-vault`)

Token-based authentication:

    auth:
      token: s.xxxxx

AppRole authentication:

    auth:
      approle:
        role_id: my-role-id
        secret_id: my-secret-id

#### Azure Key Vault (`az-vault`)

    auth:
      tenant_id: <tenant-id>
      client_id: <client-id>
      client_secret: <client-secret>

---

## ➕ Create Store

**POST** `/store`

### Request (JSON)

    {
      "type": "hc-vault",
      "url": "http://vault:8200",
      "mountPath": "secret",
      "referenceName": "my-store",
      "auth": {
        "token": "s.xxxxx"
      }
    }

### Request (YAML)

    type: hc-vault
    url: http://vault:8200
    mountPath: secret
    referenceName: my-store
    auth:
      token: s.xxxxx

### Responses

| Status                      | Description               |
|-----------------------------|---------------------------|
| `201 Created`               | Store configuration saved |
| `400 Bad Request`           | Invalid configuration     |
| `500 Internal Server Error` | Backend connection failed |

---

## ✏️ Update Store

**PUT** `/store`

Request body is identical to **Create Store** (JSON or YAML).

### Responses

| Status          | Description          |
|-----------------|----------------------|
| `200 OK`        | Store updated        |
| `404 Not Found` | Store does not exist |

---

## 🗑️ Delete Store

**DELETE** `/store`

### Request (JSON)

    {
      "referenceName": "my-store"
    }

### Request (YAML)

    referenceName: my-store

### Responses

| Status           | Description     |
|------------------|-----------------|
| `204 No Content` | Store deleted   |
| `404 Not Found`  | Store not found |

---

## 📄 List Stores

**GET** `/store`

### Response

    [
      {
        "type": "hc-vault",
        "referenceName": "my-store",
        "url": "http://vault:8200",
        "mountPath": "secret"
      }
    ]

---

## ✅ Best Practices

- Use stable `referenceName` values (used by Secrets)
- Separate stores by environment (dev / staging / prod)
- Prefer AppRole or managed identity over static tokens
- Validate backend connectivity before creating secrets