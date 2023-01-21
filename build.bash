git pull

# This step should be removed... issue is documented here:
# https://github.com/linomp/no-pasaran-backend/issues/1
cp /etc/letsencrypt/live/apps.xmp.systems/cert.pem app/cert.pm
cp /etc/letsencrypt/live/apps.xmp.systems/privkey.pem app/privkey.pm
chmod 777 app/cert.pm
chmod 777 app/privkey.pm

docker container stop no-pasaran-backend
docker container rm no-pasaran-backend

docker image rm no-pasaran-backend
docker build -t no-pasaran-backend .

docker run -d --name no-pasaran-backend -p 443:443 no-pasaran-backend