from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

from api.traffic import router as traffic_router
from api.overview import router as overview_router
from api.adoption import router as adoption_router
from api.risk import router as risk_router

app = FastAPI(
    title="FeaturePilot Analytics"
)

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

@app.get("/health")
def health():
    return {
        "status": "ok"
    }

app.include_router(
    overview_router,
    prefix="/analytics"
)

app.include_router(
    traffic_router,
    prefix="/analytics"
)

app.include_router(
    adoption_router,
    prefix="/analytics"
)

app.include_router(
    risk_router,
    prefix="/analytics"
)