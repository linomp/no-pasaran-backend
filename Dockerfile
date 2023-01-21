# Source: https://fastapi.tiangolo.com/deployment/docker/#build-a-docker-image-for-fastapi

FROM python:3.9

# set work directory
WORKDIR /code

# copy requirements.txt first to leverage Docker cache
COPY ./requirements.txt /code/requirements.txt
RUN pip install --no-cache-dir --upgrade -r /code/requirements.txt

# copy project
COPY ./app /code/app

# run server
CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "443", "--ssl-keyfile=./app/privkey.pem", "--ssl-certfile=./app/cert.pem"]

# run bash as the default command
# CMD ["/bin/bash"]

# Reference commands:

# build image
# docker build -t no-pasaran-backend .

# run container

# HTTP
# docker run -d --name devtest -p 80:80 no-pasaran-backend
# uvicorn app.main:app --host 0.0.0.0 --port 80

# HTTPS
# docker run -t --name devtest --mount type=bind,source="/etc/letsencrypt/live/apps.xmp.systems",target=/certs -p 80:80 no-pasaran-backend
# uvicorn app.main:app --host 0.0.0.0 --port 80 --ssl-keyfile=./app/privkey.pem --ssl-certfile=./app/cert.pem
