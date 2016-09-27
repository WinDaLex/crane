package history

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/windalex/crane/path"
)

func Add(log string) {
	f, _ := os.OpenFile(path.History(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	text := fmt.Sprintf("[%s] %s\n", time.Now().Format(time.UnixDate), log)
	f.WriteString(text)
}

func Show() string {
	data, _ := ioutil.ReadFile(path.History())
	return string(data)
}

func Clear() {
	os.Remove(path.History())
}
