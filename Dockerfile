FROM google/cloud-sdk:alpine
RUN apk add bash go
SHELL ["bash", "-c"]
ENV TZ UTC
WORKDIR /root
RUN go get -v github.com/gorilla/mux
RUN go get -v github.com/gorilla/websocket
RUN go get -v github.com/urfave/negroni
RUN go get -v github.com/robfig/cron
RUN go get -v github.com/shirou/gopsutil/...
COPY . /root/go/src/github.com/livecodecreator/docker-heroku
RUN go build -v -o /bin/entrypoint github.com/livecodecreator/docker-heroku
ENTRYPOINT ["entrypoint"]
