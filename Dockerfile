# ------------------------------------------------------------------------
# This is a sample Dockerfile to build the ArdynGolang sample application
# ------------------------------------------------------------------------

FROM golang:1.17.1-alpine3.14

# We require musl-dev to compile the application
RUN apk add build-base musl-dev

# Create the app dir to store our final application executable
RUN mkdir /app

# Create a temporary directory to build our executable application. This directory 
# will be deleted once the build is complete.
RUN mkdir /tmpbuild

# Switch to the temporary directory
WORKDIR /tmpbuild

# Add the source files
ADD . .

# As go to get the modules and start the build
RUN go get -tags musl ./...
RUN go build -a -v -tags musl -o /app/myservice src/*.go

# If all goes well, switch to the root directory and 
# delete the temporary directory
WORKDIR /
RUN rm -rf /tmpbuild

# Expose the port as conigured in the Docker environment variables
EXPOSE ${CONFIG_PORT}

# Switch to the directory containing our executable
WORKDIR /app

# Run the executable and pass the full path and filename of the properties file.
# If your configuration is stored in ArdynSoothsayer, then give the full URL 
# of the configuration file in .properties format
CMD ["/app/myservice", "-config=${CONFIG_LOCATION}"]