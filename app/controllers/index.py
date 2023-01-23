import datetime

import psutil as psutil
from app.models.ServerMetrics import ServerMetrics
from fastapi import APIRouter, Request
from fastapi.responses import RedirectResponse

router = APIRouter()


@router.get("/")
async def redirect_typer():
    # redirect to /health
    return RedirectResponse(url="/health")


@router.get("/health", tags=["healthcheck"])
async def root(request: Request) -> ServerMetrics:
    time = datetime.datetime.now(datetime.timezone.utc).isoformat()

    cpu_usage = psutil.cpu_percent()
    memory_usage = psutil.virtual_memory().percent

    return ServerMetrics(host=request.client.host, timestamp=time, cpu_usage=f"{cpu_usage} %",
                         memory_usage=f"{memory_usage} %")
