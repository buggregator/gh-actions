<p align="center">
    <img src="https://github.com/buggregator/gh-actions/assets/773481/a09810d4-aead-477f-86a2-01cab5f24904" width="400" alt="Buggregator">
    <h1 align="center">Buggregator GitHub Actions</h1>
    <br>
</p>

## General usage

This package allows us to reuse GitHub actions, which simplifies version management.

### Example of use of the action [PHPUnit](https://github.com/sebastianbergmann/phpunit)

```yml
on:
  push:
    branches:
      - master
      - '*.*'
  pull_request: null
  
name: phpunit

jobs:
  phpunit:
    uses: buggregator/gh-actions/.github/workflows/phpunit.yml@master
    with:
      os: >-
        ['ubuntu-latest']
      php: >-
        ['8.3']
```

### Example of use of the action [PSALM](https://github.com/vimeo/psalm)

```yml
on:
  push:
    branches:
      - master
      - '*.*'
  pull_request: null

name: static analysis

jobs:
  psalm:
    uses: buggregator/gh-actions/.github/workflows/psalm.yml@master
    with:
      os: >-
        ['ubuntu-latest']
      php: >-
        ['8.3']
```

```yml
on:
  push:
    branches:
      - master
      - '*.*'
  pull_request: null

name: coding standards

jobs:
  coding-standards:
    uses: buggregator/gh-actions/.github/workflows/cs.yml@master
    with:
      os: >-
        ['ubuntu-latest']
      php: >-
        ['8.3']
```
