FROM scratch
MAINTAINER Rafael Jesus <rafaelljesus86@gmail.com>

ADD event-tracker /event-tracker

ENV TRACKING_REST_PORT="3000"
ENV ELASTIC_SEARCH_URL="http://@docker:9200"

ENTRYPOINT ["/event-tracker"]
