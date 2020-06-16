package collector

import (
  "os/exec"
  "strings"
)

const ShellToUse = "bash"
func RunShell(scriptFile string) string {
  var line string
  cmdName := scriptFile
  cmd := exec.Command(ShellToUse, "-c", cmdName)
  stdout, _ := cmd.CombinedOutput()
  line = strings.Trim(string(stdout), "\n")
  return line
}
