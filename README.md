# Language detection plugin

This plugin is a wrapper for the [whatlanggo](https://github.com/abadojack/whatlanggo) natural language detection library.

## Installation

Follow the [instructions](https://docs.halon.io/manual/comp_install.html#installation) in our manual to add our package repository and then run the below command.

### Ubuntu

```
apt-get install halon-extras-language-detection
```

### RHEL

```
yum install halon-extras-language-detection
```

## Exported functions

### detect_language(text)

Detect the language of a string.

**Params**

- text `string` - The text

**Returns**

The language as a string. On error an exception will be thrown.

**Example**

```
echo detect_language("This is a text in English"); // "English"
```