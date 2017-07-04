FROM iron/go

WORKDIR /app

ADD myapp /app/

RUN cd $SRC_DIR; go build -o myapp; cp myapp /app/

ENV PORT=8080

ENV DB_USERNAME=postgres
ENV DB_PASSWORD=""
ENV DB_NAME=reactapp

ENV TEST_DB_USERNAME=postgres
ENV TEST_DB_PASSWORD=postgres
ENV TEST_DB_NAME=reactapptest

EXPOSE 8080

ENTRYPOINT ["./myapp"]