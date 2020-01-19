FROM python:3.7

RUN apt-get update -y  && \
    apt-get install -y python3.7-dev

COPY ./requirements.txt /app/requirements.txt

WORKDIR /app

RUN pip3 install -r requirements.txt

COPY flask_web /app
EXPOSE 5000
ENTRYPOINT [ "python3" ]

CMD [ "app.py" ]
