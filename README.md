# Schedulerr

## How to use

If hour block is not provided, it will return status 200.
Returns 403 if hour block is disabled.

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
