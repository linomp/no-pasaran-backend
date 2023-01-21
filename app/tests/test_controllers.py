import pytest

from fastapi.testclient import TestClient

from app.main import app


class TestControllers:

    @pytest.mark.parametrize("version_prefix", ["/latest", "/v1"])
    def test_root(self, version_prefix):
        with TestClient(app) as client:
            response = client.get(f"{version_prefix}/")
            assert response.status_code == 200

            resp_body = response.json()
            assert "Host" in resp_body
            assert "Timestamp" in resp_body
