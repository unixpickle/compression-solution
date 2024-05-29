# compression-solution

This is the solution to my [compression challenge](https://github.com/unixpickle/compression-challenge/).

To produce the data used for the challenge, `gunzip` the files in `inputs` and run:

```bash
$ mkdir files
$ for i in {1..7}; do go run . encode -input inputs/file${i} -output files/file${i}; done
```

The way the challenge is setup is that the input files are encrypted with AES. The underlying files themselves are simple text patterns that are easy to compress, but compressing them is difficult unless you can decrypt them first.

The compression procedure first decrypts the file, then `gzip`s the contents. The decompression procedure un-`gzip`s the contents, and then re-encrypts it.
