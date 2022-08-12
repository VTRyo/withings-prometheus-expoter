# Withings Prometheus Exporter

Body information measured by [Withings API](https://developer.withings.com/) can be monitored by Prometheus Exporter.

# Getting started

Withings setup is a prerequisite. Please refer to this document.

**Prerequisite**: https://github.com/zono-dev/withings-go#create-your-withings-account

When you use this Prometheus Exporter, remember to mount the following two files.

- settings.yaml
- access_token.json

The library's（[withings-go](https://github.com/zono-dev/withings-go#getting-started)）README is the most detailed description of the generation.

## Generate access_token.json

Enter the `Grant Code` generated based on the information in settings.yaml.

```
map[CID:YOUR-CLIENT-ID RedirectURL:https://example.com/ Secret:YOUR-SECRET]
[user.activity,user.metrics,user.info]
URL to authorize:http://account.withings.com/oauth2_user/authorize2?access_type=offline&client_id=yourclientid&redirect_uri=https%3A%2F%2Fexample.com&response_type=code&scope=user.activity%2Cuser.metrics%2Cuser.info&state=state
Open url your browser and Enter your grant code here.
 Grant Code:
```

## settings.yaml

[](https://github.com/zono-dev/withings-go#setup-your-settings-file)

Enter your Withings application information into config.

[Withings Developer Dashboard](https://developer.withings.com/dashboard/)

```yaml
CID: "" # ClientID
Secret: ""
RedirectURL: "http://localhost:8181/callback"
```

# Usage
## Docker Compose

For example, this is the config for monitoring with Prometheus + Grafana.

```
$ tree
.
├── docker-compose.yml
├── grafana.env
├── prometheus.yml
└── withings
    ├── access_token.json
    └── settings.yaml
```

```yaml
version: "3"
services:
  prometheus:
    image: prom/prometheus
    container_name: prometheus
    volumes:
      - ./prometheus.yml:/etc/prometheus/prometheus.yml
    ports:
      - 9090:9090
  grafana:
    image: grafana/grafana
    container_name: grafana
    ports:
      - 3000:3000
    env_file:
      - ./grafana.env
  withings-prometheus-expoter:
    image: withings-prometheus-expoter:latest
    container_name: withings-prometheus-expoter
    volumes:
      # If the location where you are placing the files is different, please change it.
      - ./withings/access_token.json:/access_token.json
      - ./withings/settings.yaml:/settings.yaml
    ports:
      - 8181:8181
```

## Docker

```sh
docker run \ 
  -v $(pwd)/withings/access_token.json:/access_token.json \
  -v $(pwd)/withings/settings.yaml:/settings.yaml \
  withings-prometheus-expoter:latest sh
```
