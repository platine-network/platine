###############################################################
###        STAGE 1: Build node binary pre-requisites        ###
###############################################################

FROM golang:1.19.3-alpine3.16 AS builder

ENV APPNAME=platined

# Install minimum necessary dependencies
ENV PACKAGES ca-certificates curl make git libc-dev bash gcc linux-headers file
RUN apk add --no-cache $PACKAGES

# Set working directory for the build
WORKDIR /code

COPY ../go.mod /code/
COPY ../go.sum /code/

# Download dependencies
RUN go mod download

# Add source files
COPY .. /code/

# Make node binary
RUN BUILD_TAGS=muslc LINK_STATICALLY=true make build

RUN echo "Ensuring binary is statically linked ..." \
  && (file /code/build/$APPNAME | grep "statically linked")

###############################################################
###             STAGE 2: Build platined image             ###
###############################################################

FROM alpine:3.16 AS runner

ENV APPNAME=platined

# Install pre-requisites
RUN apk add --no-cache bash ca-certificates curl jq nano busybox-extras

# Copy over binaries from the builder
COPY --from=builder /code/build/$APPNAME /usr/bin/$APPNAME

# Set user directory and details
ARG HOME_DIR="/home/platined"
ARG USER="platined"
ARG GROUP=${USER}
ARG UID=1000
ARG GID=${UID}

# Add platined user to use in the container
RUN addgroup --system ${GROUP} --gid ${GID} \
    && adduser ${USER} --uid ${UID} -G ${GROUP} --system --home ${HOME_DIR} --shell /bin/bash

# Set entrypoint script
COPY --chown=${USER}:${GROUP} docker/entrypoint.sh /bin/node-start

# Fix windows issue
RUN dos2unix /bin/node-start

RUN chmod +rx /bin/node-start

# Set working directory & bash defaults
WORKDIR ${HOME_DIR}
USER ${USER}
SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

EXPOSE 26656 26657 26660 1317 9090 9091

# Set plain vanilla default entrypoint/command
CMD [ "platined" ]
