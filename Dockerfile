FROM scratch
MAINTAINER Rafael Jesus <rafaelljesus86@gmail.com>
ADD event-tracker /event-tracker
ENV TRACKING_REST_PORT="3000"
ENTRYPOINT ["/event-tracker"]
