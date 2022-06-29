# Kambing style yaml specification

Kambing style yaml is yaml specification that is mixed with json syntax for better readability. 
This style is supported by the yaml standard and can therefore be read by any complient yaml parser.

## Syntax

### Top level

Top level must not be surrounded by neither curly `{}` nor square `[]` brackets.
Multiple key value pair on the top level must not be sepparated by commas.

### Keys

Keys must not be surrounded by any quotation mark.

### Values

#### Object Value

Object value must be surrounded by curly brackets `{}`.
Multiple key-value-pair in an object must be sepparated by commas `,`.
The last key-value-pair in an object must not have any trailing comma `,`.

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
    key1 : "value1"
    },
    {
    key2: "value2"
    }
]
```
