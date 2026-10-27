package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.0" )

func CBvky3brWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g5wUn8kaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6i5gEzG8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q4ICd8ziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qpIRyHAVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2v5GQcTwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P2Si3dNTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KHDh2SLeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l7dA8N5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T9zYtopqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J19CVwXIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ElzCpPp1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func biqI9N8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nGyJEd8cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I5VXe4M8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tZFwRDEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YsLQDNAmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func frlD80DNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2KMO8AsHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FYVrigv8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d23TPbNOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eHFSFwIJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m6gHYTDRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qlPHA2dyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7HxurV4pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3VkKWJhCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MQLIy4DRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ksdXoUOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D4GoSrvTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a7aHr0b6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C4GuL2kMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tuPEKVONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bTakY3TTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cYqADOFvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U0xepb91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6n5ULGUpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZVTZ7OirWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func amIEfwExWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7abLyt5WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9zakQAuTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D3r4Lo9iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func adgbDnj4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mQzZCeZuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func shcIeHSeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DhDqIanXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cjhnuZUxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dv1GJH4GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4s2nziylWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m7ghkm3dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TX3vCgz8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HCdobDSOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xLf6U7FeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 18fBRXjSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g4fOpJB1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ucG7mV6TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qhtBBvRSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C8Jjh9NDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DJglOgMWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CVi1HHVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6GQunkyZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jkK5oUkqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DElnOoG6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Mm7iVLwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wFNP7L4PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Vih4jHRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ctJxJNcsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QSud0KTIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DeOORRu0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sey4rh3vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YkeEyX3NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4AUtUD8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mvC1vEwgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zOtDt7LOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l3I3frtxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tjmmDJkRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func os0EqHRtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qpwHq1LnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FWfkZaEwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HTyyyRkRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uj7klfC8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eniavGC0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wi4Kxc3zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g3UWIZwmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ciTD5Ar7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kSmOPIKOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aY5j24mgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pw5F6ZhvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6foYTCqnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iZNrFBldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i1lkOwwOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bQefbNqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0VoymoVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MiqlQbNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OREflLpEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkd0z5UtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

