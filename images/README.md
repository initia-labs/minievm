# Specialized images for minitia

## `private`

Use a GitHub PAT (Personal Access Token) to access private git repos releases and git trees.
To be used for releases until repos become public.

Build like:

``` bash
docker build \
    -t minitiad \
    --build-arg LIBWASMVM_VERSION=v1.5.0 \
    --build-arg GITHUB_ACCESS_TOKEN=$PAT \
    -f images/private/Dockerfile \
    .
```
