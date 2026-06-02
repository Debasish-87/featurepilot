from fastapi import APIRouter

from services.risk_service import (
    get_risk,
)

router = APIRouter()


@router.get("/risk")
def risk():
    return get_risk()