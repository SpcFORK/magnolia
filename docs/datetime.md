# DateTime Library (datetime)

## Overview

`libdatetime` provides utilities for working with dates, times, and UNIX timestamps. It handles date arithmetic, timezone offsets, and ISO 8601 formatting/parsing.

**Scope:**
- Correct for dates from 0001-01-01 (1 CE) forward
- Works with UNIX timestamps (positive and negative)
- Handles leap years and 400-year calendar cycles
- Does not handle millisecond resolution internally (except in format/parse)
- Timezone complexity pushed to call sites

## Import

```oak
datetime := import('datetime')
// or destructure specific functions
{ describe: describe, format: format, parse: parse, timestamp: timestamp } := import('datetime')
```

## Constants

### `SecondsPerDay`
Number of seconds in a day: `86400`

### `DaysPer4Years`, `DaysPer100Years`, `DaysPer400Years`
Days in various calendar cycles accounting for leap years

### `DaysBeforeMonth`
List of cumulative days before each month (1-indexed, January = 1)

## Functions

### `leap?(year)`

Returns whether a calendar year is a leap year.

**Leap Year Rules:**
- Divisible by 4: leap year
- Except divisible by 100: not a leap year  
- Except divisible by 400: leap year

```oak
{ leap?: leap? } := import('datetime')

leap?(2024) // => true (divisible by 4)
leap?(2000) // => true (divisible by 400)
leap?(1900) // => false (divisible by 100 but not 400)
leap?(2023) // => false (not divisible by 4)
```

### `describe(timestamp)`

Converts a UNIX timestamp to a date/time description object.

**Returns:**
```oak
{
    year: 2024
    month: 3        // 1-12 (January = 1)
    day: 9          // 1-31
    hour: 14        // 0-23
    minute: 30      // 0-59
    second: 45      // 0-59
}
```

```oak
{ describe: describe } := import('datetime')

desc := describe(time())
println('Year: ' + string(desc.year))
println('Month: ' + string(desc.month))
println('Day: ' + string(desc.day))

// Describe a specific timestamp
desc := describe(0)  // Unix epoch
// => { year: 1970, month: 1, day: 1, hour: 0, minute: 0, second: 0 }
```

### `timestamp(description)`

Converts a date/time description object back to a UNIX timestamp.

**Parameters:**
```oak
{
    year: 2024
    month: 3
    day: 9
    hour: 14
    minute: 30
    second: 45
}
```

```oak
{ timestamp: timestamp } := import('datetime')

ts := timestamp({
    year: 2024
    month: 1
    day: 1
    hour: 0
    minute: 0
    second: 0
})
// => 1704067200 (Unix timestamp for 2024-01-01T00:00:00Z)

// Round-trip conversion
original := time()
desc := describe(original)
restored := timestamp(desc)
// original ≈ restored (within precision)
```

### `format(timestamp, tzOffset?)`

Formats a timestamp as an ISO 8601 date-time string.

**Parameters:**
- `timestamp` - UNIX timestamp
- `tzOffset` - Timezone offset from UTC in minutes (default: 0 = UTC)

**Format:** `YYYY-MM-DDTHH:MM:SS[.mmm][Z|±HH:MM]`

```oak
{ format: format } := import('datetime')

// UTC (default)
format(time())
// => '2024-03-09T14:30:45Z'

// With milliseconds
format(time())
// => '2024-03-09T14:30:45.123Z'

// With timezone offset (+5 hours = 300 minutes)
format(time(), 300)
// => '2024-03-09T19:30:45+05:00'

// With negative offset (-8 hours = -480 minutes)  
format(time(), -480)
// => '2024-03-09T06:30:45-08:00'

// Specific timestamp
format(0, 0)
// => '1970-01-01T00:00:00Z'
```

### `parse(isoString)`

Parses an ISO 8601 date-time string and returns a description object with timezone offset. Returns `?` on invalid input.

**Returns:**
```oak
{
    year: 2024
    month: 3
    day: 9
    hour: 14
    minute: 30
    second: 45
    tzOffset: 0     // Timezone offset in minutes
}
```

```oak
{ parse: parse } := import('datetime')

// UTC time
desc := parse('2024-03-09T14:30:45Z')
// => { year: 2024, month: 3, day: 9, hour: 14, minute: 30, second: 45, tzOffset: 0 }

// With timezone
desc := parse('2024-03-09T19:30:45+05:00')
// => { ..., hour: 19, tzOffset: 300 }

// With milliseconds
desc := parse('2024-03-09T14:30:45.123Z')
// => { ..., second: 45.123, ... }

// Invalid input
parse('invalid') // => ?
parse('2024-13-01T00:00:00Z') // => ? (month 13 invalid)
```

## Examples

### Current Date/Time

```oak
{ describe: describe, format: format } := import('datetime')

now := time()
desc := describe(now)

println('Current date: ' + string(desc.year) + '-' + string(desc.month) + '-' + string(desc.day))
println('Current time: ' + string(desc.hour) + ':' + string(desc.minute))

// Or use format for ISO 8601
println('ISO format: ' + format(now))
// => ISO format: 2024-03-09T14:30:45Z
```

