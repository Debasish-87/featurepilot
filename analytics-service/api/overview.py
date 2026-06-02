from fastapi import APIRouter

from services.overview_service import (
    get_overview,
)

router = APIRouter()


@router.get("/overview")
def overview():
    return get_overview()