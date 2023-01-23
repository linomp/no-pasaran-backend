git pull

# This step should be removed... issue is documented here:
# https://github.com/linomp/no-pasaran-backend/issues/1
cp /etc/letsencrypt/live/apps.xmp.systems/cert.pem app/cert.pem
cp /etc/letsencrypt/live/apps.xmp.systems/privkey.pem app/privkey.pem
chmod 777 app/cert.pem
chmod 777 app/privkey.pem

docker container stop no-pasaran-backend-http
docker container rm no-pasaran-backend-http

docker image rm no-pasaran-backend-http
docker build -t no-pasaran-backend-http -f DockerfileHttp .

docker run -d --name no-pasaran-backend-http -p 80:80 no-pasaran-backend-http