FROM ubuntu:latest
RUN apt update && apt install -y curl
COPY webserver /bin/webserver
CMD /bin/webserver

# -- Build our go app for linux so it can run on docker
# GOOS=linux go build

# -- build our image and tag with the name webserver
# docker build -t webserver .

# -- now run it in an interactive terminal
# docker run -ti webserver

# -- but without forwarding the port to the OS it won't work
# docker run -ti -p 8000:8080 webserver

# -- our image was given the name 'nice_burnell' and we can get a shell into that image
# docker exec -ti nice_burnell bash

# -- if we don't know the image name we can check it using images command
# docker images
