/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../../../../index";
import * as FernIr from "../../../../../../api/index";
import * as core from "../../../../../../core";
import { Declaration } from "../../declaration/types/Declaration";
import { NamedParameter } from "../../types/types/NamedParameter";
import { InlinedRequestBody } from "./InlinedRequestBody";
import { InlinedRequestMetadata } from "./InlinedRequestMetadata";

export const InlinedRequest: core.serialization.ObjectSchema<
    serializers.dynamic.InlinedRequest.Raw,
    FernIr.dynamic.InlinedRequest
> = core.serialization.objectWithoutOptionalProperties({
    declaration: Declaration,
    pathParameters: core.serialization.list(NamedParameter).optional(),
    queryParameters: core.serialization.list(NamedParameter).optional(),
    headers: core.serialization.list(NamedParameter).optional(),
    body: InlinedRequestBody.optional(),
    metadata: InlinedRequestMetadata.optional(),
});

export declare namespace InlinedRequest {
    interface Raw {
        declaration: Declaration.Raw;
        pathParameters?: NamedParameter.Raw[] | null;
        queryParameters?: NamedParameter.Raw[] | null;
        headers?: NamedParameter.Raw[] | null;
        body?: InlinedRequestBody.Raw | null;
        metadata?: InlinedRequestMetadata.Raw | null;
    }
}
