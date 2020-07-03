FROM python:3.8-slim-buster
ADD . /app
WORKDIR /app
RUN pip install -r flask_server/requirements.txt
CMD ["gunicorn", "-b", "0.0.0.0:8000","-w","9", "flask_server.server:app"]
# CMD tail -f /dev/null
EXPOSE 8000