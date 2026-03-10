# Standard Library (std)

## Overview

`libstd` is the core standard library for Oak, providing essential functions for working with Oak values, iterators, and control flow. It defines the fundamental building blocks for functional programming patterns in Oak.

## Import

```oak
std := import('std')
// or destructure specific functions
{ map: map, filter: filter, reduce: reduce } := import('std')
```

## Core Functions

### `identity(x)`

Returns its first argument unchanged.

```oak
identity(5) // => 5
identity('hello') // => 'hello'
```

### `is(x)`

Returns a predicate function that checks if its argument equals `x`.

```oak
isThree := is(3)
isThree(3) // => true
isThree(5) // => false
```

### `constantly(x)`

Returns a function that always returns `x`, regardless of arguments.

```oak
alwaysFive := constantly(5)
alwaysFive() // => 5
alwaysFive(1, 2, 3) // => 5
```

### `default(x, base)`

Returns `x` if `x` is not null, otherwise returns `base`. Useful for optional function arguments with default values.

```oak
fn greet(name) {
    name := default(name, 'World')
    'Hello, ' + name
}

greet('Alice') // => 'Hello, Alice'
greet() // => 'Hello, World'
```

## Number Conversion

### `toHex(n)`

Converts a number to its hexadecimal string representation. Fails for negative values.

```oak
toHex(255) // => 'ff'
toHex(16) // => '10'
```

### `fromHex(s)`

Parses a hexadecimal string to an integer. Returns `?` if input is invalid.

```oak
fromHex('ff') // => 255
fromHex('10') // => 16
fromHex('invalid') // => ?
```

## Utility Functions

### `clamp(min, max, n, m)`

Constrains two values `n` and `m` to the range `[min, max]`, ensuring `n <= m`.

```oak
clamp(0, 10, 5, 8) // => [5, 8]
clamp(0, 10, -5, 15) // => [0, 10]
```

### `slice(xs, min, max)`

Returns a copy of a substring or sublist from index `min` (inclusive) to `max` (exclusive). Both `min` and `max` are optional (default to 0 and `len(xs)`).

```oak
slice([1, 2, 3, 4, 5], 1, 4) // => [2, 3, 4]
slice('hello', 1, 4) // => 'ell'
slice([1, 2, 3], 1) // => [2, 3]
```

### `clone(x)`

Creates a shallow copy of any Oak value.

```oak
original := [1, 2, 3]
copy := clone(original)
copy << 4 // original remains [1, 2, 3]
```

## List Generation

### `range(start, end, step)`

Returns a list of numbers from `start` to `end` (exclusive), incrementing by `step`. Defaults: `step = 1`, `start = 0`.

```oak
range(5) // => [0, 1, 2, 3, 4]
range(2, 5) // => [2, 3, 4]
range(0, 10, 2) // => [0, 2, 4, 6, 8]
range(10, 0, -2) // => [10, 8, 6, 4, 2]
```

### `reverse(xs)`

Reverses the order of elements in an iterable, producing a copy.

```oak
reverse([1, 2, 3]) // => [3, 2, 1]
reverse('hello') // => 'olleh'
```

## Functional Iterators

### `map(xs, f)`

Transforms each element of `xs` using function `f`. The function receives `(element, index)`.

```oak
map([1, 2, 3], fn(x) x * 2) // => [2, 4, 6]
map(['a', 'b'], fn(x, i) x + string(i)) // => ['a0', 'b1']

// If f is a string/atom/int, it acts as a property accessor
map([{a: 1}, {a: 2}], 'a') // => [1, 2]
```

### `each(xs, f)`

Calls function `f` for each element. The function receives `(element, index)`. Returns `?`.

```oak
each([1, 2, 3], fn(x) println(x))
// Prints: 1, 2, 3
```

### `filter(xs, f)`

Returns only elements where predicate `f` returns true. Function receives `(element, index)`.

```oak
filter([1, 2, 3, 4], fn(x) x % 2 = 0) // => [2, 4]
filter([{age: 20}, {age: 30}], fn(p) p.age > 25) // => [{age: 30}]
```

### `exclude(xs, f)`

Returns only elements where predicate `f` returns false. Opposite of `filter`.

```oak
exclude([1, 2, 3, 4], fn(x) x % 2 = 0) // => [1, 3]
```

### `separate(xs, f)`

Divides `xs` into two lists `[is, isnt]` based on predicate `f`.

```oak
separate([1, 2, 3, 4], fn(x) x % 2 = 0)
// => [[2, 4], [1, 3]]
```

### `reduce(xs, seed, f)`

Accumulates elements starting with `seed`. The reducer receives `(accumulator, element, index)`.

```oak
reduce([1, 2, 3], 0, fn(acc, x) acc + x) // => 6
reduce(['a', 'b', 'c'], '', fn(acc, x) acc + x) // => 'abc'
```

### `flatten(xs)`

Flattens a list of lists by one level.

```oak
flatten([[1, 2], [3, 4], [5]]) // => [1, 2, 3, 4, 5]
```

### `compact(xs)`

Removes all null (`?`) elements from a list.

```oak
compact([1, ?, 2, ?, 3]) // => [1, 2, 3]
```

## Boolean Operations

### `some(xs, pred)`

