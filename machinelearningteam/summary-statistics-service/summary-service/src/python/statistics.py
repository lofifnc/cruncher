from functools import partial
from typing import List

import pandas as pd
import logging
from io import StringIO


def column_filter(df, excluded_columns, column):
    if column.startswith("Unnamed") or column in excluded_columns:
        return False
    rows = len(df.values)
    frequency = len(df.groupby(column).count().values)
    if frequency > rows / 2:
        logging.info(f"Excluding {column} too many distinct values: {frequency}.")
        return False
    else:
        return True


def calc_statistics(
    csv: str, aggregate_columns: List[str], excluded_columns: List[str]
) -> str:
    data = StringIO(csv)
    df = pd.read_csv(data)
    group_columns = list(
        filter(partial(column_filter, df, excluded_columns), df.columns.values)
    )
    agg_columns = {"Count": ("CompanyId", "size")}
    agg_columns.update({c + "Sum": (c, "sum") for c in aggregate_columns})

    result = df.groupby(group_columns, as_index=False).agg(**agg_columns)
    out = StringIO()
    result.to_csv(out)
    return out.getvalue()


# with open('../../../test.csv', 'r') as file:
#     print(calc_statistics(file.read(), [], ["BankEntryAmount"]))
