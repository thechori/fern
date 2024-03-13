# This file was auto-generated by Fern from our API Definition.

from __future__ import annotations

import datetime as dt
import typing

from ....core.datetime_utils import serialize_datetime
from .bar import Bar as resources_types_types_bar_Bar
from .foo import Foo as resources_types_types_foo_Foo

try:
    import pydantic.v1 as pydantic  # type: ignore
except ImportError:
    import pydantic  # type: ignore

T_Result = typing.TypeVar("T_Result")


class _Factory:
    def foo(self, value: resources_types_types_foo_Foo) -> UnionWithoutKey:
        return UnionWithoutKey(__root__=_UnionWithoutKey.Foo(**value.dict(exclude_unset=True), type="foo"))

    def bar(self, value: resources_types_types_bar_Bar) -> UnionWithoutKey:
        return UnionWithoutKey(__root__=_UnionWithoutKey.Bar(**value.dict(exclude_unset=True), type="bar"))


class UnionWithoutKey(pydantic.BaseModel):
    factory: typing.ClassVar[_Factory] = _Factory()

    def get_as_union(self) -> typing.Union[_UnionWithoutKey.Foo, _UnionWithoutKey.Bar]:
        return self.__root__

    def visit(
        self,
        foo: typing.Callable[[resources_types_types_foo_Foo], T_Result],
        bar: typing.Callable[[resources_types_types_bar_Bar], T_Result],
    ) -> T_Result:
        if self.__root__.type == "foo":
            return foo(resources_types_types_foo_Foo(**self.__root__.dict(exclude_unset=True, exclude={"type"})))
        if self.__root__.type == "bar":
            return bar(resources_types_types_bar_Bar(**self.__root__.dict(exclude_unset=True, exclude={"type"})))

    __root__: typing.Annotated[
        typing.Union[_UnionWithoutKey.Foo, _UnionWithoutKey.Bar], pydantic.Field(discriminator="type")
    ]

    def json(self, **kwargs: typing.Any) -> str:
        kwargs_with_defaults: typing.Any = {"by_alias": True, "exclude_unset": True, **kwargs}
        return super().json(**kwargs_with_defaults)

    def dict(self, **kwargs: typing.Any) -> typing.Dict[str, typing.Any]:
        kwargs_with_defaults: typing.Any = {"by_alias": True, "exclude_unset": True, **kwargs}
        return super().dict(**kwargs_with_defaults)

    class Config:
        extra = pydantic.Extra.forbid
        json_encoders = {dt.datetime: serialize_datetime}


class _UnionWithoutKey:
    class Foo(resources_types_types_foo_Foo):
        type: typing.Literal["foo"]

        class Config:
            allow_population_by_field_name = True
            populate_by_name = True

    class Bar(resources_types_types_bar_Bar):
        type: typing.Literal["bar"]

        class Config:
            allow_population_by_field_name = True
            populate_by_name = True


UnionWithoutKey.update_forward_refs()