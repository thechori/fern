/**
 * Converts the the given name into a valid PHP variable name
 * (i.e. it includes the required '$' prefix).
 */
export function convertToPhpVariableName(name: string): string {
    return `${name.startsWith("$") ? "" : "$"}${name}`;
}
