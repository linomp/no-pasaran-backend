import datetime
import psutil as psutil

from fastapi import Request

from app.models.ServerMetrics import ServerMetrics


def get_status(request: Request) -> ServerMetrics:
    time = datetime.datetime.now(datetime.timezone.utc).isoformat()

    cpu_usage = psutil.cpu_percent()
    memory_usage = psutil.virtual_memory().percent

    return ServerMetrics(host=request.client.host, timestamp=time, cpu_usage=f"{cpu_usage} %",
                         memory_usage=f"{memory_usage} %")
