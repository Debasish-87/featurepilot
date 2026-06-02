import pandas as pd

from repositories.postgres import engine


def get_overview():

    query = """
    SELECT
        (SELECT COUNT(*) FROM features) AS total_features,
        (SELECT COUNT(*) FROM evaluation_events) AS total_evaluations,
        (
            SELECT COUNT(*)
            FROM releases
            WHERE status IN (
                'ACTIVE',
                'ROLLING_OUT'
            )
        ) AS active_releases
    """

    df = pd.read_sql(
        query,
        engine,
    )

    return df.iloc[0].to_dict()