# Stage 1: Build
FROM golang:1.21.6 as builder
ENV CGO_ENABLED 0
ARG BUILD_REF

# Create the service directory and copy the module files first.
# download the dependencies. If this doesn't change, we won't need to do this
# again in future builds.
# RUN mkdir /service
# COPY go.* /service/
# WORKDIR /service
# RUN go mod download

# Copy the source code into the container.
COPY . /service

# Build the service binary. We are doing this last since this will be different
# every time we run through this process.
WORKDIR /service/app/services/user-service

# Build the service binary.
RUN go build -ldflags "-X main.build=${BUILD_REF}" -o /user-service

# Stage 2: Final Image
FROM alpine:3.19
ARG BUILD_REF

# Create a non-root user
RUN addgroup -g 1000 -S sales && \
    adduser -u 1000 -h /service -G sales -S sales

# Copy the binary from the builder stage
COPY --from=builder --chown=sales:sales /user-service /service/user-service-api
WORKDIR /service
USER sales

CMD ["./user-service-api"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="user-service-api" \
      org.opencontainers.image.revision="${BUILD_REF}"
