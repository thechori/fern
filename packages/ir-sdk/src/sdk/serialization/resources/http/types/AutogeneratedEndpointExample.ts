/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../..";
import * as FernIr from "../../../../api";
import * as core from "../../../../core";

export const AutogeneratedEndpointExample: core.serialization.ObjectSchema<
    serializers.AutogeneratedEndpointExample.Raw,
    FernIr.AutogeneratedEndpointExample
> = core.serialization.objectWithoutOptionalProperties({
    example: core.serialization.lazyObject(async () => (await import("../../..")).ExampleEndpointCall),
});

export declare namespace AutogeneratedEndpointExample {
    interface Raw {
        example: serializers.ExampleEndpointCall.Raw;
    }
}