Returns true if at least one element satisfies the predicate (or is truthy if no predicate given).

```oak
some([false, true, false]) // => true
some([1, 2, 3], fn(x) x > 2) // => true
some([1, 2, 3], fn(x) x > 5) // => false
```

### `every(xs, pred)`

Returns true if all elements satisfy the predicate (or are truthy if no predicate given).

```oak
every([true, true, true]) // => true
every([1, 2, 3], fn(x) x > 0) // => true
every([1, 2, 3], fn(x) x > 2) // => false
```

## List Manipulation

### `append(xs, ys)`

Joins two iterables together, **mutating** the first argument.

```oak
a := [1, 2]
append(a, [3, 4]) // => [1, 2, 3, 4]
// a is now [1, 2, 3, 4]
```

### `join(xs, ys)`

Joins two iterables without mutating either value.

```oak
a := [1, 2]
b := join(a, [3, 4]) // => [1, 2, 3, 4]
// a remains [1, 2]
```

### `zip(xs, ys, zipper)`

Pairs up elements from two iterables. Default zipper creates 2-element lists.

```oak
zip([1, 2, 3], [4, 5, 6]) // => [[1, 4], [2, 5], [3, 6]]
zip([1, 2, 3], [4, 5, 6], fn(a, b) a * b) // => [4, 10, 18]
```

### `partition(xs, by)`

Divides `xs` into partitions. If `by` is an integer, partitions have that size. If `by` is a function, partitions change when the function result changes.

```oak
partition([1, 2, 3, 4, 5, 6], 2) // => [[1, 2], [3, 4], [5, 6]]
partition([1, 1, 2, 2, 3], fn(x) x) // => [[1, 1], [2, 2], [3]]
```

### `uniq(xs, pred)`

Removes consecutive duplicate elements. For global uniqueness, sort first.

```oak
uniq([1, 1, 2, 3, 3, 3, 4]) // => [1, 2, 3, 4]
uniq([1, 2, 1], fn(x) x) // => [1, 2, 1] (not consecutive)
```

## Element Access

### `first(xs)`

Returns the first element of an iterable.

```oak
first([1, 2, 3]) // => 1
first('hello') // => 'h'
```

### `last(xs)`

Returns the last element of an iterable.

```oak
last([1, 2, 3]) // => 3
last('hello') // => 'o'
```

### `take(xs, n)`

Returns the first `n` elements.

```oak
take([1, 2, 3, 4, 5], 3) // => [1, 2, 3]
take('hello', 2) // => 'he'
```

### `takeLast(xs, n)`

Returns the last `n` elements.

```oak
takeLast([1, 2, 3, 4, 5], 3) // => [3, 4, 5]
takeLast('hello', 2) // => 'lo'
```

## Search Functions

### `find(xs, pred)`

Returns the index of the first element matching the predicate, or `-1` if not found.

```oak
find([10, 20, 30], fn(x) x > 15) // => 1
find([1, 2, 3], fn(x) x > 10) // => -1
```

### `rfind(xs, pred)`

Returns the index of the last element matching the predicate (searching backwards), or `-1`.

```oak
rfind([10, 20, 30, 20], fn(x) x = 20) // => 3
```

### `indexOf(xs, x)`

Returns the index of the first element equal to `x`, or `-1`.

```oak
indexOf([1, 2, 3, 2], 2) // => 1
indexOf([1, 2, 3], 5) // => -1
```

### `rindexOf(xs, x)`

Returns the index of the last element equal to `x` (searching backwards), or `-1`.

```oak
rindexOf([1, 2, 3, 2], 2) // => 3
```

### `contains?(xs, x)`

Returns true if the iterable contains an element equal to `x`.

```oak
contains?([1, 2, 3], 2) // => true
contains?([1, 2, 3], 5) // => false
```

## Object Functions

### `values(obj)`

Returns a list of all values in an object (order arbitrary).

```oak
values({a: 1, b: 2, c: 3}) // => [1, 2, 3] (order may vary)
```

### `entries(obj)`

Returns a list of `[key, value]` pairs (order arbitrary).

```oak
entries({a: 1, b: 2}) // => [['a', 1], ['b', 2]] (order may vary)
```

### `fromEntries(entries)`

Constructs an object from a list of `[key, value]` pairs.

```oak
fromEntries([['a', 1], ['b', 2]]) // => {a: 1, b: 2}
```

### `merge(os...)`

Merges multiple objects onto the first object, **mutating** it. Returns `?` if no objects given.

```oak
a := {x: 1}
merge(a, {y: 2}, {z: 3}) // => {x: 1, y: 2, z: 3}
// a is now {x: 1, y: 2, z: 3}
```

## Higher-Order Functions

### `once(f)`

Returns a wrapper function that ensures `f` is called exactly once, no matter how many times the wrapper is called.

```oak
init := once(fn() println('Initialized!'))
init() // Prints: Initialized!
init() // Does nothing
init() // Does nothing
```

## Notes

- Many functions accept a `pred` parameter that can be:
  - A function receiving `(element, index)`
  - A string/atom/int to access a property of each element
- Predicates are converted internally using `_asPredicate()`
- Most functions return copies rather than mutating the original, except `append()` and `merge()`
