import pandas as pd

from repositories.postgres import engine


def get_risk():

    query = """
    SELECT
        release_id,
        error_rate,
        latency_ms,
        created_at
    FROM release_metrics
    ORDER BY created_at DESC
    LIMIT 20
    """

    df = pd.read_sql(
        query,
        engine,
    )

    return df.to_dict(
        orient="records"
    )