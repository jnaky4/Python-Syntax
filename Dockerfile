FROM python:3.8.0

ARG DOCKER_IMAGE_TAG=0
#ARG PIP_INDEX_ARGS= REMOVED

ENV VERSION=$DOCKER_IMAGE_TAG
EXPOSE 5000

ADD requirements.txt .
RUN pip install ${PIP_INDEX_ARGS} --no-cache-dir -r requirements.txt

WORKDIR /project
ADD . /project
RUN useradd appuser && chown -R appuser /project

USER appuser

CMD ["uwsgi","--http=0.0.0.0:5000","--module=app:app"]
