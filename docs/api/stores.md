# Stores API

The Secret Manager supports backend storage like HashiCorp Vault or Azure Key Vault.

## Base URL

`http://localhost:8090/store`

## Create Store

**Method:** `POST`

**Request Body:**

```json
{
  "type": "hc-vault",
  "url": "http://vault:8200",
  "mountPath": "secret",
  "referenceName": "my-store",
  "auth": {
    "token": "my-token"
  }
}
```

*   `type`: Type of the backend store (`hc-vault` or `az-vault`).
*   `url`: URL of the backend store.
*   `mountPath`: Base path in the backend store.
*   `referenceName`: Unique identifier for this store configuration.
*   `auth`: Authentication details for the backend.
    *   For `hc-vault`: `token` (e.g. `s.xxxxx`) or `approle` creds.

**Response:**

*   `201 Created`: Store configuration saved.
*   `400 Bad Request`: Invalid configuration.
*   `500 Internal Server Error`: Backend connection failed.

## Update Store

**Method:** `PUT`

**Request Body:** Same as `Create Store`.

**Response:** `200 OK`

## Delete Store

**Method:** `DELETE`

**Request Body:**

```json
{
  "referenceName": "my-store"
}
```

**Response:** `204 No Content`

## List Stores

**Method:** `GET`

**Response:** Returns a list of configured stores.

```json
[
  {
    "type": "hc-vault",
    "referenceName": "my-store",
    ...
  }
]
```
