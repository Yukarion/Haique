FROM python:3.8 AS builder
WORKDIR /app
COPY ./BFF .
RUN pip install pipenv \
 && pipenv install --system
EXPOSE 5000
ENTRYPOINT ["python3","serve.py"]
