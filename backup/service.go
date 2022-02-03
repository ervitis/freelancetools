package backup

import (
	"fmt"
	"google.golang.org/api/drive/v3"
	"io"
	"log"
	"os"
	"path/filepath"
)

type (
	IBackup interface {
		DownloadFileIfNotExists(string, string) error
	}

	backup struct {
		driveService *drive.Service
	}
)

func New(drvService *drive.Service) IBackup {
	return &backup{driveService: drvService}
}

func (b *backup) DownloadFileIfNotExists(driveID, name string) error {
	if _, err := os.Stat(fmt.Sprintf("env%s%s", string(filepath.Separator), name)); err == nil {
		log.Println("file exists, no backup download needed")
		return nil
	}

	files, err := b.driveService.Files.
		List().
		IncludeItemsFromAllDrives(true).
		SupportsAllDrives(true).
		Q(fmt.Sprintf(`"%s" in parents and name = "%s" and trashed = false`, driveID, name)).
		Do()
	if err != nil {
		return fmt.Errorf("backup: getting files: %w", err)
	}

	if len(files.Files) == 0 {
		return fmt.Errorf("backup: no files in directory")
	}

	invoiceBackup := files.Files[0]
	resp, err := b.driveService.Files.Get(invoiceBackup.Id).Download()
	if err != nil {
		return fmt.Errorf("backup: getting file backup: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("backup: reading response from file: %w", err)
	}

	f, err := os.OpenFile(fmt.Sprintf("env%s%s", string(filepath.Separator), name), os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	_, _ = f.Write(data)
	return nil
}
