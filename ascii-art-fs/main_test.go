package main

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"testing"

	"ascii/asciiart"
)

func TestPrintASCIIArt(t *testing.T) {
	tests := []struct {
		name       string
		bannerFile string
		arguments  string
		want       string
		wantErr    bool
	}{
		{
			name:       "ValidInput",
			bannerFile: "standard.txt",
			arguments:  "\nhello",
			want: `
 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
`,
			wantErr: false,
		},
		{
			name:       "EmptyArgument",
			bannerFile: "standard.txt",
			arguments:  "",
			want:       "",
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		{
			t.Run(tt.name, func(t *testing.T) {
				// Read the banner file content
				fileContent, _ := asciiart.ReadBannerFile(tt.bannerFile)
				lines := asciiart.SplitLines(fileContent, "\n")

				got := captureOutput(func() {
					asciiart.PrintASCIIArt(lines, tt.arguments)
				})

				if tt.wantErr {
					if !reflect.DeepEqual(got, tt.want) {
						t.Errorf("PrintASCIIArt() error = %v, wantErr %v", got, tt.want)
					}
				} else {
					if !reflect.DeepEqual(got, tt.want) {
						t.Errorf("PrintASCIIArt() got = %v, want %v", got, tt.want)
					}
				}
			})
		}
	}
}

func TestGetBannerFileFromArgs(t *testing.T) {
	args := []string{"Program", "arg2", "standard"}
	expected := "standard.txt"
	fileName := asciiart.GetBannerFileFromArgs(args)
	if fileName != expected {
		t.Errorf("Expected filename: %s, but got: %s", expected, fileName)
	}
}

func TestReadBannerFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "ValidFileStandard",
			args:    args{filename: "standard.txt"},
			wantErr: false,
		},
		{
			name:    "ValidFileShadow",
			args:    args{filename: "shadow.txt"},
			wantErr: false,
		},
		{
			name:    "ValidFileThinkertoy",
			args:    args{filename: "thinkertoy.txt"},
			wantErr: false,
		},
		{
			name:    "EmptyFilename",
			args:    args{filename: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := asciiart.ReadBannerFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadBannerFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Only check content for valid files
			if !tt.wantErr {
				expectedContent, err := os.ReadFile(tt.args.filename)
				if err != nil {
					t.Fatalf("failed to read the expected content file: %v", err)
				}
				if got != string(expectedContent) {
					t.Errorf("ReadBannerFile() = %v, want %v", got, string(expectedContent))
				}
			}
		})
	}
}

func TestReadBannerFileHelper(t *testing.T) {
	filename := os.Getenv("TEST_FILENAME")
	if filename == "" {
		return
	}

	_, err := asciiart.ReadBannerFile(filename)
	if err == nil {
		t.Fatalf("expected error for filename %s but got none", filename)
	}
}

func TestEscapeSequence(t *testing.T) {
	args := []string{"program", "argument1"}
	arguments := asciiart.EscapeSequence(args)
	expected := "argument1"
	if arguments != expected {
		fmt.Println("Argument not found")
	}
}

func TestSplitLines(t *testing.T) {
	content := "line1\nline2\nline3\n"
	lines := asciiart.SplitLines(content, "standard.txt")
	expectedLines := []string{"line1", "line2", "line3", ""}
	if !compareslices(lines, expectedLines) {
		fmt.Println("Unable to split lines")
	}
}

// Utility functions for comparing slices and capturing outputs

func compareslices(str1, str2 []string) bool {
	if len(str1) != len(str2) {
		return false
	}
	for i := range str1 {
		if str1[i] != str2[i] {
			return false
		}
	}
	return true
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	_, _ = buf.ReadFrom(r)
	os.Stdout = old

	return buf.String()
}
