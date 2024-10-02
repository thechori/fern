/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../../index";
import * as FernConjure from "../../../../api/index";
import * as core from "../../../../core";

export const ConjureEnumDeclaration: core.serialization.ObjectSchema<
    serializers.ConjureEnumDeclaration.Raw,
    FernConjure.ConjureEnumDeclaration
> = core.serialization.objectWithoutOptionalProperties({
    values: core.serialization.list(core.serialization.string()),
});

export declare namespace ConjureEnumDeclaration {
    interface Raw {
        values: string[];
    }
}
