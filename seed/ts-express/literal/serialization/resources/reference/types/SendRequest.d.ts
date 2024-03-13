/**
 * This file was auto-generated by Fern from our API Definition.
 */
import * as serializers from "../../..";
import * as SeedLiteral from "../../../../api";
import * as core from "../../../../core";
export declare const SendRequest: core.serialization.ObjectSchema<serializers.SendRequest.Raw, SeedLiteral.SendRequest>;
export declare namespace SendRequest {
    interface Raw {
        prompt: "You are a helpful assistant";
        query: string;
        stream: false;
    }
}