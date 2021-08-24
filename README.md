# Scaffolder CLI

Go CLI to generate files based on templates.

## Usage

```
// dev version
go run . make {template name or relative path} {relative output path}

// or built version
./scaffolder make {template name or relative path} {relative output path}
```

## Building

```
go build
```

## Internal templates

Templates can be embedded into the CLI. These are made available to the user without any extra configuration. To add a new internal template, add a file on the `./templates/default-templates` folder and create a reference to it on the `./templates/defaultTemplates.go` file. For example:

```go
//go:embed templates/name-of-the-template.mustache
var nameOfTheTemplate string

var List = map[string]string{
	"name-of-the-template": nameOfTheTemplate,
}
```

## Custom templates

### Writing your template

Templates can use [mustache syntax](https://mustache.github.io/mustache.5.html). With it, you can have dynamic content on your templates. By default, you have access to a set of variables, including the output path and a set of variations of the file name. The name of the file can be defined using the `--name` flag or it's derived from the output file name. In case your filename is `index`, the name will be set to the output's dir.

Variations of the name are passed along as well, giving you the flexibility to use the variant more appropriate to the conventions of you file. For example, in JavaScript is common to use kebab case for folders and pascal case for components. Having these variations in hand allows your template to choose the most appropriate.

It takes the following shape:

```go
{
  // Output path
  PATH: string
  // File name (e.g. my component)
  NAME: {
    // The base for of the name and its variations
    VALUE              string // Raw value: my component
    UPPERCASE          string // UPPERCASE value: MY COMPONENT
    LOWERCASE          string // lowercase value: my component
    KEBAB              string // kebab-case value: my-component
    CAMEL              string // camelCase value: myComponent
    SNAKE              string // snake_case value: my_component
    PASCAL             string // PascalCase value: MyComponent
    MACRO              string // MACRO_CASE value: MY_COMPONENT

    // The singular for of the name and its variations
    SINGULAR_VALUE     string // Singular Raw value: my component
    SINGULAR_UPPERCASE string // Singular UPPERCASE value: MY COMPONENT
    SINGULAR_LOWERCASE string // Singular lowercase value: my component
    SINGULAR_KEBAB     string // Singular kebab-case value: my-component
    SINGULAR_CAMEL     string // Singular camelCase value: myComponent
    SINGULAR_SNAKE     string // Singular snake_case value: my_component
    SINGULAR_PASCAL    string // Singular PascalCase value: MyComponent
    SINGULAR_MACRO     string // Singular MACRO_CASE value: MY_COMPONENT

    // The plural for of the name and its variations
    PLURAL_VALUE       string // Plural Raw value: my components
    PLURAL_UPPERCASE   string // Plural UPPERCASE value: MY COMPONENTs
    PLURAL_LOWERCASE   string // Plural lowercase value: my components
    PLURAL_KEBAB       string // Plural kebab-case value: my-components
    PLURAL_CAMEL       string // Plural camelCase value: myComponents
    PLURAL_SNAKE       string // Plural snake_case value: my_components
    PLURAL_PASCAL      string // Plural PascalCase value: MyComponents
    PLURAL_MACRO       string // Plural MACRO_CASE value: MY_COMPONENTs
  }
}
```

### Custom template registration

To make use of the custom template, create a `config.json` file on the CWD containing:

```json
{
  "templates": {
    "name-of-the-template": "./relative/path/to/template/file.mustache"
  }
}
```

## Flags

### Custom config file path

By default, the CLI will look for a `./config.json` file on the CWD. If you want to store this under a different name or path, you can pass a `--config` flag with a relative path to the CWD.

### Custom name

If you wish to specify the `NAME` variable, you can pass a `--name` flag.
