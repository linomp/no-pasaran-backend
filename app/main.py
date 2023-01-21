from fastapi import FastAPI

from fastapi_versioning import VersionedFastAPI

from app.controllers import index

app = FastAPI(title="no-pasaran-backend", description="Backend for no-pasaran")

app.include_router(index.router)

app = VersionedFastAPI(app, version_format='{major}', prefix_format='/v{major}', enable_latest=True)
