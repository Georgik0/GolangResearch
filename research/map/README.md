### Initialisation
```go
m := make(map[key_type]value_type)
m := new(map[key_type]value_type)
m := map[key_type]value_type{key1: val1, key2: val2}

var m map[key_type]key_type
var m = map[key_type]key_type{key1: val1, key2: val2}
```

### Basic operation
```go
// Insert
m[key] = value

// Delete
delete(m, key)

// Search
value, ok = m[key] 
```
---
### For operation
```go
for k, v := range m {
	print(k,v)
}
```
- Hash table is unordered
---
```go
mInt := make(map[int]int)
mStr := make(map[int]string)
mBool := make(map[int]bool)

num := mInt[0] // v = 0
s := mStr[0] // s = ""
b := mBool[0] // b = false
```
- An attempt to fetch a map value with a key that is not present in the map will return the zero value for the type of the entries in the map.
---
### Assignment to entry in nil map
```go
var m map[int]int
m[1] = 1 // panic !!!
```
```go
var mNil map[int]int
var m = map[int]int{}

println(mNil == nil) // true
println(m == nil) // false
```
---
### Map as function argument
```go
func main() {
    m := make(map[int]int)
    f(m)
	v := m[1] // v = 1
}

func f(m map[int]int) {
	m[1] = 1
}
```
---
### How to search in a simplified form
```go
// value := m[key] --> func mapaccess1(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer 

func mapaccess1(t *mapType, m *mapHeader, key unsafe.Pointer) unsafe.Pointer {
    hash := hashfn(key)
    bucket := hash % nbuckets // Bucket which you should search in
    extra := byte(hash >> 56) // Top 8 bits of hash - extra hash 
    b := &m.buckets[bucket] // Determine the position of the “bucket” in memory (address arithmetic)
}
```