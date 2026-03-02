# Secrets API

The Secret Manager provides a RESTful API for managing secrets.

## Base URL

`http://localhost:8090/secret`

## Create Secret

**Method:** `POST`

**Request Body:**

```json
{
  "file_path": "/path/to/secret",
  "name": "my-secret",
  "store_name": "my-store",
  "reroll_time": "1h",
  "secret_wrapper": {
    "content": "raw-secret-content",
    "secret_references": {
      "key": "value"
    }
  }
}
```

*   `file_path`: Path where the secret will be written on the node.
*   `name`: Unique name of the secret config.
*   `store_name`: Reference to the `store` where the secret is fetched from.
*   `reroll_time`: Duration string (e.g., "1h", "30m").
*   `secret_wrapper`:
    *   `content`: Raw content (if not using references). In this content use go tmpl language to build your secret 
    *   `secret_references`: Map of keys to secret paths in the backend store.

**Response:**

*   `201 Created`: Secret created successfully.
*   `400 Bad Request`: Invalid request body.
*   `500 Internal Server Error`: Server error.

## Update Secret

**Method:** `PUT`

**Request Body:** Same as `Create Secret`.

**Response:** `200 OK`

## Delete Secret

**Method:** `DELETE`

**Request Body:**

```json
{
  "name": "my-secret"
}
```

**Response:** `204 No Content`

## List Secrets

**Method:** `GET`

**Response:** Returns a list of all configured secrets.

```json
[
  {
    "file_path": "/path/to/secret",
    "name": "my-secret",
    ...
  }
]
```
