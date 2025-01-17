# Envelope Zero backend

Envelope Zero is fundamentally rooted in two ideas:

- Using the [envelope method](https://en.wikipedia.org/wiki/Envelope_system) to budget expenses into envelopes.
- Zero Based Budeting, meaning that you assign all your money to an envelope. Saving for a vacation? Create an envelope and archive it after your vacation. Rent? Create an envelope that gets a fixed amount of money added every month.

## Usage

The recommended and only supported way for production deployments is to run the backend with [the OCI image](https://github.com/envelope-zero/backend/pkgs/container/backend).

### Deployment methods

If you want to deploy with a method not listed here, you are welcome to open a discussion to ask any questions needed so that this documentation can be improved.

#### On Kubernetes

You can run the backend on any Kubernetes cluster with a supported version using the [morremeyer/generic]() helm chart with the following values:

```yaml
image:
  repository: ghcr.io/envelope-zero/backend
  tag: v0.2.1

# Only set this when you want to use sqlite as database backend.
# In this case, you need to make sure the database is regularly backed up!
persistence:
  enabled: true
  mountPath: /app/data

ports:
  - name: http
    containerPort: 8080
    protocol: TCP

ingress:
  enabled: true
  hosts:
    - host: envelope-zero.example.com
      paths:
        - path: /api
  tls:
    - hosts:
        - envelope-zero.example.com
```

### Database backends

Envelope Zero currently supports sqlite and postgresql as database backends. While sqlite is supported, it is highly recommended to use postgresql for production purposes.

#### sqlite

No configuration is needed, but you need to mount a persistent volume to the `/app/data` directory or the backend won’t start.

#### postgresql

Set the following environment variables:

- `DB_HOST`: The hostname for the postgresql instance
- `DB_USER`: The username to connect with
- `DB_PASSWORD`: The password for `DB_USER`
- `DB_NAME`: The name of the database to use

## Supported Versions

This project is under heavy development. Therefore, only the latest release is supported.

Please check the [releases page](https://github.com/envelope-zero/backend/releases) for the latest release.

## Contributing

Please see [the contribution guidelines](CONTRIBUTING.md).
