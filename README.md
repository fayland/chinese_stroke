chinese_stroke
==============

Go: Retrieve the stroke count of Chinese character.


```
package main

import (
    cs "github.com/fayland/chinese_stroke"
)

func main() {
    // 6
    stroke := cs.GetStroke("好")

    // 2
    stroke = cs.GetStroke("人")
}

```