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
	"github.com/z3orc/ender-cli/internal/global"
	"github.com/z3orc/ender-cli/internal/logger"
)

const BACKUP_DIR = global.WORK_DIR + "/backups"

func init() {
	_, err := os.Stat(BACKUP_DIR)
	if err != nil {
		if err = os.Mkdir(BACKUP_DIR, os.ModePerm); err != nil {
			logger.Error.Fatalln("could not create backup directory. " + err.Error())
		}
	}
}

type Backup struct {
	ID          uuid.UUID
	Destination string
	Timestamp   int64
	Success     bool
}

func New() (*Backup, error) {
	currentTime := time.Now()
	filename := currentTime.Format("2006-01-02_15-04-05")

	backup := &Backup{
		ID:          uuid.New(),
		Destination: BACKUP_DIR + "/" + filename + ".tar.gz",
		Timestamp:   currentTime.Unix(),
	}

	backups, err := ReadOverview()
	if err != nil {
		return nil, fmt.Errorf("could not register new backup. %s", err)
	}

	errArchive := ""
	if err := createArchive(global.WORK_DIR, BACKUP_DIR+"/"+filename); err != nil {
		backup.Success = false
		errArchive = fmt.Sprintf("could not create server backup. %s. ", err)
	} else {
		backup.Success = true
	}

	backups[backup.ID.String()] = backup

	err = WriteOverview(backups)
	if err != nil {
		return nil, fmt.Errorf("%scould not register new backup. %s", errArchive, err)
	}

	err = PurgeOldBackups()
	if err != nil {
		return nil, fmt.Errorf("%scould not remove old backups. %s", errArchive, err)
	}

	if errArchive == "" {
		return backup, nil
	} else {
		return backup, fmt.Errorf(errArchive)
	}

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
		abs, err := filepath.Abs(global.WORK_DIR + "/" + file.Name())
		if err != nil {
			return fmt.Errorf("failed to find path of file: %s. %s", file.Name(), err)
		}

		//Cannot create archive if a file is being written to
		if strings.HasSuffix(abs, "testing/backups") {
			continue
		}

		fileMap[abs] = ""
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

func PurgeOldBackups() error {
	backups, err := ReadOverview()
	if err != nil {
		return err
	}

	currentTime := time.Now()
	for _, v := range backups {
		backupTimestamp := time.Unix(v.Timestamp, 0)
		diff := currentTime.Sub(backupTimestamp)
		if diff.Hours()/24 > 7 || !v.Success {
			err = os.RemoveAll(v.Destination)
			if err != nil {
				return fmt.Errorf("could not remove backup: %s", v.Destination)
			}

			delete(backups, v.ID.String())

			logger.Info.Println("Removed backup: " + v.ID.String())
		}
	}

	err = WriteOverview(backups)
	if err != nil {
		return err
	}

	return nil
}

func PurgeBackups() error {
	backups, err := ReadOverview()
	if err != nil {
		return err
	}

	for _, v := range backups {
		err = os.RemoveAll(v.Destination)
		if err != nil {
			return fmt.Errorf("could not remove backup: %s", v.Destination)
		}

		delete(backups, v.ID.String())

		logger.Info.Println("Removed backup: " + v.ID.String())
	}

	err = WriteOverview(backups)
	if err != nil {
		return err
	}

	return nil
}
