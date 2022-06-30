# Kambing Style YAML Specification

Kambing Style YAML is a YAML formatting specification that leverages YAML's support for JSON syntax to improve readability. 
This style is supported by the YAML standard and can therefore be read by any YAML-compliant parsers.

## Syntax

### Top level

Top level must not be surrounded by neither curly `{}` nor square `[]` brackets.
Multiple key value pair on the top level must not be separated by commas.

```yaml
top1: "Hello"
top2: "World"
top3: "We love goats"
```
**Why:** We allow top-level definitions using the standard YAML syntax to create an implicit rule for defining major sections in a YAML configuration. JSON-style objects are reserved for object definitions within a specific "section", while the YAML syntax is ONLY used for top level declarations. We believe this leads to improved semantics within a YAML document.

### Keys

Keys must not be surrounded by any quotation mark.
```yaml
# Wrong
"key1": "This is wrong."

# Correct
key2: "This is correct!"
"key-2": "This is correct!"

# 
10: "This is correct when you want the key to be implicitly parsed as a 'fixnum'"
"10": "This is correct when you want to force a key to be parsed as a 'string' instead of a 'fixnum'"
```

**Why:** We avoid the usage of quotation marks when defining keys that have a combination of letters and/or numbers because it will be automatically handled by a standard-complying parser. In comparison to JSON this makes it easier for developers to define a key-value pair by allowing them to omit quotation marks in most use-cases, similar to how one would define an object in JavaScript.

### Values

#### Object and Array values

Objects other than the top-level YAML object must be surrounded by curly brackets `{}`.
Arrays must be surrounded by square brackets `[]`. Elements of arrays shall be separated by commas `,`.
Multiple key-value-pair in an object must be separated by commas `,`.
The last key-value-pair in an object must not have any trailing comma `,`.

```yaml
# Standard way of writing YAML.
# This encourages inconsistent usage of quotations and indentation, 
# leading to headaches when a typo causes parsing issues.

version: "3.8"
services:
    db:
        container_name: exampleproject-db
        image: postgres:13.3-alpine
        volumes:
            - db-data:/var/lib/postgresql/data
        restart: always
        environment:
            POSTGRES_USER: ${DB_USER}
            POSTGRES_PASSWORD: ${DB_PASSWORD}
            POSTGRES_DB: ${DB_NAME}
        networks:
            - exampleproject
        ports:
            - "${DB_PORT}:5432"
            
    server:
        container_name: exampleproject-server
        build:
            context: .
            dockerfile: Dockerfile
        ports:
            - "${SERVER_PORT}:${SERVER_PORT}"
        depends_on:
            - db
        environment:
            - DB_HOST=db
            - DB_PORT=5432
        networks:
            - exampleproject
        commands:
            - "/app/main"

volumes:
    db-data:
        driver: local

networks:
    exampleproject:
        driver: bridge
```

```yaml
# KS-YAML compliant formatting
# Although more verbose, this prevents alot of problems stemming from usage of indents in a YAML file.
version: "3.8"
services: {
    db: {
        container_name: "exampleproject-db",
        image: "postgres:13.3-alpine",
        volumes: [
            "db-data:/var/lib/postgresql/data"
        ],
        restart: "always",
        environment: {
            POSTGRES_USER: "${DB_USER}",
            POSTGRES_PASSWORD: "${DB_PASSWORD}",
            POSTGRES_DB: "${DB_NAME}"
        },
        networks: [
            "exampleproject"
        ],
        ports: [
            "${DB_PORT}:5432"
        ]
    },
    server: {
        container_name: "exampleproject-server",
        build: {
            context: ".",
            dockerfile: "Dockerfile"
        },
        ports: [
            "${SERVER_PORT}:${SERVER_PORT}"
        ],
        depends_on: [
            "db"
        ],
        environment: [
            "DB_HOST=db",
            "DB_PORT=5432"
        ],
        networks: [
            "exampleproject"
        ],
        commands: [
            "/app/main"
        ]
    }
}

volumes: {
    db-data: {
        driver: "local"
    }
}

networks: {
    exampleproject: {
        driver: "bridge"
    }
}
```
        
**Why:** We define non-top level objects using a JS-like writing style to improve readability for multiple nested objects. We believe this solves a major problem with YAML-readability in exchange for just a little more effort on the developer's end when writing YAML files. Arrays are also defined in a JS-like syntax to prevent confusion surrounding the use of YAML's native `- <value>` syntax.

#### String Value

String value must be surrounded by double quotation mark `""`.

**Why:** Strings must be explicitly defined using quotation marks to prevent confusion when trying to differentiate between types of variables.

#### Number Value

Number value must not be surounded by quotation mark `""` nor any identifier.

**Why:** Numbers are automatically parsed into integers by standard-compliant parsers, hence we explicitly differentiate between a number definition and a string definition by omitting the quotation marks.

#### Boolean Value

Boolean value must not be surrounded by quotation mark `""` nor any indentifier.

**Why:** This has the same reasoning as with the formatting rules for a number primitive. It prevents confusions between string values such as `"true"` with an actual boolean value `true`

### Indentation
Since KS-YAML does not rely on indentations due to using JSON-style syntax, we do not enforce any kind of indentation. A KS-YAML compliant formatter should give developers an option to change tab sizing either by modifying a global config or a project-specific config.

## Example

```yaml
top level: {
    key1: "value1",
    array: ["a", "b", "c", 1.425]
    string: "this is string",
    number : 3.14159,
    boolean: true
    }

array top: [
    {
    key1: "value1"
    },
    {
    key2: "value2"
    }
]
```
