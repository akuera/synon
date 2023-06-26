# Synon

Synon is a lightweight Go library that simplifies updating values within data structures. It provides a single function that allows you to replace values in the left parameter with values from the right parameter, under specific conditions.

## Installation
```shell
go get github.com/your-username/synon
```

## Usage

Synon exposes a single function, `Merge`, that takes two parameters: the left parameter (destination) and the right parameter (source). The function updates the values in the left parameter according to the following conditions:

1. If a value in the left parameter is empty or nil, and the corresponding value in the right parameter is non-empty or non-nil, the value in the left parameter will be replaced with the value from the right parameter.

2. If the value in the left parameter is different from the corresponding value in the right parameter, the value in the left parameter will be updated with the value from the right parameter.

Here are examples of how to use Synon:

```go
package main

import (
	"fmt"
	"github.com/your-username/synon"
)

type Person struct {
	Name     string
	Age      int
	Email    string
	Location string
}

func main() {
	// Define the destination struct
	destination := Person{
		Name:     "John Doe",
		Age:      30,
		Email:    "",
		Location: "New York",
	}

	// Define the source struct
	source := Person{
		Name:     "Jane Smith",
		Age:      30,
		Email:    "jane@example.com",
		Location: "San Francisco",
	}

	// Merge values from the source into the destination
	synon.Merge(&destination, source)

	// Print the updated destination struct
	fmt.Println(destination)
}

```

In the updated example, we define a `Person` struct with fields such as `Name`, `Age`, `Email`, and `Location`. We then create instances of the `Person` struct for the `destination` and `source` objects. The `Merge` function is used to merge the values from the `source` struct into the `destination` struct based on the specified conditions.

## Use Cases

Synon can be particularly useful in scenarios where you need to update a document or data structure based on new information, such as when working with databases or APIs. One common use case is with MongoDB, where you may fetch a document, modify certain fields, and then update the document with the changes. Synon simplifies this process by automatically handling the updates based on the specified conditions.

By using Synon, you can avoid manually checking each field for changes and writing conditional logic to perform the updates. The library abstracts away the complexities, allowing you to focus on the core functionality of your application.