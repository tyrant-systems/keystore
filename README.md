# key store

* * *

## about

a simple server configured to serve a `user/` directory. each subdirectory maps
to a realm user and contains a single file `keys`. the contents of which are
public ssh keys owned by that user.

example directory layout:

```
user/
    nate/
        keys
    jenna/
        keys
    ...
```

the `keys`file format is exactly the same as an ssh `authorized_keys` file. 
