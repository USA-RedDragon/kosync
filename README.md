# Kosync

Kosync is a re-implementation of <https://github.com/koreader/koreader-sync-server> in Go.

## Configuration

<!-- configulator:begin -->

| Key                       | Type           | Default  | Environment               | Flag                        | Description                                                           |
|---------------------------|----------------|----------|---------------------------|-----------------------------|-----------------------------------------------------------------------|
| `log-level`               | string         | `info`   | `LOG_LEVEL`               | `--log-level`               | Logging level for the application. One of debug, info, warn, or error |
| `storage.type`            | string         | `sqlite` | `STORAGE_TYPE`            | `--storage.type`            | Storage type. One of mysql, postgres, sqlite                          |
| `storage.dsn`             | string         |          | `STORAGE_DSN`             | `--storage.dsn`             | Data source name for the storage                                      |
| `auth.salt`               | string         |          | `AUTH_SALT`               | `--auth.salt`               | Salt for hashing passwords (required) (secret)                        |
| `auth.allow-registration` | boolean        | `true`   | `AUTH_ALLOW_REGISTRATION` | `--auth.allow-registration` | Allow user registration                                               |
| `http.address`            | string         |          | `HTTP_ADDRESS`            | `--http.address`            | Address to listen on                                                  |
| `http.port`               | integer        | `8080`   | `HTTP_PORT`               | `--http.port`               | Port to listen on                                                     |
| `http.trusted-proxies`    | list of string |          | `HTTP_TRUSTED_PROXIES`    | `--http.trusted-proxies`    | Trusted proxies for the HTTP server                                   |
| `metrics.enabled`         | boolean        |          | `METRICS_ENABLED`         | `--metrics.enabled`         | Enable metrics server                                                 |
| `metrics.address`         | string         |          | `METRICS_ADDRESS`         | `--metrics.address`         | Address to listen on                                                  |
| `metrics.port`            | integer        | `9000`   | `METRICS_PORT`            | `--metrics.port`            | Port to listen on                                                     |
| `pprof.enabled`           | boolean        |          | `PPROF_ENABLED`           | `--pprof.enabled`           | Enable pprof server                                                   |
| `pprof.address`           | string         |          | `PPROF_ADDRESS`           | `--pprof.address`           | Address to listen on                                                  |
| `pprof.port`              | integer        | `9999`   | `PPROF_PORT`              | `--pprof.port`              | Port to listen on                                                     |

<!-- configulator:end -->
