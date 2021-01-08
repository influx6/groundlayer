# Miru
Miru takes the standard go template parser implementation adding more features with a different
end result where templates are parsed into pure go based packages which can be rendered faster and
more efficiently. This allows us bundle this templates directly into binaries among other benefits.

## Why Miru

Miru was created to allow an optimized generated form of any html template which would allow users
build web application using templates whilst still benefiting from a componentized view for their application.

We hope to use the benefit of a simplified html template with the speed that brings in regards setup, use and 
time compared to the time investment required by new web frameworks and libraries.


```go
var dir = Miru.NewVDir("./testdata")
var indexFile, err = dir.GetFile("index.html")
if err != nil {
	log.Fatal("Failed: ", err)
}

var parsedData, parseErr = indexFile.Parse(Miru.DefaultOption, nil)
if parseErr != nil {
	log.Fatal("Failed: ", err)
}

// Format takes the name of the package to be used for generated file.
var parsedString, parsedStrErr = parsedData.Format("example")
if parsedStrErr != nil {
	log.Fatal("Failed: ", err)
}

// parsedString contains generated go content
```

## Miru Template Language

Whilst Miru utilizes the native Go template package [text/template](https://golang.org/pkg/text/template/), 
alot of changes were done to improve support for native html handling and other features added. Hence providing
a clear and detailed view of these changes and how  they are expected to work is required. 

The described sections below are the new support capabilities which the user should be aware of as they use
Miru.


#### Template Gotchas

1. Miru provides limited internal functions even compared to go/template packages, hence only basic types are supported in the default functions that come with Miru templates. Developers can include custom helpers through [Package Directives](#package-directives) to include packages customly provided by the developer and used as functions in the templates.

2. Miru generates a go file out of a template hence it's preset and defined and not customizable how that output looks like.

### Package Directives

Miru adds the ability to include a external package into the template to provide package level variables, constants and functions. These allows a broader level of support and customization to developers.


```go
{{ import plackers | github.com/influx6/groundlayer/pkg/plackers }}
```

Above is a package directive with declared alias `plackers` pointed at package path `github.com/influx6/groundlayer/pkg/plackers`.

### Data and Types Directives
Miru adds the ability to define structs and types within the templates itself to allow a form of typesafety in the final generated code. This becomes important due to the naturally need to define
template entry expectations.

#### Model Directive
Model directives is used to define a struct type within the templates and there can be many defined across files and in templates with the only constraint is to have unique names. 

```go
{{model Address | Zip:String, Street: string, Number:Int }}
{{model Product | Isin:String, Price: Float, Address:Address }}
```

#### ModelType Directive
ModelType directives allows defining a typed definition which is exactly as it sounds as in go. A custom type aliased to another type be it basic or defined type is used as the base.

In go code:

```go
type Name string
```

is exactly the following in Miru:


```go
{{modeltype productmapping |  string }}
```

Miru provides the `Any` key to represent the `interface{}` base type.

```go
{{modelType ProductLog | Any }}
```

For List and Map types, there is specific rules unlike in go code, where we use `(KeyType, ValueType)` as the blocks to define the `Key` type and `Value` type.

```go
{{modelType ProductMapping |  Map(string, Product) }}
{{modelType ListProductMapping |   List(ProductMapping) }}
```

### Method Directives

Miru adds ability to define methods which specific purpose is to encapsulate specific UI views as methods which can be 
re-used across different parts of the template. They are similar to the [define](#define-blocks), but typed argument.

```go
{{method renderSubProducts | product:Product, subProducts:ProductMapping }}
	<div class="product">
	{{if len subProducts}}
		<div class="item">
		<span class="price">No sub products available</span>
		</div>
	{{else}}
		<div class="item">
		<span class="price">Categories</span>
		</div>
	{{end}}
	</div>
{{endMethod}}
```


#### Component Directive

Component directive provides a clear means of defining possible components which are argument to the root template.

In essence, you tell the compiler to create a grouping which allows injecting a [peji](pkg/peji/components.go)
which will be mounted as desirable as part of the layout. This provides the benefit of clarity in how a layout expectations
on its parts.

```go
{{Komponent UserComponent }}
{{Komponent profileComponent }}
```

You can then mount these components within your template using the `mount` keyword with provided context:

```go
{{ mount UserComponent . }}
```

The `mount` keyword cannot be used like other identifiers or functions, it's a special purpose identifier which does not
work with pipes or any other composition directive. It exists to instruct generated code to expect said component and mount
it with provided argument using the current scope (the current parent node) as root for the component.

#### Dom Identifiers

DOM identifiers exists to allow you to query for the name of the parent, root or previous dom parent from current, due 
to some need to pass this off to a function or to a custom function which appends these to it's parent.

Hence, there exists 3 identifiers to be used: `dom`, `parentdom` and `rootdom`.

#### Noop directive
The `noop` function call in Miru exists due to the fact that all function calls, fields, text, variables are automatically attached to 
the current node which is the parent of others (where the final is the root node if no html exists). This occurs because Miru is a compiled
expression of a output to be rendered, hence unless it's a variable declaration, or assignment then it must lead to some result to be displayed.

Due to this, the `noop` function name exists which will indicate to the 
parser to not wrap the result in a `helpers.AttachToNode` function, as we
may wish to have a function call or calls that have no output result.

##### Bad form:

```go
{{ print .TotalLives | noop | printMore }}
```

##### Correct form:

```go
{{ print .TotalLives | printMore | noop }}
```

### Basic Template Elements

As already provided by the [text/template](https://golang.org/pkg/text/template/) support for texts tokens, assignment operations, variables, functions, range statements, if statements, and pipes are maintained.

#### Text blocks 

```go
{{ $name := "alex" }}

I was born on the day of {{ $name }}, but these needs to be counted as {{ $name | count }} to be certain of the name length.

```

#### Pipe blocks 


```go
{{  .Field3 | print }}
```

#### MultiPipe Block

*A special thing to note is when having multiple pipe blocks that each pipe block should be wrapped in `(` and `)`, which if absent may not be properly handled in special cases.

##### Bad form:

```go
{{ $x := "wonder" }}
{{ $x = (print .TotalLives "alex" $x | print) }}
```

##### Correct form:

```go
{{ $x := "wonder" }}
{{ $x = (print .TotalLives "alex" $x) | print }}
```

##### Correct form:

```
{{ $x := "wonder" }}
{{ $x = ((print .TotalLives "alex" $x) | print) }}
```

#### Variable blocks 

```go
{{printf .Field1.Field2.Field3}}
{{(printf .Field1.Field3.Field6).Value}}
```

```go
{{$x := (printf .Field1.Field2.Field3 .Field1.Field2).Value}}
{{$y := (printf $x.Field1.Field2.Field3).Value}}
{{$ui = $y }}
```

#### If blocks 

```go
Hello, {{if .Name}} {{.Name}} {{else}} there {{end}}!
```

```go
{{if and .User .User.Admin}}
	You are an admin user!
{{else if or .User .User.Admin }}
	Access denied!
{{end}}
```

#### Range blocks 

```go
{{range .Items}}
	Hello, {{.Name}} 
{{else}}
	No items available
{{end}}
```

```go
{{range $x, $y := .Items}}
	Hello, {{$y.Name}}  from {{ $x }} queue.
{{else}}
	No items available
{{end}}
```

#### With blocks 

```go
{{with contains .Items }}
	{{printf "%q" . | rewrap }}
{{else}}
	{{printf "%d %d %d" 11 11 11}}
{{end}}
```

```go
{{with .Field | contains "boring" | contains "water" }}
	{{printf "%q" . | printf "%s"}}
{{else}}
	{{printf "%d %d %d" 11 11 11}}
{{end}}
```

#### Defined templates

```go
This is rendered text

{{range .Items}}
    <div class="item">{{.}}</div>
{{end}}

{{template "footer.tmper" .}}
```

We added the ability to reference specific `define` blocks all defined in a single file from the `template`
directive.

- File1.html
```go
{{define "footer" }}
	This is rendered text
	{{range .Items}}
		<div class="item">{{.}}</div>
	{{end}}
{{end}}

{{define "fender" }}
	This is rendered text
	{{range .Items}}
		<div class="item">{{.}}</div>
	{{end}}
{{end}}
```

- File 2
```go
{{template "footer.tmper" .}}
{{template "footer.tmper" .}}
```

Where each template specifically targets the specified file.

With the addition of [Data and Types Directives](#data-and-types-directives), define blocks can also be typed:

```go
{{modelType ItemList | List(string) }}
{{model Product | Items: ItemList }}

This is rendered text
```


### Html tags
Html tags are now supported directly within the text parsing code, hence allowing you to write html text like any other text.

```go
{{range $x, $y := .Items}}
	<div class="item">
		<h3 class="name item-{{$x}}">{{.Name}}</h3>
		<span class="price">{{$y.Price}}</span>
	</div>
{{else}}
	No data to count
{{end}}
```


### Theme Attributes
The current templates had supports for theme attributes in [Tailwind](https://tailwindcss.com/) style.
It allows specifying such values if so desired which will be applied and generated accordingly to the
element.

```go
<p
    class="bob"
    theme=[
        color-red-50
        sm-padding-10
        max-w
        max-h
        {{if and .User .User.Admin}}
            bob
        {{else}}
            bob-10
        {{end}}
    ]
    id=wan
>It was you</p>
```

One thing to notice with list attributes for html is a limitation in the parser which 
requires that space be only used to separate values and any value which is not meant to be
separated should generally only have space after it's definition, so:

##### Bad form:

```go
<p
    class="bob"
    theme=[
        color-red-50
        sm-padding-10
        xs:(bg-color-500, fg-color-400) // this is wrong
    ]
    id=wan
    >It was you</p>
```

##### Good form:

```go
<p
    class="bob"
    theme=[
        color-red-50
        sm-padding-10
        xs:(bg-color-500,fg-color-400) // this is correct, no space in-between
    ]
    id=wan
>It was you</p>
```
