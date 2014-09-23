package hello

import (
    "os"
    "bufio"
    "strings"
    "html/template"
)

func templateWithFile(key string, path string) *template.Template{
    t := template.New(key)
    var err error
    t, err = t.Parse(stringWithFile(path))
    if err != nil {
        panic( err.Error() )
    }
    return t
}

func stringWithFile(path string) string {
    lines, err := readLines( path )
    if err != nil {
        panic( err.Error() )
    }
    return strings.Join( lines, "" )
}

func readLines(path string) ([]string, error) {
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

type int64Array []int64
func (s int64Array) Len() int { return len(s) }
func (s int64Array) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s int64Array) Less(i, j int) bool { return s[i] < s[j] }