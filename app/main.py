import datetime

from typing import Union

from fastapi import FastAPI, Request

app = FastAPI()

api_version = "v1"
api_prefix = f"api/{api_version}"

@app.get(f"api_prefix/")
async def root(request: Request):

    time = datetime.datetime.now(datetime.timezone.utc).isoformat()

    return {"Host": request.client.host, "Timestamp": time}
