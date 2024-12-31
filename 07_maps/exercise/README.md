1. How to check if a key/element pair is in a map?

Can check the boolean which is returned when we retrieve the value from a given key.
If this value is true then the pair exists. 
In our code the key 'Ed' exists with a value of 50, whereas we do not have a key value pair with a key of Fred.
m, ok := myMap["Ed"]
fmt.Println("Key Ed:", m, ok)
Key Ed: 50 true

m, ok := myMap["Fred"]
fmt.Println("Key Fred:", m, ok)
Key Fred: 0 false


2. How are Go Maps implemented internally?

Maps are key-value pairs which are stored in a hash table.
The key is 'hashed' to provide a unique id, so it is this unique key which is stored along with the value from the original key-value pair. 
In go the hashKey-values pairs are stored in a 'bucket', which can hold a maximum of 8 hashKey-value pairs.  If the number of values being stored is larger than 8 then another bucket will be generated and a pointer to the new bucket stored on the original bucket.  Thus all buckets for these key-value pairs will be linked.  Also the hash key generated for a key includes the id of the bucket so this is one of the reasons why retrieving data is so fast.  The buckets can only contain a maximum of 8 values so there will never be more than 8 to search through to retreive the value being searched for.



3. How to iterate over a map?

Use the range function as we have done previously.
for key, value := range m {}
So we loop through all values in the map retreiving each key, value pair.



4. If a map M does not contain the key K, what will return M[K]?

go will return the null value of whatever type the value element is, so if an integer it would return 0, if a string it would return Null.
So it would be better to see whether the key exists as above:
e.g. m, ok := M[K]
fmt.Println("Key K:", m, ok)
Key K: 0 false