from fastapi import APIRouter

from services.adoption_service import (
    get_adoption,
)

router = APIRouter()


@router.get("/adoption")
def adoption():
    return get_adoption()