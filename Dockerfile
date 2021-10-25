FROM alpine:3.10.1

ENV PATH=$PATH:/opt/kcd2021-helmfile

WORKDIR /opt/kcd2021-helmfile

COPY ./migrations /opt/kcd2021-helmfile/migrations/
COPY ./bin/kcd2021-helmfile /opt/kcd2021-helmfile/

RUN chmod +x /opt/kcd2021-helmfile/kcd2021-helmfile

RUN adduser --disabled-password --gecos '' kcd
USER kcd

ENTRYPOINT ["/opt/kcd2021-helmfile/kcd2021-helmfile"]
CMD ["api"]
