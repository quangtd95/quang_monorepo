FROM golang:alpine3.10 as builder
RUN mkdir /compile
ADD ./ /compile/
WORKDIR /compile
RUN go build -mod vendor -o kbond-subscription-service

#######################################################

FROM common-docker.artifactory.kasikornbank.com:8443/alpine:3.10
RUN mkdir /myapp
RUN mkdir /myapp/images
RUN mkdir /myapp/images/barcode
RUN chmod 755 /myapp
COPY --from=builder /compile/kbond-subscription-service /myapp/
COPY --from=builder /compile/kbonduistoeai.crt /myapp/
COPY --from=builder /compile/kbonduistoeai.key /myapp/

# Cannot add tzdata because of firewall. Copy timezone manual instead.
#RUN apk add --no-cache tzdata
COPY --from=builder /compile/assets/timezone/zoneinfo/Asia/Bangkok /etc/localtime
RUN echo "Asia/Bangkok" > /etc/timezone

WORKDIR /myapp


CMD [ "/myapp/kbond-subscription-service" ]
