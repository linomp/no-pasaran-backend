import datetime

from fastapi import APIRouter, Request

router = APIRouter()


@router.get("/")
async def root(request: Request):
    time = datetime.datetime.now(datetime.timezone.utc).isoformat()
    return {"Host": request.client.host, "Timestamp": time}
