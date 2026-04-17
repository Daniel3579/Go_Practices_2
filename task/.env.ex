TASK_PORT=:443                  #For TLS
AUTH_SERVER=host.docker.internal:<External_AUTH_SERVER_PORT>
DATABASE_URL=postgresql://<NAME>:<PASSWORD>@host.docker.internal:<PORT>/<DB_NAME>?sslmode=disable
TASK_CERT_FILE=certs/server.crt #Path to server.crt
TASK_KEY_FILE=certs/server.key  #Path to server.key
CA_CERT_FILE=certs/ca.crt       #Path to ca.crt