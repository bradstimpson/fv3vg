FROM python:3.8
ADD . /app
WORKDIR /app
RUN pip install -r fastapi_server/requirements.txt
CMD ["gunicorn", "-k", "uvicorn.workers.UvicornWorker", "-b", "0.0.0.0:8001", "-w", "9", "fastapi_server.server:app"]
EXPOSE 8001
