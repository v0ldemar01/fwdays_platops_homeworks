#Person: {
    name:     string
    age:      int & >=0 & <=120
    email:    string & =~"^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
    hobbies?: [...string]
}

// Example instance
person: #Person & {
    name:    "John Doe"
    age:     30
    email:   "john@example.com"
    hobbies: ["reading", "coding"]
}
