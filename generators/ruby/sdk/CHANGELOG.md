# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased

## [0.0.5] - 2024-02-06

- This release contains no changes to the API.

- Bug fixes:
  - Syntactic error in block parameter usage: we now ensure block parameters are the final parameter for functions
  - Add properties to subpackages: previously properties on subpackages were not being exposed
  - Ensure optional properties in from_json are only parsed if present

## [0.0.4] - 2024-02-05

- This release contains no changes to the API.

- Bug fixes: ensures files are written at the gem name path over client name, and addresses string escaping on one of the Fern headers.

## [0.0.3] - 2024-02-05

- This release contains no changes to the API.

- Bug fixes: addresses a number of typos and other issues previously contained within the generation code.

## [0.0.2] - 2024-02-01

- Improvement: Consolidate imports into the main file for easier gem usage, as opposed to directly importing the specific file (e.g. `require "fern" vs. require "fern/types/folder_a/object"`)

- Improvement: Added `rubygems` output block to generators.yml. To publish your ruby gem, setup your generators.ym like this:

```yaml
# generators.yml

groups:
 - name: ...
    version: ...
    location:
       output: rubygems
       api-key: ${API_KEY}
       package-name: "petstore"
```

## [0.0.1] - 2024-02-01

- Improvement: Support client generation (async and async) as well as most endpoint types (except streaming)

## [0.0.0] - 2024-01-30

- Chore: Intialize this changelog