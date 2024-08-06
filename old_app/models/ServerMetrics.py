from pydantic import BaseModel


class ServerMetrics(BaseModel):
    host: str
    timestamp: str
    cpu_usage: str
    memory_usage: str
