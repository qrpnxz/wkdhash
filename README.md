# wkdhash

Command wkdhash takes a mail user as its first argument and returns its
sha1 encoded in base32 for use in a [Web Key Directory][1].

Example:
```
 $ wkdhash test
iffe93qcsgp4c8ncbb378rxjo6cn9q6u
 $ wkdhash TEST
iffe93qcsgp4c8ncbb378rxjo6cn9q6u
 $ wkdhash lest
4ykxa7wnwasr7goc634k6w1ej1ib8zge
```

## Copyright
See [COPYRIGHT.md](COPYRIGHT.md).

## Contributing
See [CONTRIBUTING.md](CONTRIBUTING.md).

[1]: <https://tools.ietf.org/html/draft-koch-openpgp-webkey-service-11>
"OpenPGP Web Key Directory draft standard"