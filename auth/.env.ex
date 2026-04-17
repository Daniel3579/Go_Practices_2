AUTH_PORT=:443                  #For TLS
DATABASE_URL=postgresql://<NAME>:<PASSWORD>@host.docker.internal:<PORT>/<DB_NAME>?sslmode=disable
SECRET_KEY=secret_key
AUTH_METRICS_PORT=:9090
AUTH_CERT_FILE=certs/server.crt #Path to server.crt
AUTH_KEY_FILE=certs/server.key  #Path to server.key