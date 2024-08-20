package lsm_workload

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type LsmWorkload struct {
	gets     int
	puts     int
	deletes  int
	hitRatio float32

	keyPoolCount int
	keyPool      []string

	rs *RandomString
}

const (
	GET = iota
	PUT
	DELETE

	MaxBufferSize = 10 >> 20

	MaxKeyPoolCount = 10000000
)

func NewLsmWorkload(mode string, gets, puts, deletes int, hitRatio float32) *LsmWorkload {
	return &LsmWorkload{
		gets:     gets,
		puts:     puts,
		deletes:  deletes,
		hitRatio: hitRatio,

		keyPool: make([]string, 0, MaxKeyPoolCount),

		rs: NewRandomString(mode),
	}
}

func (lw *LsmWorkload) Generate() error {
	file, err := os.Create("workload.txt")
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	getCount, putCount, deleteCount := 0, 0, 0
	bufferSize := 0
	cnt := 0
	for getCount < lw.gets || putCount < lw.puts || deleteCount < lw.deletes {
		operation := rand.Intn(3)

		if (operation == GET && getCount >= lw.gets) || (operation == PUT && putCount >= lw.puts) || (operation == DELETE && deleteCount >= lw.deletes) {
			continue
		}

		cnt++
		if cnt%1000000 == 0 {
			fmt.Println("cnt: ", cnt)
		}

		var line string

		if operation == GET && getCount < lw.gets {
			key := lw.keyPool[rand.Intn(len(lw.keyPool))]
			line = fmt.Sprintf("g %s\n", key)

			getCount++
		}

		if operation == PUT && putCount < lw.puts {
			key := lw.rs.RandomKey()
			value := lw.rs.RandomValue()
			line = fmt.Sprintf("p %s %s\n", key, value)

			if len(lw.keyPool) >= MaxKeyPoolCount {
				lw.keyPool[rand.Intn(len(lw.keyPool))] = key
			} else {
				lw.keyPool = append(lw.keyPool, key)
			}

			putCount++
		}

		if operation == DELETE && deleteCount < lw.deletes {
			key := lw.rs.RandomKey()
			line = fmt.Sprintf("d %s\n", key)

			deleteCount++
		}

		if line != "" {
			_, err := writer.WriteString(line)
			if err != nil {
				return fmt.Errorf("failed to write to buffer: %w", err)
			}

			bufferSize += len(line)

			if bufferSize >= MaxBufferSize {
				if err := writer.Flush(); err != nil {
					return fmt.Errorf("failed to flush buffer: %w", err)
				}
				bufferSize = 0
			}
		}
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush buffer: %w", err)
	}

	return nil
}
