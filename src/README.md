## Manifest.dovego.json

```json5
{
  "name": "DoveGo",
  "version": "0.0.0",
  "server": {
    "port": 8080,
    "shutdown": {
      "url": "" // empty string by default. If you want to use the special URL to shut down the server then input it.
    },
    "start_url": "app/urls/tmpl/index.html", // OK
    // "start_url": "./app/urls/tmpl/index.html" // OK
    // "start_url": "/app/urls/tmpl/index.html"  // Error
  },
  "plugins": [
    {
      "name": "helloWorld",
      "path": "/plugin/hello-world/"
    }
  ],
  "debug": {
    "enable": true
  }
}
```
