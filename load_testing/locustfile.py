import numpy as np
from locust import HttpUser, task, between


class MyLocust(HttpUser):

    wait_time = between(0.05,0.2)

    @task
    def predict(self):
        payload = {"X": np.random.randn(2, 13).tolist()}
        self.client.post("/predict", json=payload)

