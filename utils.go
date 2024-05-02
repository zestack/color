package color

import (
	"regexp"
	"strconv"
	"sync"
)

const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var (
	resetBytes = Bytes(Reset)
	ansiRegex  = regexp.MustCompile(ansi)
)

// StripAnsi 清除 ANSI 转义代码
// https://en.wikipedia.org/wiki/ANSI_escape_code
// https://github.com/acarl005/stripansi/blob/master/stripansi.go
func StripAnsi(str []byte) []byte {
	return ansiRegex.ReplaceAll(str, []byte{})
}

// Bytes converts a list of ColorAttributes to a byte array
func Bytes(attrs ...Attribute) []byte {
	bts := make([]byte, 0, 20)
	bts = append(bts, escape[0], '[')
	if len(attrs) > 0 {
		bts = append(bts, strconv.Itoa(int(attrs[0]))...)
		for _, a := range attrs[1:] {
			bts = append(bts, ';')
			bts = append(bts, strconv.Itoa(int(a))...)
		}
	} else {
		bts = append(bts, strconv.Itoa(int(Bold))...)
	}
	bts = append(bts, 'm')
	return bts
}

var (
	colors = []Attribute{
		20, 21, 26, 27, 32, 33, 38, 39, 40, 41,
		42, 43, 44, 45, 56, 57, 62, 63, 68, 69,
		74, 75, 76, 77, 78, 79, 80, 81, 92, 93,
		98, 99, 112, 113, 128, 129, 134, 135,
		148, 149, 160, 161, 162, 163, 164, 165,
		166, 167, 168, 169, 170, 171, 172, 173,
		178, 179, 184, 185, 196, 197, 198, 199,
		200, 201, 202, 203, 204, 205, 206, 207,
		208, 209, 214, 215, 220, 221,
	}
	mutex      sync.Mutex
	namespaces map[string]*Value
)

func Namespace(name string) *Value {
	mutex.Lock()
	defer mutex.Unlock()
	if namespaces == nil {
		namespaces = make(map[string]*Value)
	}
	ns, ok := namespaces[name]
	if !ok {
		ns = NewValue(name, 38, 5, selectColor(name))
		namespaces[name] = ns
	}
	return ns
}

func selectColor(namespace string) Attribute {
	var hash int32
	ns := []rune(namespace)
	for _, charCode := range ns {
		hash = ((hash << 5) - hash) + charCode
		hash |= 0 // Convert to 32bit integer
	}
	if hash < 0 {
		hash *= -1
	}
	// 76 == len(colors)
	return colors[int32(int(hash)%76)]
}
