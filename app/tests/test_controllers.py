import pytest

from fastapi.testclient import TestClient

from app.main import app


class TestControllers:

    def test_root(self):
        with TestClient(app) as client:
            response = client.get("/latest/")
            assert response.status_code == 200

            resp_body = response.json()
            assert "Host" in resp_body
            assert "Timestamp" in resp_body
