FROM iron/go:dev
WORKDIR /app
ENV SRC_DIR=/go/src/github.com/volkovartem/waitfor/
# Add the source code:
ADD . $SRC_DIR
# Build it:
RUN cd $SRC_DIR; go build -o waitfor; cp waitfor /app/
ENTRYPOINT ["./waitfor"]