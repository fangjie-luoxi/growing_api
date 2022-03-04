FROM alpine

WORKDIR /
COPY lgapi /lgapi
COPY conf /conf

#RUN apk update \
#    && apk add --no-cache curl ca-certificates bash
#RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
#RUN echo "Asia/Shanghai" >> /etc/timezone
#RUN apk add tzdata

ENTRYPOINT ["/lgapi"]