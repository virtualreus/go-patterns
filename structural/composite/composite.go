package composite

import "fmt"

// Позволяет данный структурный паттерн сгруппировать множество обьектов в древовидную структуру и работать с ними как с единым обьектом

// FileSystemComponent Интерфейс компонента
type FileSystemComponent interface {
	Name() string
	Size() int
	IsDir() bool
	Print(indent string)
}

type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{name: name, size: size}
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Size() int {
	return f.size
}

func (f *File) IsDir() bool {
	return false
}

func (f *File) Print(indent string) {
	fmt.Printf("%s📄 %s (%d bytes)\n", indent, f.name, f.size)
}

type Directory struct {
	name     string
	children []FileSystemComponent
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name:     name,
		children: []FileSystemComponent{},
	}
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) Size() int {
	total := 0
	for _, child := range d.children {
		total += child.Size()
	}
	return total
}

func (d *Directory) IsDir() bool {
	return true
}

func (d *Directory) Add(component FileSystemComponent) {
	d.children = append(d.children, component)
}

func (d *Directory) Print(indent string) {
	fmt.Printf("%s📁 %s/ (total: %d bytes)\n", indent, d.name, d.Size())

	for _, child := range d.children {
		child.Print(indent + "  ")
	}
}
