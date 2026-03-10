# Sort Library (sort)

## Overview

`libsort` implements efficient in-place list sorting using the quicksort algorithm with Hoare partitioning. Provides both mutating and non-mutating variants.

## Import

```oak
sort := import('sort')
// or destructure specific functions
{ sort: sort, sort!: sort! } := import('sort')
```

## Functions

### `sort!(xs, pred?)`

Sorts the list `xs` **in-place** (mutates the original list) by each item's `pred` value using quicksort. Returns the sorted list.

**Parameters:**
- `xs` - List to sort
- `pred` - Optional predicate function or key (default: sorts by item value)

**Mutates the original list** for efficiency.

```oak
{ sort!: sort! } := import('sort')

// Sort numbers in-place
numbers := [3, 1, 4, 1, 5, 9, 2, 6]
sort!(numbers)
// numbers is now [1, 1, 2, 3, 4, 5, 6, 9]

// Sort strings
words := ['banana', 'apple', 'cherry']
sort!(words)
// words is now ['apple', 'banana', 'cherry']

// Sort by predicate
users := [
    { name: 'Bob', age: 30 }
    { name: 'Alice', age: 25 }
    { name: 'Charlie', age: 35 }
]
sort!(users, fn(user) user.age)
// Sorted by age: Alice (25), Bob (30), Charlie (35)

// Sort by property name (shorthand)
sort!(users, 'name')
// Sorted alphabetically by name
```

### `sort(xs, pred?)`

Returns a **sorted copy** of `xs` without mutating the original. Internally calls `clone()` then `sort!()`.

**Parameters:**
- `xs` - List to sort
- `pred` - Optional predicate function or key

**Does not mutate the original list**, returns a new sorted list.

```oak
{ sort: sort } := import('sort')

original := [3, 1, 4, 1, 5, 9]
sorted := sort(original)
// sorted = [1, 1, 3, 4, 5, 9]
// original = [3, 1, 4, 1, 5, 9] (unchanged)

// Sort by predicate
data := [
    { id: 3, value: 'c' }
    { id: 1, value: 'a' }
    { id: 2, value: 'b' }
]
sorted := sort(data, fn(item) item.id)
// Sorted copy by id
// original data unchanged
```

## Predicate Functions

The `pred` parameter can be:

1. **Function** - Receives each item, returns comparable value
2. **String/Atom** - Property name to sort by (uses `std._asPredicate()`)
3. **Omitted/null** - Sorts by item value directly

```oak
{ sort: sort } := import('sort')

data := [
    { name: 'Bob', score: 85 }
    { name: 'Alice', score: 92 }
    { name: 'Charlie', score: 78 }
]

// By function
sort(data, fn(item) item.score)

// By property name (string)
sort(data, 'name')

// By property name (atom)
sort(data, :score)

// Direct value (numbers, strings, etc.)
sort([3, 1, 4, 1, 5, 9])
```

## Examples

### Sorting Numbers

```oak
{ sort: sort } := import('sort')

numbers := [42, 17, 93, 8, 51, 26]
ascending := sort(numbers)
// => [8, 17, 26, 42, 51, 93]

// Reverse sort by negating
descending := sort(numbers, fn(n) -n)
// => [93, 51, 42, 26, 17, 8]
```

### Sorting Strings

```oak
{ sort: sort } := import('sort')

words := ['zebra', 'apple', 'mango', 'banana']
sorted := sort(words)
// => ['apple', 'banana', 'mango', 'zebra']

// Case-insensitive sort
{ lower: lower } := import('str')
sorted := sort(words, lower)
```

### Sorting Objects by Property

```oak
{ sort: sort } := import('sort')

products := [
    { name: 'Widget', price: 19.99 }
    { name: 'Gadget', price: 9.99 }
    { name: 'Tool', price: 29.99 }
]

// Sort by price (ascending)
byPrice := sort(products, 'price')
// => Gadget ($9.99), Widget ($19.99), Tool ($29.99)

// Sort by name (alphabetical)
byName := sort(products, 'name')
// => Gadget, Tool, Widget
```

### Complex Sorting Criteria

```oak
{ sort: sort } := import('sort')

students := [
    { name: 'Alice', grade: 'A', score: 95 }
    { name: 'Bob', grade: 'B', score: 85 }
    { name: 'Charlie', grade: 'A', score: 92 }
]

// Sort by grade, then score within grade
// (Two separate sorts: stable for strings)
sorted := students |>
    sort(fn(s) s.score) |>
    sort(fn(s) s.grade)
```

