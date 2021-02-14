# üWave Sigil

Adaptation of Cupcake's [Sigil](https://github.com/cupcake/sigil), for use by üWave servers as a default avatar generator.

## Changes
 - Removed SVG icon generation because üWave does not use it
 - Set default icon size to 120px because that's somewhat closer to üWave's
   avatar size
 - Add a `run.sh` script to make development runs easier

## Running it
Clone the repository, then:

- Locally, use `./run.sh`
- For production, use `./build.sh` and then run the `sigil` binary on your server

## License

See [LICENSE](./LICENSE)
