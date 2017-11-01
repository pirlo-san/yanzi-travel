docker build -t server:v0.1 .
docker run -d -p 80:80 server:v0.1 /bin/bash