package main

import (
  "fmt"
  "flag"
  "path/filepath"
  "archive/tar"
  "os"
  "crypto/md5"
  "crypto/sha1"
  "crypto/sha256"
  "crypto/sha512"
  "io"
  "encoding/json"
  "encoding/hex"
  "time"
  "hash"
)

type Options struct {
  SkipMd5    *bool `json:"skipMd5"`
  SkipSha1   *bool `json:"skipSha1"`
  SkipSha256 *bool `json:"skipSha256"`
  SkipSha512 *bool `json:"skipSha512"`
}

type Checksum struct {
  Path   string `json:"path"`
  Md5    string `json:"md5,omitempty"`
  Sha1   string `json:"sha1,omitempty"`
  Sha256 string `json:"sha256,omitempty"`
  Sha512 string `json:"sha512,omitempty"`
}

type ResultFile struct {
  Checksums []Checksum `json:"checksums"`
  Size      int `json:"size"`
}

const (
  printCount = 100
)


var checksums []Checksum

func tarChecksums(path string, options Options) error {
  f, err := os.Open(path)
  if err != nil {
    return err
  }
  defer f.Close()

  tarReader := tar.NewReader(f)

  for {
    header, err := tarReader.Next()

    if err == io.EOF {
      return nil
    }

    if err != nil {
      return err
    }
    // fmt.Println(name)
    if header.Typeflag == tar.TypeReg {
      // size := header.Size
      name := header.Name
      err = addToChecksum(name, tarReader, options)
      if err != nil {
        return err
      }
    }
  }
  return nil
}

func dirChecksum(path string, options Options) error {
  return filepath.Walk("test", func(path string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }

    if !info.IsDir() {
      f, err := os.Open(path)
      if err != nil {
        return err
      }
      defer f.Close()

      err = addToChecksum(path, f, options)
      if err != nil {
        return err
      }
    }
    return nil
  })
}

func pop(a []hash.Hash) (hash.Hash, []hash.Hash) {
  len := len(a) - 1
  return a[len].(hash.Hash),a[:len]
}

func addToChecksum(path string, readIo io.Reader, options Options) error {
  var hashes []hash.Hash

  if !*options.SkipMd5 {
    hashes = append(hashes, md5.New())
  }
  if !*options.SkipSha1 {
    hashes = append(hashes, sha1.New())
  }
  if !*options.SkipSha256 {
    hashes = append(hashes, sha256.New())
  }
  if !*options.SkipSha512 {
    hashes = append(hashes, sha512.New())
  }


  writers := make([]io.Writer, len(hashes))
  for i, h := range hashes {
    writers[i] = io.Writer(h)
  }

  multiWriter := io.MultiWriter(writers...)

  if _, err := io.Copy(multiWriter, readIo); err != nil {
    return err
  }

  checksum := Checksum{Path: path}

  if !*options.SkipMd5 {
    var md5Hash hash.Hash
    md5Hash, hashes = pop(hashes)
    checksum.Md5 = hex.EncodeToString(md5Hash.Sum(nil))
  }
  if !*options.SkipSha1 {
    var sha1Hash hash.Hash
    sha1Hash, hashes = pop(hashes)
    checksum.Sha1 = hex.EncodeToString(sha1Hash.Sum(nil))
  }
  if !*options.SkipSha256 {
    var sha256Hash hash.Hash
    sha256Hash, hashes = pop(hashes)
    checksum.Sha256 = hex.EncodeToString(sha256Hash.Sum(nil))
  }
  if !*options.SkipSha512 {
    var sha512Hash hash.Hash
    sha512Hash, hashes = pop(hashes)
    checksum.Sha512 = hex.EncodeToString(sha512Hash.Sum(nil))
  }

  checksums = append(checksums, checksum)
  len := len(checksums)
  if (len % printCount == 0) {
    fmt.Printf("  Checking: %v\n", len)
  }
  return nil
}

func main() {
  var pathP = flag.String("name", "", "path to check(sum)")
  var outP = flag.String("out", "out.json", "file output")

  var options = Options{
    flag.Bool("skip-md5", false, "skip md5"),
    flag.Bool("skip-sha1", false, "skip sha1"),
    flag.Bool("skip-sha256", false, "skip sha256"),
    flag.Bool("skip-sha512", false, "skip sha512"),
  }

  flag.Parse()
  path := *pathP
  if len(path) == 0 {
    panic("needs name")
  }

  fmt.Printf("Check(sum)ing %v... ", path)
  fileInfo, err := os.Stat(path)
  if err != nil {
  	panic(err)
  }

  start := time.Now()
  if fileInfo.IsDir() {
    err = dirChecksum(path, options)
  } else {
    err = tarChecksums(path, options)
  }
  duration := time.Since(start)
  fmt.Println(duration)
  if err != nil {
    panic(err)
  }

  fmt.Println("Saving JSON")
  resultFile := ResultFile{ Checksums: checksums, Size: len(checksums) }
  b, err := json.MarshalIndent(resultFile, "", "  ")
  if err != nil {
    panic(err)
  }

  err = os.WriteFile(*outP, b, 0644)
  if err != nil {
    panic(err)
  }
}
