# Healthcheck

A REST API to provide your healcheck scripts result.

## Install

```bash
make install
```

> Use `sudo make install` to export some admin data(like network speed).

## Test

```bash
make test
```

## Uninstall

```bash
make uninstall
```

## Use
 - add some healcheck scripts inside the `scripts` folder
 - install the process
 - curl the corresponding address(by default `curl http://localhost:8080/healthcheck`)
 - Enjoy

## Data output
It is a JSON formated output with the output off all scripts inside.

For example, if you created a script called `testing_script.sh` doing this:
```bash
#!/bin/bash

echo "This is a test"
```

The output would be:
```JSON
{
    "testing_script_CODE": 0,
    "testing_script_STDOUT": "This is a test",
    "testing_script_STDERR": ""
}
```

By default, the output shall have an error field which is empty if there is no error, and some text to indicate the error

## Configure
All the configuration you need is inside the `.env` file.

Some sample(and usefull) scripts are provided in the `scripts_template` folder.

## Authors

 - [Tom Moulard](mailto:tom@moulard.org)
 - [Sebastien Goubeau](mailto:sebastien.goubeau@epita.fr)
