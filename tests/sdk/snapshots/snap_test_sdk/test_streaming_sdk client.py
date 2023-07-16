# This file was auto-generated by Fern from our API Definition.

import typing

import httpx

from .core.client_wrapper import AsyncClientWrapper, SyncClientWrapper
from .resources.ai.client import AiClient, AsyncAiClient


class FernIr:
    def __init__(self, *, environment: str, timeout: typing.Optional[float] = 60):
        self._environment = environment
        self._client_wrapper = SyncClientWrapper(httpx_client=httpx.Client(timeout=timeout))
        self.ai = AiClient(environment=environment, client_wrapper=self._client_wrapper)


class AsyncFernIr:
    def __init__(self, *, environment: str, timeout: typing.Optional[float] = 60):
        self._environment = environment
        self._client_wrapper = AsyncClientWrapper(httpx_client=httpx.AsyncClient(timeout=timeout))
        self.ai = AsyncAiClient(environment=environment, client_wrapper=self._client_wrapper)
