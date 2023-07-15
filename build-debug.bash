git pull

## This step should be removed... issue is documented here:
## https://github.com/linomp/no-pasaran-backend/issues/1
#cp /etc/letsencrypt/live/apps.xmp.systems/cert.pem app/cert.pem
#cp /etc/letsencrypt/live/apps.xmp.systems/privkey.pem app/privkey.pem
#chmod 777 app/cert.pem
#chmod 777 app/privkey.pem

docker container stop devtest
docker container rm devtest

docker image rm pointless-backend-debug
docker build -t pointless-backend-debug -f DockerfileDebug .

docker run -it --name devtest -p 8080:8080 -p 8443:8443 -v /etc/letsencrypt/archive/apps.xmp.systems:/code/certs:ro pointless-backend-debug
