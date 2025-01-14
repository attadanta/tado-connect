# Tado Connect

## The Tado API

This is a toy project to connect to the Tado API and get the current state of the zones in a home.

### Authentication

```
POST https://auth.tado.com/oauth/token
```

Content-Type: `application/x-www-form-urlencoded`

| Parameter       | Type   | Description        |
| --------------- | ------ | ------------------ |
| `client_id`     | string | `tado-web-app`     |
| `scope`         | string | `home.user`        |
| `grant_type`    | string | `password`         |
| `username`      | string | Tado username      |
| `password`      | string | Tado password      |
| `client_secret` | string | Tado client secret |


### Get Home ID

```
GET https://my.tado.com/api/v2/me
```

### Get Zone States

```
GET https://my.tado.com/api/v2/homes/{home_id}/zoneStates
```

