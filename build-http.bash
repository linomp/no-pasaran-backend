git pull

docker container stop no-pasaran-backend-http
docker container rm no-pasaran-backend-http

docker image rm no-pasaran-backend-http
docker build -t no-pasaran-backend-http -f DockerfileHttp .

docker run -d --name no-pasaran-backend-http -p 80:80 no-pasaran-backend-http