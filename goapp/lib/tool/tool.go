package tool

import (
    "os"
    "bufio"
    "strings"
    "html/template"
	"strconv"
)

func Str2Int64(str string) int64{
	ret, err := strconv.ParseInt(str, 10, 0)
	if err != nil { panic(err.Error()) }
	return ret
}

func TemplateWithFile(key string, path string) *template.Template{
    t := template.New(key)
    var err error
    t, err = t.Parse(StringWithFile(path))
    if err != nil {
        panic( err.Error() )
    }
    return t
}

func StringWithFile(path string) string {
    lines, err := ReadLines( path )
    if err != nil {
        panic( err.Error() )
    }
    return strings.Join( lines, "" )
}

func ReadLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

type ByTypeInt64 []int64
func (s ByTypeInt64) Len() int { return len(s) }
func (s ByTypeInt64) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByTypeInt64) Less(i, j int) bool { return s[i] < s[j] }