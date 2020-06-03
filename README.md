## Introduction
This module is used to calculate the number comment line and its percentage in a project.

## Usage

```shell script
calculator -dir=<dir>
# dir need to be full path
```

> WARNING:
> Since this project hasn't using parser, so you need to make sure the format of comments in your code is correct.
> You also need to check out there is not comment string in the string literal.
> For Python, currently we cannot deal with inline doc string (e.g., '''hello'')