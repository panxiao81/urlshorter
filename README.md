# Yet Another Naive URL Shorter

Yet another naive URL shorter.

## build

Need Go environment

```shell
go build ./main
```

## Run

We have a pre-built Docker image

```shell
docker run -d -p 8080:8080 -v path/to/config.yaml:/config/config.yaml panxiao81/urlshorter
```

Or you can build binary by yourself.

You need a config file first.

The config file is being read at `config/config.yaml` by default.

You can use `-c` flag to change the location of the config file.

```shell
./urlshorter

# or

./urlshorter -c config.yaml
```

Then browse to `localhost:8080`.

## TODO

- [ ] Support SQL database
- [ ] Config hot reload
- [ ] Support memory cache
- [ ] A handier Web Interface