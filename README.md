Golang builder pattern example for assigment 1

main.go - client
iBuilder, result - builder and result interfaces with structures of final products
getBuilder, postBuilder - concrete builders. getBuilder creates curl command string, postBuilder creates http request body 
director - director used for encapsulation, reuse of code, etc.
