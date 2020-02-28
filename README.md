Running `go mod tidy` once gives different results than running in twice. 
This repository reproduces this issue.

Run following:
```
cd hello
for id in a b c d; do cp go.mod $id; go mod tidy; done
shasum a b c d
```

The output:
```
f99722d4551542f425478d1606680dc030e20de7  a
ca3670789242a94a99aeede3d4593ab6fa8579eb  b
4865db2ffbdd6379e06f30c69574fade80269bef  c
4865db2ffbdd6379e06f30c69574fade80269bef  d
```

a != b != c

It's broken in Go 1.13, works fine in 1.14
