# tech test

Examples about how to solve the following challenges using golang and netcore (8).

## palindrome

validate if a given word is a palindrome or not

## two items 

validate whether there are two elements in an array whose sum result is a given specific value

## string times

retrieve how many times a specific character is in a string
 
## anagrama

validate if a string is a permutation of another

## compactor/decompactor

compact and optionally decompact a string

`string = "aaaaabbbadddcc"`

`compactador(string) = "a5b3a1d3c2"`

## maximum difference

the difference between the elements of a matrix is defined by a[j] - a[i] where i < j and a[i] < a[j]
complete la funcion maximumDifference which has and integer array as an input and returns the maximum difference
if the maximum difference cannot be calculated (for example, the elements are in descedent order) the function should return -1

example: `[15, 3, 6, 10]`

differences: 

```
6 - 3 = 3
10 - 3 = 7
10 - 6 = 4
```
maximum difference = `7`

## merge ordenado

In merge ordenado arrays are created and merged to create a new sorted array containing the values from both arrays

example: `a=[1, 2, 3]` `b=[2, 5, 5]` then `c=[1, 2, 2, 3, 5, 5]`

## merge arrays

receives two sorted arrays and produce a new array with all the elements from the two arrays sorted in ascending order
