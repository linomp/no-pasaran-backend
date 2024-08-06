from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import RedirectResponse

from app.routers import healthcheck

app = FastAPI(title="pointless-backend", description="Backend for my Pointless personal site")

app.include_router(healthcheck.router)

origins = [
    "https://pointless.xmp.systems",
    "https://apps.xmp.systems:8001",
    "https://apps.xmp.systems:443",
    "http://apps.xmp.systems:80",
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


@app.get("/", tags=["healthcheck"])
async def redirect_to_healthcheck_page():
    return RedirectResponse(url="/health/page")
