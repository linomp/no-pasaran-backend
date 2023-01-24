from fastapi import FastAPI

from app.controllers import healthcheck
from fastapi.responses import RedirectResponse

app = FastAPI(title="no-pasaran-backend", description="Backend for no-pasaran")

app.include_router(healthcheck.router)


@app.get("/", tags=["healthcheck"])
async def redirect_to_healthcheck_page():
    return RedirectResponse(url="/health/page")
