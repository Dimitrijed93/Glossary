package models

type FolderAndLanguages struct {
	Folders   []Folder   `json:"folders"`
	Languages []Language `json:"languages"`
}

func NewFolderAndLanguages(folders []Folder, languages []Language) *FolderAndLanguages {
	return &FolderAndLanguages{
		Folders:   folders,
		Languages: languages,
	}
}
