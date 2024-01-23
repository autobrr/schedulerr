# Schedulerr

## How to use

Use with our without a config file. See example config [here.](./config.yaml)

`schedulerr --config config.yaml`

If hour block is enabled, it will return status 200.
Returns 403 if hour block is disabled or missing.

Send the request with or without a payload. If a config file is loaded, then that will be used unless the request comes with a payload. A payload will always take precedence over the config file.

```bash
curl -v -X POST http://localhost:8585/webhook
```

```bash
curl -X POST -H "Content-Type: application/json" \
-d '{
    "sunday": [
        { "hour": 0, "enabled": false },
        { "hour": 1, "enabled": true },
        { "hour": 21, "enabled": false }
    ],
    "monday": [
        { "hour": 0, "enabled": false },
        { "hour": 1, "enabled": true },
        { "hour": 21, "enabled": false }
    ],
    "tuesday": [
        { "hour": 0, "enabled": false },
        { "hour": 1, "enabled": true },
        { "hour": 21, "enabled": false }
    ],
    "wednesday": [
        { "hour": 0, "enabled": false },
        { "hour": 1, "enabled": true },
        { "hour": 21, "enabled": false }
    ],
    "thursday": [
        { "hour": 0, "enabled": false },
        { "hour": 1, "enabled": true },
        { "hour": 21, "enabled": false }
    ],
    "friday": [
        { "hour": 0, "enabled": false },
        { "hour": 1, "enabled": true },
        { "hour": 21, "enabled": false }
    ],
    "saturday": [
        { "hour": 0, "enabled": false },
        { "hour": 1, "enabled": true },
        { "hour": 21, "enabled": false }
    ]
}' http://localhost:8585/webhook
```
