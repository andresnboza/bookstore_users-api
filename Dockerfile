FROM golang:1.12.13

# Configure repo url so we can configure our work directory
ENV REPO_URL=github.com/andresnboza/bookstore_users-api

# Setup ou t$GOPATH
ENV GOPATH=/app

ENV APP_PATH=$GOPATH/src/${REPO_URL}

# Copy the entire source code from the current directory to $WORKDIR
ENV WORKPATH=$APP_PATH/src
COPY src $WORKPATH
WORKDIR $WORKPATH

RUN go build -o users-api .

# Expose port 8080 to the world
EXPOSE 8080

CMD ["./users-api"]