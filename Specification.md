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

#### Object Value

Objects other than the top-level YAML object must be surrounded by curly brackets `{}`.
Multiple key-value-pair in an object must be separated by commas `,`.
The last key-value-pair in an object must not have any trailing comma `,`.

**Why:** We define non-top level objects using a JS-like writing style to improve readability for multiple nested objects. We believe this solves a major problem with YAML-readability in exchange for just a little more effort on the developer's end when writing YAML files.

#### String Value

String value must be surrounded by double quotation mark `""`.

#### Number Value

Number value must not be surounded by quotation mark `""` nor any identifier.

#### Array Values

Arrays must be surrounded by square brackets `[]`. Elements of arrays shall be sepparated by commas `,`.

#### Boolean Value

Boolean value must not be surrounded by quotation mark `""` nor any indentifier.

### Indentation
Kambing style yaml doesn't enforce any specific indentation.

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