### Date Arithmetic

```oak
{ describe: describe, timestamp: timestamp } := import('datetime')

// Add 7 days to current time
now := time()
sevenDaysLater := now + (7 * 86400)  // 7 days * 86400 seconds/day

desc := describe(sevenDaysLater)
println('7 days from now: ' + string(desc.year) + '-' + string(desc.month) + '-' + string(desc.day))

// Add 1 month (approximately)
desc := describe(now)
desc.month := desc.month + 1
if desc.month > 12 -> {
    desc.month <- desc.month - 12
    desc.year <- desc.year + 1
}
futureTs := timestamp(desc)
```

### Timezone Conversion

```oak
{ format: format, parse: parse } := import('datetime')

// Current time in different timezones
utcTime := time()

println('UTC: ' + format(utcTime, 0))
println('New York (EST): ' + format(utcTime, -300))     // UTC-5
println('London: ' + format(utcTime, 0))                // UTC+0
println('Tokyo: ' + format(utcTime, 540))               // UTC+9
println('Los Angeles (PST): ' + format(utcTime, -480)) // UTC-8
```

### Parsing and Converting Timestamps

```oak
{ parse: parse, timestamp: timestamp } := import('datetime')

dateStr := '2024-12-25T00:00:00Z'
desc := parse(dateStr)

if desc != ? -> {
    ts := timestamp(desc)
    println('Christmas 2024 timestamp: ' + string(ts))
    
    // Calculate days until Christmas
    now := time()
    daysUntil := int((ts - now) / 86400)
    println('Days until Christmas: ' + string(daysUntil))
}
```

### Age Calculation

```oak
{ describe: describe, timestamp: timestamp } := import('datetime')

fn calculateAge(birthYear, birthMonth, birthDay) {
    birth := timestamp({
        year: birthYear
        month: birthMonth
        day: birthDay
        hour: 0
        minute: 0
        second: 0
    })
    
    now := time()
    ageInSeconds := now - birth
    ageInYears := int(ageInSeconds / (86400 * 365.25))
    ageInYears
}

age := calculateAge(1990, 5, 15)
println('Age: ' + string(age))
```

### Log Timestamps

```oak
{ format: format } := import('datetime')

fn log(level, message) {
    timestamp := format(time())
    println('[' + timestamp + '] ' + level + ': ' + message)
}

log('INFO', 'Application started')
log('WARN', 'Low memory')
log('ERROR', 'Connection failed')
```

### Date Range Generation

```oak
{ describe: describe, timestamp: timestamp } := import('datetime')
std := import('std')

fn dateRange(startTs, endTs) {
    dates := []
    current := startTs
    
    fn sub if current <= endTs {
        true -> {
            dates << current
            sub() where current := current + 86400  // Add one day
        }
        _ -> dates
    }
    
    sub()
}

start := timestamp({ year: 2024, month: 1, day: 1, hour: 0, minute: 0, second: 0 })
end := timestamp({ year: 2024, month: 1, day: 7, hour: 0, minute: 0, second: 0 })

dates := dateRange(start, end)
println('Generated ' + string(len(dates)) + ' dates')
```

### Duration Formatting

```oak
{ describe: describe } := import('datetime')

fn formatDuration(seconds) {
    days := int(seconds / 86400)
    hours := int((seconds % 86400) / 3600)
    mins := int((seconds % 3600) / 60)
    secs := int(seconds % 60)
    
    if days > 0 -> string(days) + 'd ' + string(hours) + 'h'
    if hours > 0 -> string(hours) + 'h ' + string(mins) + 'm'
    if mins > 0 -> string(mins) + 'm ' + string(secs) + 's'
    string(secs) + 's'
}

startTime := time()
// ... do work ...
endTime := time()

println('Duration: ' + formatDuration(endTime - startTime))
```

## Implementation Notes

- Uses Gregorian calendar with 400-year cycles
- Leap day is February 29
- Months are 1-indexed (January = 1, December = 12)
- Timezone offset in `format()` adjusts the displayed time
- Timezone offset in `parse()` is returned but not applied to values
- Milliseconds are supported in format/parse but not in internal representation
- Negative timestamps represent dates before Unix epoch (1970-01-01)

## Leap Year Behavior

```oak
{ leap?: leap?, describe: describe, timestamp: timestamp } := import('datetime')

// Leap day timestamp
leapDay := timestamp({ year: 2024, month: 2, day: 29, hour: 0, minute: 0, second: 0 })
println(leap?(2024)) // => true

// Non-leap year - February only has 28 days
// Attempting Feb 29 on non-leap year may produce unexpected results
```

## Limitations

- No Duration type—use raw seconds
- No built-in timezone database—offsets must be provided
- No daylight saving time handling
- No calendar-aware date arithmetic (e.g., "add 1 month" can be ambiguous)
- Millisecond support only in format/parse, not in internal calculations
- No locale-specific formatting
- No relative time calculations (e.g., "3 days ago")

## See Also

- Oak built-in `time()` - Get current UNIX timestamp
- `fmt` library - For custom formatting
- `str` library - For string manipulation
