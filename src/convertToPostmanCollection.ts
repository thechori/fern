import { ApiAuth, AuthScheme } from "@fern-fern/ir-model/auth";
import { ErrorDeclaration } from "@fern-fern/ir-model/errors";
import { IntermediateRepresentation } from "@fern-fern/ir-model/ir";
import { HttpEndpoint, HttpService } from "@fern-fern/ir-model/services/http";
import { TypeDeclaration } from "@fern-fern/ir-model/types";
import {
    PostmanCollectionEndpointItem,
    PostmanCollectionItem,
    PostmanCollectionSchema,
    PostmanHeader,
} from "@fern-fern/postman-sdk/resources";
import { startCase } from "lodash";
import { convertAuth, getAuthHeaders, getVariablesForAuthScheme } from "./auth";
import { convertExampleEndpointCall } from "./convertExampleEndpointCall";
import { GeneratedDummyRequest } from "./request/GeneratedDummyRequest";
import { GeneratedExampleRequest } from "./request/GeneratedExampleRequest";
import { ORIGIN_VARIABLE_NAME } from "./utils";

const ORIGIN_DEFAULT_VAULE = "http://localhost:8080";

export function convertToPostmanCollection(ir: IntermediateRepresentation): PostmanCollectionSchema {
    const id = ir.apiName;

    const authSchemes = filterAuthSchemes(ir.auth);
    const authHeaders = getAuthHeaders(authSchemes);

    return {
        info: {
            name: ir.apiDisplayName ?? id,
            schema: "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
        },
        variable: [
            {
                key: ORIGIN_VARIABLE_NAME,
                value: ORIGIN_DEFAULT_VAULE,
                type: "string",
            },
            ...authSchemes.flatMap(getVariablesForAuthScheme),
        ],
        auth: convertAuth(authSchemes),
        item: getCollectionItems({ ir, authHeaders }),
    };
}

function filterAuthSchemes(auth: ApiAuth): AuthScheme[] {
    let hasSeenAuthorizationHeader = false;
    return auth.schemes.filter((scheme) => {
        return AuthScheme._visit(scheme, {
            basic: () => {
                if (hasSeenAuthorizationHeader) {
                    return false;
                }
                return (hasSeenAuthorizationHeader = true);
            },
            bearer: () => {
                if (hasSeenAuthorizationHeader) {
                    return false;
                }
                return (hasSeenAuthorizationHeader = true);
            },
            header: () => true,
            _unknown: () => {
                throw new Error("Unknown auth scheme: " + scheme._type);
            },
        });
    });
}

function getCollectionItems({
    ir,
    authHeaders,
}: {
    ir: IntermediateRepresentation;
    authHeaders: PostmanHeader[];
}): PostmanCollectionItem[] {
    const rootItems: PostmanCollectionItem[] = [];
    for (const httpService of ir.services.http) {
        let container = rootItems;
        for (const fernFilepathPart of httpService.name.fernFilepathV2) {
            const existingContainerForFernFilepathPart = rootItems
                .filter((item): item is PostmanCollectionItem.Container => item.type === "container")
                .find((item) => item.name === fernFilepathPart.unsafeName.originalValue);
            if (existingContainerForFernFilepathPart != null) {
                container = existingContainerForFernFilepathPart.item;
            } else {
                const newContainer = PostmanCollectionItem.container({
                    name: startCase(fernFilepathPart.unsafeName.originalValue),
                    item: [],
                });
                container.push(newContainer);
                container = newContainer.item;
            }
        }

        container.push(
            ...httpService.endpoints.map((httpEndpoint) =>
                PostmanCollectionItem.endpoint(
                    convertEndpoint({
                        authHeaders,
                        httpEndpoint,
                        httpService,
                        allTypes: ir.types,
                        allErrors: ir.errors,
                    })
                )
            )
        );
    }
    return rootItems;
}

function convertEndpoint({
    authHeaders,
    httpEndpoint,
    httpService,
    allTypes,
    allErrors,
}: {
    authHeaders: PostmanHeader[];
    httpEndpoint: HttpEndpoint;
    httpService: HttpService;
    allTypes: TypeDeclaration[];
    allErrors: ErrorDeclaration[];
}): PostmanCollectionEndpointItem {
    const example = httpEndpoint.examples[0];
    const generatedRequest =
        example != null
            ? new GeneratedExampleRequest({ authHeaders, httpService, httpEndpoint, allTypes, example })
            : new GeneratedDummyRequest({ authHeaders, httpService, httpEndpoint, allTypes });

    return {
        name: httpEndpoint.id,
        request: generatedRequest.get(),
        response: httpEndpoint.examples.map((example) =>
            convertExampleEndpointCall({ authHeaders, httpService, httpEndpoint, allTypes, allErrors, example })
        ),
    };
}
