FROM golang:buster
ENV GOPATH /go
RUN apt-get update

COPY d /go/d

WORKDIR /go/d
RUN go get github.com/go-sql-driver/mysql github.com/gorilla/sessions github.com/shurcooL/github_flavored_markdown \
&& go build && mkdir /app && cp d /app && cp -R ./static /app && cp -R ./tmpl /app

WORKDIR /app

CMD /app/d


