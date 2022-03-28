package dir

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(t *testing.T) string
		isEmpty bool
		wantErr bool
	}{
		{
			name: "directory is empty",
			setup: func(t *testing.T) string {
				return t.TempDir()
			},
			isEmpty: true,
			wantErr: false,
		},
		{
			name: "directory has one directory",
			setup: func(t *testing.T) string {
				path := t.TempDir()
				err := os.Mkdir(filepath.Join(path, "tmp"), 0644)
				if err != nil {
					t.Fail()
				}
				return path
			},
			isEmpty: false,
			wantErr: false,
		},
		{
			name: "directory has two directories",
			setup: func(t *testing.T) string {
				path := t.TempDir()
				for _, subpath := range []string{"tmp", "tmp1"} {
					err := os.Mkdir(filepath.Join(path, subpath), 0644)
					if err != nil {
						t.Fail()
					}
				}
				return path
			},
			isEmpty: false,
			wantErr: false,
		},
		{
			name: "directory has two directories",
			setup: func(t *testing.T) string {
				return "/random"
			},
			isEmpty: false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := tt.setup(t)
			isEmpty, err := IsEmpty(path)
			if isEmpty != tt.isEmpty {
				t.Errorf("IsEmpty() got = %v, want %v", isEmpty, tt.isEmpty)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("IsEmpty() got error = %v, want error %v", err, tt.wantErr)
			}
		})
	}
}
