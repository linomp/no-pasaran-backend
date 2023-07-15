git pull

docker container stop devtest
docker container rm devtest

docker image rm pointless-backend-debug
docker build -t pointless-backend-debug -f DockerfileDebug .

docker run -it --name devtest -p 443:443 -v /etc/letsencrypt/archive/apps.xmp.systems:/code/app/certs:ro pointless-backend-debug

# uvicorn app.main:app --host "0.0.0.0" --port 443 --ssl-keyfile=./app/certs/privkey3.pem --ssl-certfile=./app/certs/cert3.pem
