import pytest

from fastapi.testclient import TestClient

from app.main import app


class TestControllers:

    def test_root(self):
        with TestClient(app) as client:
            response = client.get(f"/health")
            assert response.status_code == 200

            # assert that body contains ServerMetrics fields
            assert "host" in response.json()
            assert "timestamp" in response.json()
            assert "cpu_usage" in response.json()
            assert "memory_usage" in response.json()
