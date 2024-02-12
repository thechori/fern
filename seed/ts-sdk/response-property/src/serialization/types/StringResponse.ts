/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "..";
import * as SeedResponseProperty from "../../api";
import * as core from "../../core";

export const StringResponse: core.serialization.ObjectSchema<
    serializers.StringResponse.Raw,
    SeedResponseProperty.StringResponse
> = core.serialization.object({
    data: core.serialization.string(),
});

export declare namespace StringResponse {
    interface Raw {
        data: string;
    }
}