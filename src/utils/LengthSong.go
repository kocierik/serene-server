package utils

import (
	"fmt"
	"io"
	"mime/multipart"

	"github.com/tcolgate/mp3"
)

func GetLengthSong(mp3File multipart.File) int {
	d := mp3.NewDecoder(mp3File)
	var f mp3.Frame
	skipped := 0
	t := 0.0
	for {
		if err := d.Decode(&f, &skipped); err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return 0
		}
		t = t + f.Duration().Seconds()
	}
	return int(t)
}
