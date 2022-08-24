/* Alta3 Research | RZFeeser
   Refer - using the defer statement  */

package main

import(
    "os"
    "io"
    "os/exec"
)

func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    defer src.Close()

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }
    defer dst.Close()

    return io.Copy(dst, src)
}

func main() {

    cmd := exec.Command("clear") // Clear the term screen before program run
    cmd.Stdout = os.Stdout
    cmd.Run()

    CopyFile("/tmp/example-copy2.txt", "/tmp/example-copy.txt")
}

