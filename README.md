# Confiar

Confiar is a tool to generate and manage self-signed certificate as if they are trusted by a usual certificate authority. 

This was built to assist provisioning (virtual) machines in restricted environments.

**HEADS UP**: You should really consider using real certificates or robust certificate management (such as [Vault](https://vaultproject.io)).
In the event that none of the above is applicable, let's do this together painlessly!

## Direction

### Goals

- Reduce friction to manage (self-signed) certificates in restricted environments (partial or no internet access).

### Non-goals

- Replace existing cryptographic tools (e.g. OpenSSL, BoringSSL)
- Manage REAL certificates (security concerns)
- To be used in a public environment

## Usage

### Create a self-signed certificate

The output certificate signed itself as certificate authority.
Complete specification of the certificate is availabe through `confiar generate --help`.

```sh
❯ confiar generate --fqdn myserver.corp
```

The command above will generate `cert.pem` and `key.pem` in the current working directory.
The `cert.pem` will have `myserver.corp` in `Subject Alternative Name` as `DNS` entry.
IP address can also be specified with `--ip` flag.
Both `--fqdn` and `--ip` accepts multiple entries as comma-separated list.

### Install a self-signed certificate

**HEADS UP**: some targets may require `sudo` privileges.

While most applications will rely on underlying operating system's trusted certificate authorities, some applications also allow specific certificate authorities to be trusted manually.
One example of a supported application is Docker.

```sh
❯ confiar install --target docker --from cert.pem
```

The command above will install certificate specified by `--from` as a trusted certificate authority to Docker, which allows `docker (pull|push)` operations to work smoothly.
Docker requires every certificate to be placed according to their used hostname and Confiar automatically handles that by parsing the `Subject Alternative Name` field in the provided certificate.

## Design principles

### Optional dependencies

Confiar currently only supports its own as a cryptographer to generate certificates, but [the interface in place allows substitution](internal/cryptographer) and in the future, users can use `--cryptographer` flag to specify other providers, such as OpenSSL, BoringSSL, etc.

Such pattern will persist throughout the development of Confiar, where built-ins will be the first supported provider.

### Integrates to modern infrastructure

While Confiar strives towards zero hard dependencies at runtime, the inverse is applied towards the output.
Confiar aims to support integration with any application / platform / operating system, particularly in installing certificates.

## Contributing

For any feature request / proposal, please start with opening issues.
Opening PRs without issues / prior discussion is strongly discouraged.

Be excellent to each other!

## Towards v1.0.0

The following list will _eventually_ be converted to issues and projects, though if you have thoughts before they were converted, feel free to open one and discuss!

- Support `--cryptographer` variants
  - Required: OpenSSL
  - Optionally: LibreSSL, BoringSSL, cfssl
- Support `--target` variants
  - Required: Ubuntu
  - Optionally: Any Linux distribution, maybe macOS
- Support `--from` remote (and therefore figure out a way to serve the generated certificate)
