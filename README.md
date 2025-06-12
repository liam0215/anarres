# anarres

## Requirements

Same as the [requirements for blueprint](https://github.com/Blueprint-uServices/blueprint/blob/main/docs/manual/requirements.md):

- gRPC
- Docker
- Kubernetes

And in addition:

- [Kompose](https://kompose.io/installation/) to convert the docker-compose file to Kubernetes files.
- [Intel QPL](https://intel.github.io/qpl/documentation/get_started_docs/installation.html) to run the QPL modules.

## Setup

```bash
make
make run
```
