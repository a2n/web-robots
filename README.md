# Web Robots.txt Checker
* A web robots.txt checker. 


# Usage
```
import (
	"github.com/a2n/web-robots"
	"fmt"
)

func main() {
	r := robots.NewRobots()
	fmt.Println(r.IsAllowURLString("Safari", "https://www.google.com/search/a"))
	fmt.Println(r.IsAllowURLString("Chrome", "https://www.yahoo.com/p/a"))
}

```

#License
Web-robots is released under the CC0 1.0 Universal.
