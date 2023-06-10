# Indexer
This is a sample App for the [Mister](https://github.com/davidfregoli/mister) MapReduce framework

### Sample input:
- file `input-one`:
  ```
  this is the first file
  ```
- file `input-two`:
  ```
  this is certainly the second file
  ```
### Sample output:
```
this 2 input-one,input-two
is 2 input-one,input-two
the 2 input-one,input-two
first 1 input-one
file 2 input-one,input-two
certainly 1 input-two
second 1 input-two
```
