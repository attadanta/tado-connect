# Tado Connect

## The Tado API

### Authentication

```
POST https://auth.tado.com/oauth/token
```

Content-Type: `application/x-www-form-urlencoded`

| Parameter     | Type   | Description        |
| ------------- | ------ | ------------------ |
| `client_id`   | string | `tado-web-app`     |
| `scope`       | string | `home.user`        |
| grant_type    | string | `password`         |
| username      | string | Tado username      |
| password      | string | Tado password      |
| client_secret | string | Tado client secret |

