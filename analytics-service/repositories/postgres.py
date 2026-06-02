import os
from sqlalchemy import create_engine

DATABASE_URL = os.getenv(
    "DATABASE_URL",
    "postgresql://postgres:postgres@postgres:5432/featurepilot",
)

DATABASE_URL = DATABASE_URL.replace(
    "postgres://",
    "postgresql://",
    1,
)

engine = create_engine(DATABASE_URL)