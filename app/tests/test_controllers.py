import pytest

from fastapi.testclient import TestClient

from app.main import app


class TestControllers:

    def test_redirect(self):
        with TestClient(app) as client:
            # test redirection from "/" to "/health"

            response = client.get(f"/")
            assert response.headers["location"] == "/health"

    def test_root(self):
        with TestClient(app) as client:
            response = client.get(f"/health")
            assert response.status_code == 200

            # assert that body contains ServerMetrics fields
            assert "host" in response.json()
            assert "timestamp" in response.json()
            assert "cpu_usage" in response.json()
            assert "memory_usage" in response.json()
