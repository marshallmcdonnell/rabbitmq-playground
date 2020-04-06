FROM python:3.7

ENV PYTHONDONTWRITEBYTECODE 1
ENV FLASK_APP "app.py"
ENV FLASK_ENV "development"
ENV FLASK_DEBUG True

RUN mkdir /app
WORKDIR /app

COPY Pip* /app/

RUN apt update && \
    apt install netcat -y && \
    pip install --upgrade pip && \
    pip install pipenv && \
    pipenv install --dev --system --deploy --ignore-pipfile && \
    rm -rf /var/lib/apt/lists/*

ADD . /app

EXPOSE 3000
CMD ["bash", "docker-start.sh"]
