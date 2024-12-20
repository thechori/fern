/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../index";
import * as SeedObject from "../../api/index";
import * as core from "../../core";

export const NestedInlineType1: core.serialization.ObjectSchema<
    serializers.NestedInlineType1.Raw,
    SeedObject.NestedInlineType1
> = core.serialization.object({
    foo: core.serialization.string(),
    bar: core.serialization.string(),
    myEnum: core.serialization.lazy(() => serializers.InlineEnum),
});

export declare namespace NestedInlineType1 {
    interface Raw {
        foo: string;
        bar: string;
        myEnum: serializers.InlineEnum.Raw;
    }
}
