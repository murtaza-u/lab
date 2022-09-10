FROM ubuntu:latest
RUN apt-get update -y && apt-get install fortune -y
ADD ./fortuneloop /usr/bin
ENTRYPOINT ["/usr/bin/fortuneloop"]
CMD ["5"]
