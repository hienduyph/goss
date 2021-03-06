// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package chunks

import "math"

func ChunkString(inputs []string, batch int) <-chan []string {
	out := make(chan []string, len(inputs)/batch+1)
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

func ChunkStreamString(inputs <-chan string, batch int) <-chan []string {
	out := make(chan []string, batch+1)
	go func() {
		chunk := make([]string, 0, batch)
		for item := range inputs {
			chunk = append(chunk, item)
			if len(chunk) == batch {
				out <- chunk
				chunk = make([]string, 0, batch)
			}
		}
		if len(chunk) > 0 {
			out <- chunk
		}
		close(out)
	}()
	return out
}
