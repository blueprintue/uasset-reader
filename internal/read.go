package internal

import (
	"errors"
	"fmt"
	"os"
)

type Header struct {
	EPackageFileTag EPackageFileTag

	LegacyVersion    int32
	LegacyVersionUE3 int32

	FileVersionUE4         int32
	FileVersionLicenseeUE4 int32

	CustomVersionContainer uint64

	TotalHeaderSize int32
	A1              int32
	A2              int32
	A3              int32
	A4              int32
	A5              int32
	A6              int32
}

func ReadUasset(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("error while opening file: %v", err)
	}
	defer f.Close()

	// header struct
	header := Header{}

	// check file is uasset
	Read(f, 4, &header.EPackageFileTag)
	if header.EPackageFileTag != PACKAGE_FILE_TAG && header.EPackageFileTag != PACKAGE_FILE_TAG_SWAPPED {
		return errors.New("invalid uasset")
	}

	// check file is not UE3
	Read(f, 4, &header.LegacyVersion)
	if header.LegacyVersion != -6 && header.LegacyVersion != -7 {
		return errors.New("unsupported version")
	}

	// next bytes is for UE3
	Read(f, 4, &header.LegacyVersionUE3)

	// check file is valid UE4
	Read(f, 4, &header.FileVersionUE4)
	Read(f, 4, &header.FileVersionLicenseeUE4)

	if header.FileVersionUE4 == 0 && header.FileVersionLicenseeUE4 == 0 {
		return errors.New("asset unversioned")
	}

	//offset, err := f.Seek(0, os.SEEK_CUR)

	Read(f, 8, &header.CustomVersionContainer)

	/*
		let _legacy_ue3_version: i32 = reader.read_le()?;

		let file_version = reader.read_le()?;
		let file_licensee_version = reader.read_le()?;
		if file_version == 0 && file_licensee_version == 0 {
			return Err(Error::UnversionedAsset);
		}
	*/
	/*Read(f, 4, &header.FileVersionLicenseeUE4)
	Read(f, 4, &header.CustomVersionContainer)
	Read(f, 4, &header.TotalHeaderSize)*/
	Read(f, 4, &header.A1)
	Read(f, 4, &header.A2)
	Read(f, 4, &header.A3)
	Read(f, 4, &header.A4)
	Read(f, 4, &header.A5)
	Read(f, 4, &header.A6)

	fmt.Println("===")
	fmt.Println("EPackageFileTag", header.EPackageFileTag)
	fmt.Printf("Tag %x\n", header.EPackageFileTag)
	fmt.Println("LegacyVersion", header.LegacyVersion)
	fmt.Printf("LegacyVersion %x\n", header.LegacyVersion)
	fmt.Println("LegacyVersionUE3", header.LegacyVersionUE3)
	fmt.Printf("LegacyVersionUE3 %x\n", header.LegacyVersionUE3)
	fmt.Println("FileVersionUE4", header.FileVersionUE4)
	fmt.Printf("FileVersionUE4 %x\n", header.FileVersionUE4)
	fmt.Println("FileVersionLicenseeUE4", header.FileVersionLicenseeUE4)
	fmt.Printf("FileVersionLicenseeUE4 %x\n", header.FileVersionLicenseeUE4)
	fmt.Println("CustomVersionContainer", header.CustomVersionContainer)
	fmt.Printf("CustomVersionContainer %x\n", header.CustomVersionContainer)
	fmt.Println("TotalHeaderSize", header.TotalHeaderSize)
	fmt.Printf("TotalHeaderSize %x\n", header.TotalHeaderSize)
	fmt.Println("A1", header.A1)
	fmt.Printf("A1 %x\n", header.A1)
	fmt.Println("A2", header.A2)
	fmt.Printf("A2 %x\n", header.A2)
	fmt.Println("A3", header.A3)
	fmt.Printf("A3 %x\n", header.A3)
	fmt.Println("A4", header.A4)
	fmt.Printf("A4 %x\n", header.A4)
	fmt.Println("A5", header.A5)
	fmt.Printf("A5 %x\n", header.A5)
	fmt.Println("A6", header.A6)
	fmt.Printf("A6 %x\n", header.A6)
	fmt.Println("===")

	return nil
}
