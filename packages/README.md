Packages 
-------

Contains Javascript packages for interacting with the frontend easily.

- DOM reconciling
- Polyfills


## Building

To build Markup typescript sources into javascript using the `commonjs` module system, simply
execute:

```bash
npm run build
```

We use [Browserify](http://browserify.org/) to make these files fully ready for the browser as a single bundle, to get minified and unminified
versions of these files, execute:

```bash
npm run unminified
```

Or

```bash
npm run minified
```


## Example

See demo of DOM patching taking place in [Index.html](./dist/index.html)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Markup Sample Load Page</title>
    <style>
        body{
            padding: 10px;
            width: 960px;
            margin: 0 auto;
        }

        h1#counter{
            text-align: center;
            font-size: 15em;
        }
    </style>
</head>
<body>
    <div id="app"></div>
    <script src="webvm.js" type="text/javascript"></script>
    <script type="text/javascript">
        var counter = 0;
        var mount = new Markup.mount.DOMMount(window.document, "#app");

        setInterval(function patchDOM() {
          counter++;
          mount.patch(`
            <h1 id="counter">${counter}</h1>
          `);
        }, 1000);
    </script>
</body>
</html>
```

![Patch Demo](./media/patch.gif)
