# üWave Sigil

Adaptation of Cupcake's [Sigil](https://github.com/cupcake/sigil), for the üWave
instance for [WE ♥ KPOP](https://welovekpop.club).

## Changes

 - Removed SVG icon generation because üWave does not need it
 - Set default icon size to 120px because that's somewhat closer to üWave's
   avatar size
 - Add a `run.sh` script to make running it with `pm2` easier
 - Add files for serverless Now deployment

## Running it

Clone the repository, then:

- Locally, use `./run.sh`
- For production, use `./build.sh` and then run the `sigil` binary on your server
- Serverlessly, run `now`

## License

See [LICENSE](./LICENSE)
