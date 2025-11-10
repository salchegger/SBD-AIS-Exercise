#### 1: Builder

# Uses the official Go 1.25 image as the build environment
# Contains Go tools needed to compile the backend application
# AS builder names this stage so its output can be copied later
FROM golang:1.25 AS builder

#Sets the working directory inside the container to /app
#all following commands run in this directory.
WORKDIR /app

#Copies only dependency files first
#allows Docker to cache dependencies and avoid re-downloading them every build
COPY go.mod go.sum ./

#Downloads all Go module dependencies
#Speeds up later builds by reusing cached modules
RUN go mod download

#Copies the rest of the source code into the image
COPY . .

#Makes the custom build script executable and runs it
#compiles the Go project -> and creates the ordersystem binary
RUN chmod +x scripts/build-application.sh && ./scripts/build-application.sh




#### 2: Runner

#Starts a minimal, lightweight Alpine Linux image (~5 MB)
#Keeps the final image small and secure
FROM alpine:latest

#Creates a non-root user named appuser with a home directory -> improves security
RUN adduser -D -h /home/appuser appuser

#Installs required tools
  #ca-certificates -> enables HTTPS requests
  #bash -> for shell scripts (e.g., startup scripts)
  #netcat-openbsd -> used by wait-for-postgres.sh to check database availability
RUN apk add --no-cache ca-certificates bash netcat-openbsd

#Sets /app as the working directory again (in this new image)
WORKDIR /app

#Copies the compiled binary (ordersystem) from the builder stage into the runner image
#Avoids copying source code or dependencies, making the image much smaller
COPY --from=builder /app/ordersystem /app/ordersystem

#Gives file ownership to appuser and makes it executable
#Ensures proper permissions for the non-root user
RUN chown appuser:appuser /app/ordersystem && chmod +x /app/ordersystem

#Switches to appuser for running the application
#Prevents root-level access inside the container (best practice)
USER appuser


#EXPOSE doesn't actually do anything!
#Documents that the app listens on port 3000
#This doesn’t open the port; it’s just for reference and tooling
EXPOSE 3000

#Defines the command to run when the container starts
#Launches the compiled Go backend binary
ENTRYPOINT ["./ordersystem"]
