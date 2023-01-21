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
CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "80"]

# Reference commands:

# build image
#docker build -t no-pasaran-be .

# run container
# docker run -t --name no-pasaran-be -p 80:80 no-pasaran-be