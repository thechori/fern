# This file was auto-generated by Fern from our API Definition.

import typing

from ..types.migration import Migration as resources_migration_types_migration_Migration


class Migration:
    def __init__(self, *, environment: str):
        ...

    def get_attempted_migrations(self) -> typing.List[resources_migration_types_migration_Migration]:
        ...
