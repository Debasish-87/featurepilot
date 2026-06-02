from fastapi import APIRouter

from services.traffic_service import (
    get_traffic,
)

router = APIRouter()


@router.get("/traffic")
def traffic():

    return get_traffic()