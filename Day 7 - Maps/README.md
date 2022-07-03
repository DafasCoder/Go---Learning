# Day 7 - Maps

What did we learn today!!!!!!!

1. Creating maps are definited like this "value map[type] type"
2. Maps are like reference type so there no need to add a &.A map value is a pointer to a runtime.hmap structure.
3. With maps they can have a value of nil. So attempt to read from a empty map will cause runtime panic.
   - So intialize the map either like this
     var dictionary = map[string] string{}
     or
     var dictionary = make(map[string]string)
4. To delete from a map you can use this global function delete(map,key)
