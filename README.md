# fv3vg = flask vs fastapi vs fiber vs gin

An exploration of 4 different web frameworks, 2 in python, 2 in go.  This is inspired from [flask-vs-fastapi](https://github.com/jeremyjordan/flask-vs-fastapi) which is also used as the base code but updated to the latest dependency versions python 3.8.3 and go 1.14.

## Running the containers

Running Flask:

```bash
docker build -t flask-image -f flask_server/Dockerfile .
docker run -it --rm -p 8000:8000 flask-image
locust -f load_testing/locustfile.py --host=http://127.0.0.1:8000
```

Running FastAPI:

```bash
docker build -t fastapi-image -f fastapi_server/Dockerfile .
docker run -it --rm -p 8001:8001 fastapi-image
locust -f load_testing/locustfile.py --host=http://127.0.0.1:8001
```

Running Fiber:

```bash
docker build -t fiber-image -f fiber_server/Dockerfile .
docker run -it --rm -p 8002:8002 fiber-image
locust -f load_testing/locustfile.py --host=http://127.0.0.1:8001
```

Running Gin:

```bash
docker build -t gin-image -f gin_server/Dockerfile .
docker run -it --rm -p 8003:8003 gin-image
locust -f load_testing/locustfile.py --host=http://127.0.0.1:8001
```

All servers have two basic routes (index and predict) and are load tested with Locust. Find it here: [locust.io](https://locust.io/)
