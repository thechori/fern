/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as core from "../../../../core";
import urlJoin from "url-join";
import * as errors from "../../../../errors";

export declare namespace Service {
    interface Options {
        environment: core.Supplier<string>;
    }

    interface RequestOptions {
        timeoutInSeconds?: number;
        maxRetries?: number;
    }
}

export class Service {
    constructor(protected readonly _options: Service.Options) {}

    public async getText(requestOptions?: Service.RequestOptions): Promise<unknown> {
        const _response = await core.fetcher({
            url: urlJoin(await core.Supplier.get(this._options.environment), "text"),
            method: "POST",
            headers: {
                "X-Fern-Language": "JavaScript",
                "X-Fern-SDK-Name": "",
                "X-Fern-SDK-Version": "0.0.1",
            },
            contentType: "application/json",
            timeoutMs: requestOptions?.timeoutInSeconds != null ? requestOptions.timeoutInSeconds * 1000 : 60000,
            maxRetries: requestOptions?.maxRetries,
        });
        if (_response.ok) {
            return _response.body;
        }

        if (_response.error.reason === "status-code") {
            throw new errors.SeedPlainTextError({
                statusCode: _response.error.statusCode,
                body: _response.error.body,
            });
        }

        switch (_response.error.reason) {
            case "non-json":
                throw new errors.SeedPlainTextError({
                    statusCode: _response.error.statusCode,
                    body: _response.error.rawBody,
                });
            case "timeout":
                throw new errors.SeedPlainTextTimeoutError();
            case "unknown":
                throw new errors.SeedPlainTextError({
                    message: _response.error.errorMessage,
                });
        }
    }
}