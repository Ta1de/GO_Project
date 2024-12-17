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
		expected   string
	}{
		{
			name: "Test -A (after)",
			args: []string{"-A", "2", "test", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "line1\nline2\ntest\nline4\nline5\nline6",
			},
			expected: "test\nline4\nline5\n",
		},
		{
			name: "Test -B (before)",
			args: []string{"-B", "1", "test", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "line1\nline2\ntest\nline4\nline5",
			},
			expected: "line2\ntest\n",
		},
		{
			name: "Test -C (context)",
			args: []string{"-C", "1", "test", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "line1\nline2\ntest\nline4\nline5",
			},
			expected: "line2\ntest\nline4\n",
		},
		{
			name: "Test -c (count)",
			args: []string{"-c", "test", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "line1\nline2\ntest\nline4\ntest\n",
			},
			expected: "2\n",
		},
		{
			name: "Test -i (ignore case)",
			args: []string{"-i", "Test", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "line1\nline2\ntest\nline4\nTEST\n",
			},
			expected: "test\nTEST\n",
		},
		{
			name: "Test -v (invert match)",
			args: []string{"-v", "test", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "line1\nline2\ntest\nline4\n",
			},
			expected: "line1\nline2\nline4\n",
		},
		{
			name: "Test -F (fixed string)",
			args: []string{"-F", "test", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "line1\nline2\ntesting\nline4\ntest\n",
			},
			expected: "test\n",
		},
		{
			name: "Test -n (line number)",
			args: []string{"-n", "test", "file1.txt"},
			inputFiles: map[string]string{
				"file1.txt": "line1\nline2\ntest\nline4\n",
			},
			expected: "3: test\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем временные файлы
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
			//if strings.TrimSpace(got) != strings.TrimSpace(tt.expected) {
			//	t.Errorf("Ожидалось:\n%s\nПолучено:\n%s", tt.expected, got)
			//}

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
