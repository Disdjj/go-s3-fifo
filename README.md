# Go-S3-FIFO

> a simple Go implementation of the [S3FIFO](https://s3fifo.com/blog/2023/08/01/fifo-queues-are-all-you-need-for-cache-eviction/#implication-of-a-large-one-hit-wonder-ratio) 

This is a Go project that implements a variant of the S3FIFO (Segmented LRU with Frequency in Count) cache. The cache is divided into three segments: Small (S), Medium (M), and Ghost (G). The cache uses the frequency count of the keys to decide which segment a key should reside in.

## Getting Started

To get a local copy up and running, follow these simple steps.

### Prerequisites

- Go 1.22.2 or later


## Usage

The cache is implemented as a Go generic type, so it can be used with any comparable type as the key and any type as the value.

Here is a basic example of how to use the cache:

```go
package main

import (
	"fmt"
	"github.com/Disdjj/go-s3-fifo"
)

func main() {
	cache := go_s3_fifo.NewS3FIFOCache[int, string](10)

	cache.Set(1, "one")
	value, ok := cache.Get(1)

	if ok {
		fmt.Println(value) // prints: one
	}
}
```

## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are greatly appreciated.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Project Link: [https://github.com/Disdjj/go-s3-fifo](https://github.com/Disdjj/go-s3-fifo)

Please note that the "Contact" section needs to be filled out with your personal contact information.