package backup

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mholt/archiver/v4"
)

type Backup struct {
	ID          uuid.UUID
	Destination string
	Timestamp   int64
}

func New() (*Backup, error) {
	currentTime := time.Now()
	filename := currentTime.Format("2006-01-02_15-04-05")

	if err := createArchive("./testing", "./testing/backups/"+filename); err != nil {
		return nil, fmt.Errorf("could not create server backup. %s", err)
	}

	backup := &Backup{
		ID:          uuid.New(),
		Destination: "./testing/backups/" + filename + ".tar.gz",
		Timestamp:   currentTime.Unix(),
	}

	backups, err := ReadOverview()
	if err != nil {
		return nil, fmt.Errorf("could not register new backup. %s", err)
	}

	backups[backup.ID.String()] = backup

	err = WriteOverview(backups)
	if err != nil {
		return nil, fmt.Errorf("could not register new backup. %s", err)
	}

	err = removeOldBackups()
	if err != nil {
		return nil, fmt.Errorf("could not remove old backups. %s", err)
	}

	return backup, nil
}

func createArchive(source string, destination string) error {
	destination = fmt.Sprintf("%s.tar.gz", destination)
	destinationFile, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("could not create destination file. %s", err)
	}
	defer destinationFile.Close()

	_, err = os.Stat(source)
	if err != nil {
		return fmt.Errorf("could not create tar archive. %s", err)
	}

	files, err := os.ReadDir(source)
	if err != nil {
		return fmt.Errorf("could not create overview of server files. %s", err)
	}

	fileMap := make(map[string]string)
	for _, file := range files {
		abs, err := filepath.Abs("./testing/" + file.Name())
		if err != nil {
			return fmt.Errorf("failed to find path of file: %s. %s", file.Name(), err)
		}
		if !strings.Contains(file.Name(), "backup") {
			fileMap[abs] = ""
		}
	}

	fileForArchive, err := archiver.FilesFromDisk(nil, fileMap)
	if err != nil {
		return err
	}

	format := archiver.CompressedArchive{
		Compression: archiver.Gz{},
		Archival:    archiver.Tar{},
	}

	err = format.Archive(context.Background(), destinationFile, fileForArchive)
	if err != nil {
		return err
	}

	return nil
}

func removeOldBackups() error {
	backups, err := ReadOverview()
	if err != nil {
		return err
	}

	backupFiles, err := ReadOverview()
	if err != nil {
		return err
	}

	currentTime := time.Now()
	for _, v := range backupFiles {
		backupTimestamp := time.Unix(v.Timestamp, 0)
		diff := currentTime.Sub(backupTimestamp)
		if diff.Hours()/24 > 60 {
			err = os.RemoveAll(v.Destination)
			if err != nil {
				return fmt.Errorf("could not remove backup: %s", v.Destination)
			}

			delete(backups, v.ID.String())
		}
	}

	err = WriteOverview(backups)
	if err != nil {
		return err
	}

	return nil
}
