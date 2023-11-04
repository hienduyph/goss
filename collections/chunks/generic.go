package chunks

import (
	"math"
)

func Chunk[T any](inputs []T, batch int) <-chan []T {
	out := make(chan []T, len(inputs)/batch+1)
	go func() {
		length := len(inputs)
		chunks := int(math.Ceil(float64(length) / float64(batch)))
		for i, end := 0, 0; chunks > 0; chunks-- {
			end = (i + 1) * batch
			if end > length {
				end = length
			}
			chunk := inputs[i*batch : end]
			out <- chunk
			i++
		}
		close(out)
	}()
	return out
}

func ChunkStream[T any](inputs <-chan T, batch int) <-chan []T {
	out := make(chan []T, batch+1)
	go func() {
		chunk := make([]T, 0, batch)
		for item := range inputs {
			chunk = append(chunk, item)
			if len(chunk) == batch {
				out <- chunk
				chunk = make([]T, 0, batch)
			}
		}
		if len(chunk) > 0 {
			out <- chunk
		}
		close(out)
	}()
	return out
}
