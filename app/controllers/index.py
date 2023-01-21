import datetime

from fastapi import APIRouter, Request
from fastapi_versioning import version

router = APIRouter()


@router.get("/")
@version(1)
async def root(request: Request):
    time = datetime.datetime.now(datetime.timezone.utc).isoformat()
    return {"Host": request.client.host, "Timestamp": time}
