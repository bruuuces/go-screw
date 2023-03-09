package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const PathSeparator = string(os.PathSeparator)

// RemoveContents 删除目录下的所有文件和子目录(不包括该目录)
func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer func() {
		_ = d.Close()
	}()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

// RenameContents 读取目录下的所有文件和子目录(不包括该目录), 并移动到新目录
func RenameContents(oldpath, newpath string) error {
	d, err := os.Open(oldpath)
	if err != nil {
		return err
	}
	defer func() {
		_ = d.Close()
	}()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.Rename(filepath.Join(oldpath, name), filepath.Join(newpath, name))
		if err != nil {
			return err
		}
	}
	return nil
}

// ReadContentsByExt 读取指定目录下指定后缀的文件
func ReadContentsByExt(path, ext string) ([]string, error) {
	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}
	readDir, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var files []string
	for _, file := range readDir {
		if file.IsDir() {
			continue
		}
		if ext != "" && !strings.HasSuffix(strings.ToUpper(file.Name()), strings.ToUpper(ext)) {
			continue
		}
		files = append(files, path+PathSeparator+file.Name())
	}
	return files, nil
}

// IsExist 判断所给路径文件/文件夹是否存在
// path: 文件/文件夹路径
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func MoveFile(oldpath, newpath string) error {
	err := os.Rename(oldpath, newpath)
	if err == nil {
		return nil
	}
	oldfile, err := os.Open(oldpath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}
	newfile, err := os.Create(newpath)
	if err != nil {
		oldfile.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer newfile.Close()
	_, err = io.Copy(newfile, oldfile)
	oldfile.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(oldpath)
	if err != nil {
		return fmt.Errorf("Failed removing original file: %s", err)
	}
	return nil
}
