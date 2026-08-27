package asposecellscloud

// WorkbookRef locates a workbook file in cloud storage for the high-level
// functional APIs (export/import/report/editor/search). Name is required;
// Folder and StorageName are optional.
type WorkbookRef struct {
	// Name is the workbook file name in cloud storage (required).
	Name string
	// Folder is the cloud directory that contains the file (optional).
	Folder string
	// StorageName is the cloud storage name (optional).
	StorageName string
}

// NewWorkbookRef creates a WorkbookRef with the given cloud file name.
func NewWorkbookRef(name string) *WorkbookRef {
	return &WorkbookRef{Name: name}
}
