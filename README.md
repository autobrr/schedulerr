# Schedulerr

Schedulerr is a webhook companion for autobrr. It's made for those who wants an easy way to stop racing during certain hours throughout the day/week by rejecting releases within the hours set to false in the config/payload.

## Configuration

Schedulerr will look for a configuration file in the following locations (in order):

1. Path specified by `--config` flag
2. `$HOME/.config/schedulerr/config.yaml`
3. `./config.yaml` (current directory)

### Docker Usage

Example setup:

```yaml
services:
  schedulerr:
    image: ghcr.io/autobrr/schedulerr:latest
    volumes:
      - "/path/to/docker/data/schedulerr:/config"
```

Place your `config.yaml` in the schedulerr directory. See example config [here.](config.yaml)

## How to use

Start the server:

```bash
schedulerr serve
```

The service will automatically look for a config file in the standard locations. You can also specify a custom config location:

```bash
schedulerr serve --config /path/to/config.yaml
```

- If hour block is set to true or missing, it will return status 200.
- If hour block is set to false it will return 403.

Send the request with or without a payload. If a config file is loaded, then that will be used unless the request comes with a payload for the relevant hour block. Hour blocks in a payload will always take precedence over the hour blocks in the config file.

### Setup in autobrr

In autobrr you would set this up as an external webhook in the filters you want to limit.

Go to the `External` tab -> Add new `External Filter` of type `Webhook`.

1. Endpoint: `http://schedulerr:8585/webhook`
2. HTTP Method: `POST`
3. Expected HTTP Status Code: `200`
4. Payload (not needed when using a config file, but could be used to override the config file for a specific filter etc.):

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

### Curl example for testing

```bash
curl -v -X POST http://localhost:8585/webhook \
-H "Content-Type: application/json" \
-d '{
  "sunday": [
    { "hour": 0, "enabled": false },
    { "hour": 1, "enabled": false },
    { "hour": 23, "enabled": false }
  ],
  "monday": [
    { "hour": 9, "enabled": false }
  ],
  "friday": [
    { "hour": 0, "enabled": false },
    { "hour": 23, "enabled": false }
  ],
  "saturday": [
    { "hour": 0, "enabled": false },
    { "hour": 23, "enabled": false }
  ]
}'
```
