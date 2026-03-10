# Test Library (test)

## Overview

`libtest` is a unit testing framework for Oak, providing assertions, test reporting, and test management capabilities.

## Import

```oak
test := import('test')
// Use constructors
suite := test.new('My Tests')
```

## Test Suite API

### `new(title)`

Creates a new test suite with the given title.

**Methods:**
- `eq(name, result, expect)` - Assert equality
- `approx(name, result, expect, epsilon?)` - Assert approximate equality
- `assert(name, result)` - Assert truthiness
- `skip(name, result, expect)` - Skip a test
- `reportFailed()` - Report only failed tests
- `report()` - Report all tests
- `exit()` - Exit with code 0 (pass) or 1 (fail)

```oak
{ new: new } := import('test')

suite := new('Math Tests')

suite.eq('addition', 2 + 2, 4)
suite.eq('subtraction', 5 - 3, 2)
suite.eq('multiplication', 3 * 4, 12)

suite.report()
suite.exit()
```

## Assertion Methods

### `eq(name, result, expect)`

Asserts that `result` equals `expect` using `=` comparison.

```oak
suite := new('Equality Tests')

suite.eq('string equality', 'hello', 'hello')
suite.eq('number equality', 42, 42)
suite.eq('boolean equality', true, true)
suite.eq('list equality', [1, 2, 3], [1, 2, 3])
suite.eq('object equality', {a: 1, b: 2}, {a: 1, b: 2})

// This will fail
suite.eq('failing test', 2 + 2, 5)
```

### `approx(name, result, expect, epsilon?)`

Asserts that `result` is approximately equal to `expect` within `epsilon`. Default epsilon: `0.00000001`.

Works recursively for lists and objects.

```oak
suite := new('Approximate Tests')

suite.approx('floating point', 0.1 + 0.2, 0.3)
suite.approx('with custom epsilon', 3.14159, 3.14, 0.01)

// Works with lists
suite.approx('list of floats', [1.0001, 2.0001], [1.0, 2.0], 0.001)

// Works with objects
suite.approx('object values', {x: 1.0001, y: 2.0001}, {x: 1.0, y: 2.0}, 0.001)
```

### `assert(name, result)`

Asserts that `result` is truthy (equivalent to `eq(name, result, true)`).

```oak
suite := new('Boolean Tests')

suite.assert('should be true', 5 > 3)
suite.assert('should exist', fileExists('config.json'))
suite.assert('non-empty string', len('hello') > 0)

// These will fail
suite.assert('false value', false)
suite.assert('null value', ?)
```

### `skip(name, result, expect)`

Marks a test as skipped. The test is not run but is counted in the report.

```oak
suite := new('Tests with Skips')

suite.eq('passing test', 1 + 1, 2)
suite.skip('not implemented yet', someFunc(), expectedValue)
suite.eq('another passing test', 'a' + 'b', 'ab')
```

## Reporting

### `reportFailed()`

Prints only the tests that failed, plus an aggregate summary.

```oak
suite := new('Some Tests')

suite.eq('test 1', 1, 1) // Pass
suite.eq('test 2', 2, 3) // Fail
suite.eq('test 3', 3, 3) // Pass
suite.eq('test 4', 4, 5) // Fail

suite.reportFailed()
// Only prints test 2 and test 4
// Output:
// Failed Some Tests tests:
//  ✘ test 2
//    expected: 3
//      result: 2
//  ✘ test 4
//    expected: 5
//      result: 4
// 2 / 4 tests passed.
```

### `report()`

Prints all tests with their pass/fail status.

```oak
suite := new('All Tests')

suite.eq('test 1', 1, 1) // Pass
suite.eq('test 2', 2, 3) // Fail

suite.report()
// Output:
// All Tests tests:
//  ✔ test 1
//  ✘ test 2
//    expected: 3
//      result: 2
// 1 / 2 tests passed.
```

### `exit()`

Exits the program with exit code 0 if all tests passed, or 1 if any failed.

