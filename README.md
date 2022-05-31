# pair

## Description
This package provides generic KeyValue pair, along with couple of useful methods. It is dead simple as code and have 100% code coverage.

## Motivation
In my personal project I came across a situation, where I had 2 modules, which I didn't want to depend on each other. In module A I used "dependency injection", which I implemented as an interface that module B should (implicitly) fulfil. The methods of the interface returned list of key-value pairs. In that case using ad-hoc `string {key string; value string}` in both places couldn't work, becayse they are seen by the compiler as different types. Thus, in order not to introduce dependancy between them, I needed both the interface in module A and the implementation in module B to depend on the same struct defined somewhere in third module. It seemed excessive to me to introduce a third module only for such a generic thing in this specific project, so coming from the C# world, I decided to create a separate package that basically copies the functionality of C#'s KeyValuePair type.

## Usage

### Creating KeyValue pair

`kvp := pair.KeyValue[string, string]{Key: "key", Value: "value"}`

and then simply

`kvp.Key` will give you "key" and `kvp.Value` will give you "value", as one might expect

### Deconstructing

Following C# lead, there is `Deconstruct` method:

`key, value := kvp.Deconstruct()`

### String

`kvp.String()` will give you `{key value}`

## About me
Written by Victor Tashkov <vtashkov@gmail.com>
Published with MIT License.