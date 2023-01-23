from fastapi import FastAPI

from app.controllers import index

app = FastAPI(title="no-pasaran-backend", description="Backend for no-pasaran")

app.include_router(index.router)
