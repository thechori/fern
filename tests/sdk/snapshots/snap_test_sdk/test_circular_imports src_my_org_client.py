# This file was auto-generated by Fern from our API Definition.

import typing

import httpx

from .core.client_wrapper import AsyncClientWrapper, SyncClientWrapper


class MyOrgIr:
    def __init__(self, *, environment: str, timeout: typing.Optional[float] = 60):
        self._environment = environment
        self._client_wrapper = SyncClientWrapper(httpx_client=httpx.Client(timeout=timeout))


class AsyncMyOrgIr:
    def __init__(self, *, environment: str, timeout: typing.Optional[float] = 60):
        self._environment = environment
        self._client_wrapper = AsyncClientWrapper(httpx_client=httpx.AsyncClient(timeout=timeout))
