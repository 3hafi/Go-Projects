package main
import (
	"fmt"
	"os"
	"os/user"
	"sort"
	"strconv"
	"strings"
	"syscall"
	// "time"
	// "errors"
)
var (
	flagLong      bool // -l: long listing
	flagRecursive bool // -R: recursive listing
	flagAll       bool // -a: show hidden files
	flagReverse   bool // -r: reverse order
	flagSortTime  bool // -t: sort by modification time
)
func main() {
	dir := "."
	// Parse cl arguments, Flags can be combined (e.g. "-laRt")
	for _, arg := range os.Args[1:] {
		if arg == "-" {
			fmt.Println("my-ls: cannot access '-': No such file or directory")
			os.Exit(1)
		} else if strings.HasPrefix(arg, "-") {

			for _, ch := range arg[1:] {
				switch ch {
				case 'l':
					flagLong = true
				case 'R':
					flagRecursive = true
				case 'a':
					flagAll = true
				case 'r':
					flagReverse = true
				case 't':
					flagSortTime = true
				default:
					fmt.Printf("Unknown flag: -%c\n", ch)
					os.Exit(1)
				}
			}
		} else {
			// Assuming non-flag argument is a directory path
			dir = arg
		}
	}
	// Start listing from the given directory.
	if err := listDirectory(dir, ""); err != nil {
		fmt.Println("Error:", err)
	}
}
// Hint code 
// listDirectory lists files in a directory.  If recursive (-R) flag is set, it also lists subdirectories.// 'prefix' is used for formatting output in recursive mode.
func listDirectory(dir string, prefix string) error {
	fileInfos, err := getFileInfos(dir)
	if err != nil {
		return err
	}
	// if -a flag is not set, filter out hidden files
	if !flagAll {
		fileInfos = filterHidden(fileInfos)
	}
	// if -t flag is set sort files
	if flagSortTime {
		sort.Slice(fileInfos, func(i, j int) bool {
			return fileInfos[i].ModTime().After(fileInfos[j].ModTime())
		})
	} else {
		sort.Slice(fileInfos, func(i, j int) bool {
			return strings.ToLower(fileInfos[i].Name()) < strings.ToLower(fileInfos[j].Name())
		})
	}
	if flagReverse {
		reverseFileInfos(fileInfos)
	}

	if prefix != "" {
		fmt.Printf("\n%s:\n", dir)
	}
	// Print each file using either long format or simple name.
	for _, info := range fileInfos {
		if flagLong {
			fmt.Println(formatLong(info))
		} else {
			fmt.Println(info.Name())
		}
	}
	if flagRecursive {
		for _, info := range fileInfos {
			if info.IsDir() {
				subdir := dir + "/" + info.Name()
				listDirectory(subdir, prefix+"    ")
			}
		}
	}
	
	return nil
}
// getFileInfos returns file information for all items in the given directory.
func getFileInfos(dir string) ([]os.FileInfo, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	infos, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}
	return infos, nil
}
// filterHidden filters out hidden files if -a flag is not set.
func filterHidden(files []os.FileInfo) []os.FileInfo {
	var result []os.FileInfo
	for _, file := range files {
		if len(file.Name()) > 0 && file.Name()[0] == '.' {
			continue
		}
		result = append(result, file)
	}
	return result
}
func reverseFileInfos(files []os.FileInfo) { // reverse order of files
	for i, j := 0, len(files)-1; i < j; i, j = i+1, j-1 {
		files[i], files[j] = files[j], files[i]
	}
}

// formatLong returns a long listing string similar to "ls -l".
func formatLong(info os.FileInfo) string {
	// gets hard link count
	stat := info.Sys().(*syscall.Stat_t)
	// File mode string (permissions)
	mode := info.Mode().String()
	user, err := user.Current()
	if err != nil {
		fmt.Print(err)
	}
	username := user.Username
	// group, err := user.LookupGroupId(strconv.Itoa(int(stat.Gid)))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// groupname := group.Name
	size := strconv.FormatInt(info.Size(), 10) // 6
	modTime := info.ModTime().Format("Jan 2 15:04")
	fileName := info.Name()
	return fmt.Sprintf("%s %d %s %s %s %s", mode, stat.Nlink, username, size, modTime, fileName)
}

// still need group