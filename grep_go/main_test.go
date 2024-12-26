package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestFlags(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		inputFiles map[string]string // Карта временных файлов: имя -> содержимое
	}{
		{
			name: "Test -A (after)",
			args: []string{"-A", "2", "test", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "line1\nline2\ntest\nline4\nline5\nline6",
			},
		},
		{
			name: "Test -B (before)",
			args: []string{"-B", "1", "test", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "line1\nline2\ntest\nline4\nline5",
			},
		},
		{
			name: "Test -C (context)",
			args: []string{"-C", "1", "test", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "line1\nline2\ntest\nline4\nline5",
			},
		},
		{
			name: "Test -c (count)",
			args: []string{"-c", "test", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "line1\nline2\ntest\nline4\ntest\n",
			},
		},
		{
			name: "Test -i (ignore case)",
			args: []string{"-i", "Test", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "line1\nline2\ntest\nline4\nTEST\n",
			},
		},
		{
			name: "Test -v (invert match)",
			args: []string{"-v", "test", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "line1\nline2\ntest\nline4\n",
			},
		},
		{
			name: "Test -F (fixed string)",
			args: []string{"-F", "a*b", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "a*b\nb\nab\naaab\ntest\na*b\n",
			},
		},
		{
			name: "Test -n (line number)",
			args: []string{"-n", "test", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "line1\nline2\ntest\nline4\n",
			},
		},
		{
			name: "Test no -F (regular)",
			args: []string{"a*b", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "a*b\nb\nab\naaab\ntest\na*b\n",
			},
		},
		{
			name: "Test -A -n (regular)",
			args: []string{"-A", "3", "-n", "test", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "a*b\nb\nab\naaab\ntest\na*b\n",
			},
		},
		{
			name: "Test -B -n (regular)",
			args: []string{"-B", "3", "-n", "test", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "a*b\nb\nab\naaab\ntest\na*b\n",
			},
		},
		{
			name: "Test -C -n (regular)",
			args: []string{"-C", "3", "-n", "test", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "a*b\nb\nab\naaab\ntest\na*b\n",
			},
		},
		{
			name: "Test -c -n (regular)",
			args: []string{"-c", "-n", "a*b", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "a*b\nb\nab\naaab\ntest\na*b\n",
			},
		},
		{
			name: "Test -i -n (regular)",
			args: []string{"-i", "-n", "a*b", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "a*b\nb\nab\naaab\ntest\na*b\n",
			},
		},
		{
			name: "Test -v -n (regular)",
			args: []string{"-v", "-n", "a*b", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "a*b\nb\nab\naaab\ntest\na*b\n",
			},
		},
		{
			name: "Test -F -n (regular)",
			args: []string{"-n", "-F", "a*b", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "a*b\nb\nab\naaab\ntest\na*b\n",
			},
		},
		{
			name: "Test -A -F (regular)",
			args: []string{"-A", "3", "-F", "a*b", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "a*b\nb\nab\naaab\ntest\na*b\na\nb\nc\n",
			},
		},
		{
			name: "Test -B -F (regular)",
			args: []string{"-B", "3", "-F", "a*b", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "a*b\nb\nab\naaab\ntest\na*b\n",
			},
		},
		{
			name: "Test -C -F (regular)",
			args: []string{"-C", "3", "-F", "a*b", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "a*b\nb\nab\naaab\ntest\na*b\n",
			},
		},
		{
			name: "Test -c -F (regular)",
			args: []string{"-c", "-F", "a*b", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "a*b\nb\nab\naaab\ntest\na*b\n",
			},
		},
		{
			name: "Test -i -F (regular)",
			args: []string{"-i", "-n", "a*b", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "a*b\nb\nab\naaab\ntest\na*b\n",
			},
		},
		{
			name: "Test -v -F (regular)",
			args: []string{"-v", "-n", "a*b", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "a*b\nb\nab\naaab\ntest\na*b\n",
			},
		},
		{
			name: "Test -F -n (regular)",
			args: []string{"-n", "-F", "a*b", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "a*b\nb\nab\naaab\ntest\na*b\n",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempFiles := make([]string, 0, len(tt.inputFiles))
			for filename, content := range tt.inputFiles {
				tmpFile, err := os.CreateTemp("", filename)
				if err != nil {
					t.Fatalf("Ошибка создания временного файла: %v", err)
				}
				defer os.Remove(tmpFile.Name()) // Удаляем файл после теста
				if _, err := tmpFile.Write([]byte(content)); err != nil {
					t.Fatalf("Ошибка записи в файл: %v", err)
				}
				tmpFile.Close()
				tempFiles = append(tempFiles, tmpFile.Name())
				for i := range tt.args {
					if tt.args[i] == filename {
						tt.args[i] = tmpFile.Name() // Заменяем имя файла в аргументах
					}
				}
			}

			cmd := exec.Command("./program", tt.args...)
			var out bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &out

			if err := cmd.Run(); err != nil {
				t.Fatalf("Ошибка выполнения программы: %v\nВывод: %s", err, out.String())
			}

			got := out.String()

			grepCmd := exec.Command("grep", tt.args...)
			var grepOut bytes.Buffer
			grepCmd.Stdout = &grepOut
			grepCmd.Stderr = &grepOut

			if err := grepCmd.Run(); err != nil {
				t.Fatalf("Ошибка выполнения grep: %v\nВывод: %s", err, grepOut.String())
			}

			if strings.TrimSpace(got) != strings.TrimSpace(grepOut.String()) {
				t.Errorf("Результат программы не совпадает с grep.\nОжидалось:\n%s\nПолучено:\n%s", grepOut.String(), got)
			}

		})
	}
}
