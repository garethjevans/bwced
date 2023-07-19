## scm run

Run the webserver

```
scm run [flags]
```

### Examples

```
bwced run
```

### Options

```
      --bind-address string    The address to bind to (default: localhost) (default "localhost")
  -d, --document-root string   The root to store all documents
      --enable-cors            Whether to allow CORS requests (default: false)
      --max-upload-size int    The max upload size in bytes (default: 1024) (default 1024)
  -p, --port int               The port to run the webserver on (default: 8080) (default 8080)
```

### Options inherited from parent commands

```
  -v, --debug   Debug Output
      --help    Show help for command
```

### SEE ALSO

* [scm](scm.md)	 - provides commands for interacting with different scm providers

