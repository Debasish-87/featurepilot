import pandas as pd

from repositories.postgres import engine


def get_adoption():

    query = """
    SELECT
        feature_key,

        COUNT(*) AS total,

        SUM(
            CASE
                WHEN enabled = true
                THEN 1
                ELSE 0
            END
        ) AS enabled
    FROM evaluation_events
    GROUP BY feature_key
    ORDER BY total DESC
    """

    df = pd.read_sql(
        query,
        engine,
    )

    return df.to_dict(
        orient="records"
    )