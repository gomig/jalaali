# Jalaali

Structural time package for jalaali (persian) calendar. This package support parse from string, json and time.

## Structures

There are three data structures for jalaali:

- **Jalaali** jalaali date with current system time zone.
- **JalaaliTehran** jalaali date with Asia/Tehran time zone.
- **JalaaliKabul** jalaali date with Asia/Kabul time zone.

### Methods

All structures has following methods for work with date.

#### IsNil

Check if date object is empty.

```go
jDate.IsNil()
```

#### SetTime

Set date from standard time object.

```go
jDate.SetTime(time.Now())
```

#### Parse

Parse date from jalaali date string.

**Note:** if date is invalid date object will be nil.

```go
jDate.Parse("1400-01-02")
```

#### Time

Get current date as standard time. This function return nil if date is empty.

```go
t := jDate.Time()
```

#### JTime

Get current date as jalaali date object. This function return nil if date is empty.

```go
j := jDate.JTime()
```

#### UTC

Get current date as standard time object with utc timezone. This function return nil if date is empty.

```go
utcT :=jDate.UTC()
```

#### Format

Format jalaali date. This function support standard go time formatting parameters.

```go
f := jDate.Format("2006-01-02")
```

## Helper Functions

### TehranTz

Get timezone for Asia/Tehran.

### KabulTz

Get timezone for Asia/Kabul.

### Parse

Parse jalaali date string as time. This function use system timezone.

### ParseForLocale

Parse jalaali date string as time for timezone.

### New

Create new jalaali structure from time object. This function use system timezone.

### NewTehran

Create new JalaaliTehran structure from time object. This function use Asia/Tehran timezone.

### NewKabul

Create new JalaaliKabul structure from time object. This function use Asia/Kabul timezone.
