# marvel

*marvel* is a Go client library for the [Marvel](https://developer.marvel.com/)
Comic API.

## Very much a work in progress

Things will be bumpy at first; I am slightly undecided about this package's API.
The initial effort is to implement

* ~~server-side authentication~~
* client-side authentication
* ~~character service~~ *partially complete*
* comic service
* creator service
* event service
* series service
* story service
* various walk functions and other helpers

## Testing

Running the tests require your own [Marvel](https://developer.marvel.com/) developer
API keys. Set the environment variables `MARVEL_PUBLIC_KEY` and `MARVEL_PRIVATE_KEY`
appropriately. For example:

```
$ MARVEL_PUBLIC_KEY=abcd MARVEL_PRIVATE_KEY=1234 go test -v .
```

The tests utilize [go-vcr](https://github.com/dnaeon/go-vcr) and record to cassettes
under the fixtures directory, which is not tracked in the repository. Simply delete
the this directory to receive updated responses from the live Marvel API service.
