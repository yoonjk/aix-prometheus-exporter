package collector

import (
  "bufio"
  "io"
  "os"
  "time"
  "regexp"
  "strings"
)

const (
  COLUMN_SEPARATOR = ";"
)

var (
  logStats map[string]logFileStats
)

/*
  load log list file
*/
func load() {
  var c chan logFileStats = make(chan logFileStats)

  loadListFile("config/logfiles.lst")
  go updateLogInfo(c)

  for _, value := range logStats {
    go tail(value, c);
  }
}

/*
  update logging event to hash map memorys
*/
func updateLogInfo(c chan logFileStats) {
  for {
    msg := <- c
    logStats[msg.LabelName] = msg
    time.Sleep(time.Second * 1)
  }
}

/*
  load log list file
*/
func loadListFile(filename string) {
  file, err := os.Open(filename)
  logStats = make(map[string]logFileStats)

  if err != nil {
    panic(err)
  }

  defer file.Close()

  reader := bufio.NewReader(file)
    
  for {
    line, _, err := reader.ReadLine()
   
    if err == io.EOF {
      break;
    }

    columns := strings.Split(string(line), COLUMN_SEPARATOR)

    logStats[columns[0]] = logFileStats{columns[0], columns[1], columns[2], columns[3], 0, false, getTime()}
  }   
}

/*
 tail specific file
*/
func tail(fileInfo logFileStats, c chan logFileStats) {
  var (
    filename string
    matched bool = false
  )

  pattern := regexp.MustCompile(fileInfo.Pattern)
  filename = fileInfo.FileName
  f, err := os.Open(filename)

  if err != nil {
    panic(err)
  }

  defer f.Close()

  reader := bufio.NewReader(f)
  info, err := f.Stat()

  if err != nil {
      panic(err)
  }

  oldSize := info.Size()
  f.Seek(0, os.SEEK_END)

  for {
    for line, _, err := reader.ReadLine(); err != io.EOF; line, _, err = reader.ReadLine() {
      matched = pattern.MatchString(string(line))

      if matched == true{
        fileInfo.Count++
        fileInfo.Status = true
        fileInfo.LogTime = getTime() 

        c <- fileInfo
      }
    }

    pos, err := f.Seek(0, io.SeekCurrent)

    if err != nil {
      panic(err)
    }

    for {
      time.Sleep(time.Second)
      newinfo, err := f.Stat()
      if err != nil {
        panic(err)
      }

      newSize := newinfo.Size()

      if newSize != oldSize {
        if newSize < oldSize {
          f.Seek(0, 0)
        } else {
          f.Seek(pos, io.SeekStart)
        }

        reader = bufio.NewReader(f)
        oldSize = newSize

        break
      }
    }
  }
}

/*
 get statics of logging events
*/
func (c *logEventCollector) getLogInfo() (map[string]logFileStats, error) {
  return logStats, nil
}


