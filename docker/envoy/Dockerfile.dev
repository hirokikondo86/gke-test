FROM  envoyproxy/envoy:v1.12.2

WORKDIR /envoy
COPY ./docker/envoy/envoy.dev.yaml .
CMD /usr/local/bin/envoy -c /envoy/envoy.dev.yaml

EXPOSE 8080