```oak
suite := new('Tests')

suite.eq('test', 1, 1)

suite.report()
suite.exit() // Exits with code 0 (success)
```

## Complete Example

```oak
{ new: new } := import('test')

suite := new('String Library Tests')

// Test string concatenation
suite.eq('concat two strings', 'hello' + ' world', 'hello world')

// Test string length
suite.eq('string length', len('hello'), 5)

// Test string slicing
{ slice: slice } := import('str')
suite.eq('slice string', slice('hello', 0, 3), 'hel')

// Test string contains
{ contains?: contains? } := import('str')
suite.assert('contains substring', contains?('hello world', 'world'))

// Test case conversion
{ upper: upper, lower: lower } := import('str')
suite.eq('uppercase', upper('hello'), 'HELLO')
suite.eq('lowercase', lower('HELLO'), 'hello')

// Report and exit
suite.report()
suite.exit()
```

## Testing Functions

```oak
{ new: new } := import('test')

// Function to test
fn add(a, b) a + b
fn multiply(a, b) a * b

suite := new('Math Functions')

suite.eq('add positive numbers', add(2, 3), 5)
suite.eq('add negative numbers', add(-2, -3), -5)
suite.eq('add zero', add(5, 0), 5)

suite.eq('multiply positive', multiply(3, 4), 12)
suite.eq('multiply by zero', multiply(5, 0), 0)
suite.eq('multiply negative', multiply(-2, 3), -6)

suite.report()
suite.exit()
```

## Testing with Setup/Teardown

```oak
{ new: new } := import('test')

suite := new('Database Tests')

// Setup
db := connectToDatabase()

// Tests
suite.assert('connection established', db != ?)
suite.eq('insert record', db.insert({id: 1, name: 'Alice'}), true)
suite.eq('fetch record', db.get(1).name, 'Alice')

// Teardown
db.close()

suite.report()
suite.exit()
```

## Grouped Tests

```oak
{ new: new } := import('test')

// Create separate suites for different modules
stringSuite := new('String Tests')
mathSuite := new('Math Tests')

stringSuite.eq('concat', 'a' + 'b', 'ab')
stringSuite.eq('length', len('hi'), 2)

mathSuite.eq('add', 1 + 1, 2)
mathSuite.eq('multiply', 2 * 3, 6)

// Report each suite
stringSuite.report()
mathSuite.report()

// Exit based on combined results
if stringsSuite.hasFailures() | mathSuite.hasFailures() -> exit(1)
```

## Output Formatting

Test output uses ANSI color codes:
- ✔ Green for passing tests
- ✘ Red for failing tests

Failed tests show:
- Expected value (indented debug output)
- Actual result (indented debug output)

## Tips and Best Practices

### Descriptive Test Names

```oak
// Good
suite.eq('parseJSON handles empty object', parse('{}'), {})

// Bad
suite.eq('test1', parse('{}'), {})
```

### Test One Thing Per Assertion

```oak
// Good - separate tests
suite.eq('user has correct name', user.name, 'Alice')
suite.eq('user has correct age', user.age, 30)

// Bad - combined
suite.assert('user is valid', user.name = 'Alice' & user.age = 30)
```

### Use approx for Floating Point

```oak
// Good
suite.approx('calculate average', average([1, 2, 3]), 2.0, 0.01)

// Risky - may fail due to precision
suite.eq('calculate average', average([1, 2, 3]), 2.0)
```

### Test Edge Cases

```oak
suite := new('Edge Cases')

suite.eq('empty list', len([]), 0)
suite.eq('null value', stringify(?), 'null')
suite.eq('zero', abs(0), 0)
suite.eq('negative', abs(-5), 5)
```

## Limitations

- No async test support
- No test fixtures/hooks (before/after)
- No test organization (suites within suites)
- No test timing/performance metrics
- No test coverage reporting
- No test filtering/selection
- No parallel test execution
- No automatic test discovery

## See Also

- `debug.inspect()` - For debugging test failures
- Oak `exit()` builtin - For programmatic exit codes
