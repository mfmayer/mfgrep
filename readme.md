# **mf**grep

Similar to grep, but for searching **m**ultiple lines that contain **f**ixed strings

## Install

**mf**grep is a written in go and can therefore be installed via: `go get github.com/mfmayer/mfgrep` (go must be installed).

## Usage of **mf**grep

```
$ mfgrep
Usage: mfgrep [SUBSTRING]
  mfgrep is intended to work with pipes and searches SUBSTRINGs in consecutive lines.
  The last SUBSTRING is searched in all consecutive lines that are printed when it is found..
Example: 'ls -al | mfgrep . ..'
  searches for lines containing "." and multiple lines containing "." in first and ".." in all consecutive lines.
```

## Example

Reason for writing this tool was filtering kubernetes logs for specific log entries and to include all following lines beginning with prefixed spaces.

Example: Only log entries with `01052cbf-c6de-4a97-8f40-4c2f674ab855` and corrsesponding consecutive lines with additional attributes shall be printed:
```
TRACE: 2020/02/05 16:06:48 source.go:289: Something happend [17cfaf94-cc0f-4ec3-9735-02d4eb0b1721]
TRACE: 2020/02/05 16:06:48 dump.go:227: Something else happend [01052cbf-c6de-4a97-8f40-4c2f674ab855]
     any_pointer               : 0xc000340060
     any_attribute             : "attribute_value_A"
     len(any_array)            : 2
TRACE: 2020/02/05 16:06:49 dump.go:227: Something else happend [17cfaf94-cc0f-4ec3-9735-02d4eb0b1721]
     any_pointer               : 0xc000120300
     any_attribute             : "attribute_value_B"
     len(any_array)            : 5
```

```
$ cat test/testtrace.log | mfgrep 01052cbf-c6de-4a97-8f40-4c2f674ab855 "     "
TRACE: 2020/02/05 16:06:48 dump.go:227: Something else happend [01052cbf-c6de-4a97-8f40-4c2f674ab855]
     any_pointer               : 0xc000340060
     any_attribute             : "attribute_value_A"
     len(any_array)            : 2
```