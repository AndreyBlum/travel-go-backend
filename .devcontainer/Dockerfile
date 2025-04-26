FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
# [Optional] Uncomment this section to install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends postgresql-client

# [Optional] Uncomment the next lines to use go get to install anything else you need
# RUN go get -x <your-dependency-or-tool>
# USER root
