FROM alpine:3.10
LABEL DevOps Team <devops@opportunityto.work>

RUN apk add -U --no-cache ca-certificates

ENTRYPOINT []
WORKDIR /

ADD DataService /

EXPOSE 8000

CMD ["./DataService"]