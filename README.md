# Schedulerr

Schedulerr is a webhook companion for autobrr. It's made for those who wants an easy way to stop racing during certain hours throughout the day/week by rejecting releases within the hours set to false in the config/payload.

## How to use

Use with our without a config file. See example config [here.](config/config.yaml)
Using the docker image requires a config file unless you build it yourself.

`schedulerr --config /config/config.yaml`

- If hour block is set to true or missing, it will return status 200.  
- If hour block is set to false it will return 403.

Send the request with or without a payload. If a config file is loaded, then that will be used unless the request comes with a payload for the relevant hour block. Hour blocks in a payload will always take precedence over the hour blocks in the config file.

### Setup in autobrr

In autobrr you would set this up as an external webhook in the filters you want to limit.

Go to the `External` tab -> Add new `External Filter` of type `Webhook`.

1. Endpoint: `http://schedulerr:8585/webhook`
2. HTTP Method: `POST`
3. Expected HTTP Status Code: `200`
4. Payload (optional when using a config file):

```json
{
    "sunday": [
        { "hour": 0, "enabled": false },
        { "hour": 1, "enabled": false },
        { "hour": 23, "enabled": false }
    ],
    "friday": [
        { "hour": 0, "enabled": false },
        { "hour": 23, "enabled": false }
    ],
    "saturday": [
        { "hour": 0, "enabled": false },
        { "hour": 23, "enabled": false }
    ]
}
```