### Sorting by Multiple Criteria

```oak
{ sort: sort } := import('sort')

// Custom composite key
records := [
    { category: 'B', priority: 1 }
    { category: 'A', priority: 2 }
    { category: 'A', priority: 1 }
]

// Sort by category, then priority
sorted := sort(records, fn(r) r.category + string(r.priority))
```

### Sorting Lists of Lists

```oak
{ sort: sort } := import('sort')

pairs := [[3, 'c'], [1, 'a'], [2, 'b']]

// Sort by first element
sorted := sort(pairs, fn(pair) pair.0)
// => [[1, 'a'], [2, 'b'], [3, 'c']]

// Sort by second element
sorted := sort(pairs, fn(pair) pair.1)
// => [[1, 'a'], [2, 'b'], [3, 'c']]
```

### In-Place vs. Copy Performance

```oak
{ sort: sort, sort!: sort! } := import('sort')

largeList := range(10000) |> map(fn(_) rand())

// Mutating (faster, modifies original)
start := time()
sort!(largeList)
mutatingTime := time() - start

// Non-mutating (slower, creates copy)
largeList2 := clone(largeList)
start := time()
sorted := sort(largeList2)
copyingTime := time() - start

// Mutating is faster as it avoids the clone()
```

### Top-N Elements

```oak
{ sort: sort } := import('sort')
{ take: take } := import('std')

data := [42, 17, 93, 8, 51, 26, 35, 19]

// Get top 3
topThree := data |>
    sort(fn(n) -n) |>  // Sort descending
    take(3)
// => [93, 51, 42]
```

### Grouping After Sorting

```oak
{ sort: sort } := import('sort')
{ partition: partition } := import('std')

data := [
    { type: 'A', value: 3 }
    { type: 'B', value: 1 }
    { type: 'A', value: 2 }
    { type: 'B', value: 4 }
]

// Sort by type, then partition
sorted := sort(data, 'type')
grouped := partition(sorted, fn(item) item.type)
// => [[A, A], [B, B]]
```

### Stable Sort Simulation

```oak
{ sort: sort } := import('sort')

// Quicksort is not stable, but can achieve stability
// by including original index in sort key

items := ['b', 'a', 'b', 'c', 'a']

// Add indices
indexed := items |> map(fn(item, i) [item, i])

// Sort by value, but preserve order for equal items
sorted := sort(indexed, fn(pair) pair.0)

// Extract values
result := sorted |> map(fn(pair) pair.0)
```

## Algorithm Details

- **Algorithm:** Quicksort with Hoare partitioning
- **Time Complexity:** 
  - Average: O(n log n)
  - Worst: O(n²) (rare with random data)
- **Space Complexity:** O(log n) stack space for recursion
- **Stability:** Not stable (equal elements may be reordered)
- **In-place:** Yes (`sort!`) or creates copy (`sort`)

## Performance Characteristics

```oak
{ sort!: sort!, sort: sort } := import('sort')

// Small lists (n < 100)
// - Overhead of quicksort is minimal
// - Both variants perform similarly

// Medium lists (100 < n < 10,000)
// - Quicksort efficiency becomes apparent
// - sort! much faster (no clone)

// Large lists (n > 10,000)
// - sort! recommended for performance
// - clone() overhead significant in sort()
```

## Comparison with Other Approaches

```oak
// Manual bubble sort (slow)
fn bubbleSort(list) {
    n := len(list)
    fn pass(i) if i < n {
        true -> {
            // bubble largest to end
            pass(i + 1)
        }
        _ -> list
    }
    pass(0)
}

// Library quicksort (fast)
{ sort!: sort! } := import('sort')
sort!(list)
```

## Limitations

- Not stable—equal elements may reorder
- No custom comparison function (only key extraction)
- No reverse flag (must negate values)
- No partial sorting (must sort entire list)
- No guaranteed O(n log n) worst-case (use merge sort for that)
- Mutating version (`sort!`) offers no protection against concurrent modification

## When to Use Each Variant

### Use `sort!()` when:
- Performance is critical
- List is already a copy
- You don't need the original order
- Working with large lists

### Use `sort()` when:
- Need to preserve original list
- Working with immutable data patterns
- List is small (< 100 elements)
- Copying overhead is acceptable

## See Also

- `std` library - For `clone()`, `map()`, and list operations
- Oak comparison operators: `<`, `>`, `=` for sorting
