/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as fs from "fs";

/**
 * @example
 *     {
 *         file: fs.createReadStream("/path/to/your/file"),
 *         maybeString: "string",
 *         integer: 1,
 *         maybeInteger: 1,
 *         listOfStrings: "string",
 *         optionalListOfStrings: "string"
 *     }
 */
export interface JustFileWithQueryParamsRequet {
    maybeString?: string;
    integer: number;
    maybeInteger?: number;
    listOfStrings: string | string[];
    optionalListOfStrings?: string | string[];
    file: File | fs.ReadStream;
}