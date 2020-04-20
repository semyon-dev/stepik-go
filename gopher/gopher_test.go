/*
 * Этот файл содержит тесты для нашего пакета. Это отдельная тема, возможно мы
 * расскажем о ней в дальнейшем, но не во втором модуле. Файл с тестами имеет
 * суффикс _test.go. Если вы потратите немного времени на изучение этого вопроса,
 * то сможете тестировать свои программы заранее, здесь мы рассмотрим лишь несколько
 * примеров.
 */

package gopher

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGopher(t *testing.T) {
	var test interface{} = new(gopherStruct)
	if _, ok := test.(Gopher); !ok {
		t.Error("*gopherStruct не удовлетворяет интерфейсу Gopher")
	}
}

func Test_helloWorld(t *testing.T) {
	buf := bytes.NewBuffer(nil)

	if err := helloWorld(buf, "test message"); err != nil {
		t.Errorf("helloWorld error = %v", err)
	}

	if buf.String() != "test message\n" {
		t.Errorf("buf.String() = %v", buf.String())
	}
}

func ExampleGopher() {
	gopher := NewGopher()
	if err := gopher.Go(); err != nil {
		fmt.Printf("Что-то пошло не так: %v\n", err)
	}

	// Output:
	// I am Gopher version 1.14
	// Hello, World!
}
