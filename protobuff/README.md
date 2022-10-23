# Protocol Buffers

## References

* [proto3 language guide](https://developers.google.com/protocol-buffers/docs/proto3)
* [proto3 style guide](https://developers.google.com/protocol-buffers/docs/style)

## Example

```proto
syntax = "proto3";

message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
}

// this is a comment
```

> `proto3` uses C/C++ like comment style - //, /* */

## Field numbers

* Unique
* Used to identify your fields in the message binary format
* Should not be changed once your message type is in use
* Range 1 - 15 take one byte to encode, including field number and
  field's type, however field number in the range 16 - 2047 take 2
  bytes. This is why you should reserve numbers 1 - 15 for frequently
  occurring fields.

## Field rules

* Singular - default rule when nothing is specified
* Optional - you can check if the value is explicitly set
* Repeated - arrays. Order is preserved
* Map - key value pairs
* reserved - compiler will complain if you try to use these field ids

```proto
message Foo {
    reserved 2, 15, 9 to 11;
    reserved "foo", "bar";
}

// Note that you can't mix field names and field numbers in the same
// reserved statement.
```

> Note that if a scalar message field is set to its default, the value
> will not be serialized on the wire.
> — such a good design decision, especially when working with a
> programming language like Go. Think about it for a second.

## Enums

```proto
enum Corpus {
  CORPUS_UNSPECIFIED = 0; // the default enum value
  CORPUS_UNIVERSAL = 1;
  CORPUS_WEB = 2;
  CORPUS_IMAGES = 3;
  CORPUS_LOCAL = 4;
  CORPUS_NEWS = 5;
  CORPUS_PRODUCTS = 6;
  CORPUS_VIDEO = 7;
}

message SearchRequest {
  ...
  Corpus corpus = 4;
}
```

### Aliases in Enums

```proto
enum EnumAllowingAlias {
  // if this option is not specified, protoc will generate an error
  option allow_alias = true;

  // the default enum value
  EAA_UNSPECIFIED = 0;

  // EAA_RUNNING is an alias to EAA_STARTED
  EAA_STARTED = 1;
  EAA_RUNNING = 1;
}
```

### Reserved enum values

The same `reserved` keyword works for enums too.

```proto
enum Foo {
  reserved 2, 15, 9 to 11, 40 to max;
  reserved "BAR", "BLAH";
}
```

## Any

**TODO**: discover possible use-case

```proto
import "google/protobuf/any.proto";

message ErrorStatus {
  string message = 1;
  repeated google.protobuf.Any details = 2;
}
```

## OneOf

* Like unions in C
* If multiple values are set, the last set value will overwrite all
  previous ones.
* You can add fields of any type, **except map fields and repeated
  fields**.

```proto
message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;

  oneof engine_id {
      string google = 4;
      string bing = 5;
      string duck = 6;
  }
}
```

## Maps

```proto
message SearchRequest {
  ...
  map<string, int32> visited = 7;
  // map<key_type, value_type> field_name = field_number;
}
```

* `key_type` can be any scalar type except for floating point types and
  bytes.
* `value_type` can be any type other than map.
* Order of map items is not maintained.

## Import path

```
protoc -I=. -I=/usr/local/include/ --go_out=. search.proto
```

## Style guide

* Keep the line length to 80 characters.
* Use an indent of 2 spaces.

```vim
autocmd FileType proto set expandtab tabstop=2 softtabstop=2 shiftwidth=2
```

* Prefer the use of double quotes for strings.
* Files should be named `lower_snake_case.proto`
* Lower snake case for field names.
* Camel-Case (with initial capital) for message, service, rpc method
  names and enum type name.

```proto
message SearchResponse {
  repeated Result results = 1;
}

service SearchService {
  rpc Search(SearchRequest) returns (SearchResponse);
}

enum FooBar {
  FOO_BAR_UNSPECIFIED = 0;
  FOO_BAR_FIRST_VALUE = 1;
  FOO_BAR_SECOND_VALUE = 2;
}
```

* `CAPITALS_WITH_UNDERSCORES` for enum value names
* Pluralized names for repeated fields.
* Hierarchy convention:

```
syntax (1st non-commented line of file. Comments containing license
header and file overview may preceed)

package

imports(sorted)

options

everything else
```
