git pull

# This step should be removed... issue is documented here:
# https://github.com/linomp/no-pasaran-backend/issues/1
cp /etc/letsencrypt/live/apps.xmp.systems/cert.pem app/cert.pem
cp /etc/letsencrypt/live/apps.xmp.systems/privkey.pem app/privkey.pem
chmod 777 app/cert.pem
chmod 777 app/privkey.pem

docker container stop no-pasaran-backend-https
docker container rm no-pasaran-backend-https

docker image rm no-pasaran-backend-https
docker build -t no-pasaran-backend-https -f DockerfileHttps .

docker run -d --name no-pasaran-backend-https -p 443:443 no-pasaran-backend-https