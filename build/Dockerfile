FROM ubuntu:18.04

ENV OPERATOR=/usr/local/bin/gittrack \
    USER_UID=1001 \
    USER_NAME=gittrack

RUN apt-get update -y

RUN apt-get install -y git

# install operator binary
COPY build/_output/bin/gittrack ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
