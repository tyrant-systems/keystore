# key store

* * *

## about

a simple server configured to serve a `user/` directory. each subdirectory maps
to a realm user and contains a single file `keys`. the contents of which are
public ssh keys owned by that user. the `keys`file format is exactly the same as an ssh `authorized_keys` file.

## usage

create a `json` configuration file:

```
{
  "configuration": {
    "user_dir_abs_path": "/srv/realm-manifest/",
    "listen_addr": ":3030"
  }
}
```

directory layout under `/srv/realm-manifest/` in the example config:

```
user/
    nate/
        keys
    jenna/
        keys
    ...
```

build the server:

```
~$ go build main.go -o keysrv
```

run it:

```
~$ ./keysrv -config path/to/config.json
```

that's it.
