# marvel

*marvel* is a Go client library for the [Marvel](https://developer.marvel.com/)
Comic API.

## Very much a work in progress

Things will be bumpy at first; I am slightly undecided about this package's API.
The initial effort is to implement:

* ~~server-side authentication~~
* client-side authentication
* character service - *awaiting other services' data types*
* comic service
* creator service
* event service
* series service
* story service
* various walk functions and other helpers

## Testing

Running the tests will require your own [developer](https://developer.marvel.com/)
API keys. Set the environment variables `MARVEL_PUBLIC_KEY` and `MARVEL_PRIVATE_KEY`
appropriately. For example:

```
$ MARVEL_PUBLIC_KEY=abcd MARVEL_PRIVATE_KEY=1234 go test -v .
```

The tests utilize [go-vcr](https://github.com/dnaeon/go-vcr). The Marvel API's responses
are recorded to cassettes under the fixtures directory and replayed as necessary.
This lessens the chance of being rate limited when iteratively testing. The fixtures
directory is not tracked in the repository. Simply delete it to receive updated responses
from the live API service and record fresh cassettes.
