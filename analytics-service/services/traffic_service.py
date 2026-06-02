import pandas as pd

from repositories.postgres import engine


def get_traffic():

    query = """
    SELECT
        feature_key,
        COUNT(*) AS requests
    FROM evaluation_events
    GROUP BY feature_key
    ORDER BY requests DESC
    """

    try:
        df = pd.read_sql(
            query,
            engine,
        )

        return df.to_dict(
            orient="records"
        )

    except Exception as e:
        return {
            "error": str(e)
        }