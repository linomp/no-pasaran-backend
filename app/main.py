import datetime

from typing import Union

from fastapi import FastAPI, Request

app = FastAPI()


@app.get("/")
async def root(request: Request):

    time = datetime.datetime.now(datetime.timezone.utc).isoformat()

    return {"Host": request.client.host, "Timestamp": time}
