# 🔐 Secrets API

The **Secret Manager API** is a RESTful interface for creating, updating, deleting, and listing secrets.  
It supports **JSON** and **YAML** request bodies and uses **Go templates** to render secrets.

---

## 📌 Rest Endpoint
    /secret
---

## 🧱 Data Model

### Field Description

| Field                              | Type   | Description                                           |
|------------------------------------|--------|-------------------------------------------------------|
| `file_path`                        | string | Path where the rendered secret is written on the node |
| `name`                             | string | Unique name of the secret configuration               |
| `store_name`                       | string | Backend secret store reference                        |
| `reroll_time`                      | string | Rotation interval (`1h`, `30m`, etc.)                 |
| `secret_wrapper`                   | object | Secret rendering configuration                        |
| `secret_wrapper.content`           | string | Go template used to render the secret                 |
| `secret_wrapper.secret_references` | map    | Template variable → backend secret path               |

---

### Secret Schema (JSON)

    {
      "file_path": "/path/to/secret",
      "name": "my-secret",
      "store_name": "my-store",
      "reroll_time": "1h",
      "secret_wrapper": {
        "content": "raw-secret-content",
        "secret_references": {
          "key": "backend/secret/path"
        }
      }
    }

### Secret Schema (YAML)

    file_path: /path/to/secret
    name: my-secret
    store_name: my-store
    reroll_time: 1h
    secret_wrapper:
      content: raw-secret-content
      secret_references:
        key: backend/secret/path

---


## 🧠 Template & Reference Matching

The `content` field uses **Go template syntax**.  
Keys defined in `secret_references` become template variables.

### Example (YAML)

    file_path: /etc/app/config.yaml
    name: database-config
    store_name: vault-main
    reroll_time: 1h
    secret_wrapper:
      content: |
        username: {{ .db_user }}
        password: {{ .db_password }}
      secret_references:
        db_user: secret/data/db/username
        db_password: secret/data/db/password

✔ The variable names in `content` **must match** the keys in `secret_references`.

Rendered file:

    username: admin
    password: s3cr3t

---

## ➕ Create Secret

**POST** `/secret`

### Request (JSON)

    {
      "file_path": "/etc/app/secret.env",
      "name": "app-secret",
      "store_name": "vault-main",
      "reroll_time": "30m",
      "secret_wrapper": {
        "content": "API_KEY={{ .api_key }}",
        "secret_references": {
          "api_key": "secret/data/app/api_key"
        }
      }
    }

### Request (YAML)

    file_path: /etc/app/secret.env
    name: app-secret
    store_name: vault-main
    reroll_time: 30m
    secret_wrapper:
      content: API_KEY={{ .api_key }}
      secret_references:
        api_key: secret/data/app/api_key

### Responses

| Status | Description |
|------|-------------|
| `201 Created` | Secret created |
| `400 Bad Request` | Invalid request |
| `500 Internal Server Error` | Server error |

---

## ✏️ Update Secret

**PUT** `/secret`

Request body is identical to **Create Secret** (JSON or YAML).

### Responses

| Status | Description |
|------|-------------|
| `200 OK` | Secret updated |
| `404 Not Found` | Secret does not exist |

---

## 🗑️ Delete Secret

**DELETE** `/secret`

### Request (JSON)

    {
      "name": "app-secret"
    }

### Request (YAML)

    name: app-secret

### Responses

| Status | Description |
|------|-------------|
| `204 No Content` | Secret deleted |
| `404 Not Found` | Secret not found |

---

## 📄 List Secrets

**GET** `/secret`

### Response

    [
      {
        "name": "app-secret",
        "file_path": "/etc/app/secret.env",
        "store_name": "vault-main",
        "reroll_time": "30m"
      }
    ]

---

## ✅ Best Practices

- Prefer **YAML** for GitOps workflows
- Keep `name` stable — it is the primary identifier
- Ensure template variables match `secret_references`
- Use `reroll_time` for automatic rotation