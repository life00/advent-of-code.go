# [Day 1](https://adventofcode.com/2023/day/1)

## Part 1

### Summary

I have the following input:

```
erjfk2khk38f
4jk3jk29
okrjo3lk
```

I should trim each line to only contain the first and last number. If there is only one number in the line then it will be both the first and the last. Afterwards, sum all the numbers.

In the above example it would be 29, 48, 33 which equals to 110.

### Solution

- read the text file and save contents
- split the variable into an array based on newline character
  - remove the faulty last array element
- for further operations loop through every array entry (every line)
- remove all non digit characters in each string
- get the first and last character of a string
- convert string to integer
- add the value to the sum
- print the results

## Part 2

### Summary

The input also has digits written as words. It can be: one, two, three, four, five, six, seven, eight, and nine. It is necessary to account for them too.

Also notice that the following string: `jofrkjneightwolrk3k` has both `eight` and `two` at the beginning, but they share a `t`. The word closer to the (left or right) end of string is always prioritized.

### Solution

- firstly fix overlapping word numbers like (tw(o)ne) by just replacing them with twoone, to prevent further issues
- use initial raw data and replace words to numbers with RegEx in a for loop
- use the rest of the code as it was before
