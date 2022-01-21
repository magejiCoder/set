## Generic Set

[![UnitTest](https://github.com/magejiCoder/set/actions/workflows/testing.yaml/badge.svg)](https://github.com/magejiCoder/set/actions/workflows/testing.yaml)
[![codecov](https://codecov.io/gh/magejiCoder/set/branch/master/graph/badge.svg?token=OS0J3ZALXS)](https://codecov.io/gh/magejiCoder/set)
[![CodeQL](https://github.com/magejiCoder/set/actions/workflows/codeql.yaml/badge.svg)](https://github.com/magejiCoder/set/actions/workflows/codeql.yaml)


Port [scylladb/go-set](https://github.com/scylladb/go-set) to Go Generics.

## Usage

The usage is almost identical to the original set package.

But for customized type , the original set package needs to **generate** new type package via `gen_set.sh`; For generics set package , just uses it as the builtin type:

```go

type CustomizedType struct {
	f string
}

 s:= New[CustomizedType]()
 s.Add(CustomizedType{"a"})
 if s.Has(CustomizedType{"a"}) {
	// do something
 }
```
