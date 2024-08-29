# EncryptDB

**EncryptDB** is a decentralized database engine designed in Go, providing secure and scalable data storage solutions. As part of the Encrypt Distributed Ledger (EDL) organization, EncryptDB is built to support distributed systems and blockchain projects by ensuring data integrity, high availability, and enhanced security.

## Table of Contents

1. [Overview](#overview)
2. [Features](#features)
3. [Installation](#installation)
4. [Quick Start](#quick-start)
5. [Usage](#usage)
6. [Configuration](#configuration)
7. [Contributing](#contributing)
8. [License](#license)

## Overview

EncryptDB is engineered to offer a robust decentralized database solution that integrates seamlessly with blockchain and distributed applications. It is written in Go to leverage the language's concurrency model and strong standard library, aiming to provide a high-performance, reliable, and secure data storage engine.

## Features

- **Decentralized Architecture**: Supports distributed data storage across multiple nodes, ensuring high availability and fault tolerance.
- **Secure Storage**: Uses advanced cryptographic techniques to protect data integrity and confidentiality.
- **Scalable**: Designed to handle large datasets and high throughput, optimizing both read and write operations.
- **Transaction Support**: Provides ACID compliance for transactional operations, ensuring data consistency.
- **Lightweight and Fast**: Minimal dependencies and optimized for performance in resource-constrained environments.
- **Easy Integration**: Seamlessly integrates with existing Go applications and distributed systems.

## Installation

To install EncryptDB, you'll need to have Go installed on your machine. You can install the library using `go get`:

```bash
go get github.com/EncrypteDL/EncryptDB
```

## Quick Start

Here's a simple example to get you started with EncryptDB:

```go
package main

import (
    "fmt"
    "log"
    "github.com/EncrypteDL/EncryptDB"
)

func main() {
    // Initialize the database
    db, err := EncryptDB.Open("path/to/db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Insert a record
    err = db.Put("key", "value")
    if err != nil {
        log.Fatal(err)
    }

    // Retrieve a record
    value, err := db.Get("key")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Value:", value)
}
```

## Usage

### Inserting JSON Data
EncryptDB allows you to store and manage JSON data efficiently. Here’s how you can insert a JSON object into the database:

### Example JSON Object

```json
{
  "Date": "2024-08-23T20:49:39.779374+01:00",
  "FileFormat": "tar.gz",
  "FilePath": "/var/folders/_0/h7kj277n3bvbdh1x9ymfdh1w0000gn/T/TestNewTarGzBackupoutpath_directory973048178/002/001.tar.gz",
  "Stores": ["yeet"],
  "Checksum": {
    "Type": "sha256",
    "Value": "61b64d243a90ba8a946268f0f81481666641e1afddd5820f6e9ec1be1ffe445f"
  },
  "Size": 0
}
```

### Storing JSON in EncryptDB

To store a JSON object, you can serialize it to a string and then use the `Put` method:

```go
package main

import (
    "encoding/json"
    "log"
    "github.com/EncrypteDL/EncryptDB"
)

func main() {
    db, err := EncryptDB.Open("path/to/db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // JSON object to be stored
    jsonData := map[string]interface{}{
        "Date": "2024-08-23T20:49:39.779374+01:00",
        "FileFormat": "tar.gz",
        "FilePath": "/var/folders/_0/h7kj277n3bvbdh1x9ymfdh1w0000gn/T/TestNewTarGzBackupoutpath_directory973048178/002/001.tar.gz",
        "Stores": []string{"yeet"},
        "Checksum": map[string]string{
            "Type":  "sha256",
            "Value": "61b64d243a90ba8a946268f0f81481666641e1afddd5820f6e9ec1be1ffe445f",
        },
        "Size": 0,
    }

    // Serialize JSON data to a string
    jsonString, err := json.Marshal(jsonData)
    if err != nil {
        log.Fatal(err)
    }

    // Store JSON string in EncryptDB
    err = db.Put("backup:001", string(jsonString))
    if err != nil {
        log.Fatal(err)
    }

    log.Println("JSON data has been stored successfully.")
}
```

### Opening a Database

To open a database, use the `Open` method, specifying the path to your database file:

```go
db, err := EncryptDB.Open("path/to/db")
if err != nil {
    // handle error
}
defer db.Close()
```

### Inserting and Retrieving Data

Use `Put` to insert data and `Get` to retrieve it:

```go
err := db.Put("username", "zakaria")
if err != nil {
    // handle error
}

value, err := db.Get("username")
if err != nil {
    // handle error
}
fmt.Println("Username:", value)
```

## Configuration

EncryptDB can be configured via a configuration file or environment variables. Configuration options include:

- `StoragePath`: Path where the database files will be stored.
- `EncryptionKey`: Key used for encrypting data at rest.
- `ReplicationFactor`: Number of copies of the data to be replicated across nodes.

## Contributing

We welcome contributions to EncryptDB! To contribute:

1. Fork the repository.
2. Create a new branch with a descriptive name.
3. Make your changes.
4. Submit a pull request detailing your changes.

Please ensure that your code adheres to the project's coding standards and includes appropriate tests.

## License

EncryptDB is released under the MIT License. See the [LICENSE](LICENSE) file for more details.
