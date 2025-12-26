package flyweight

import (
	"fmt"
	"math/rand"
)

// позволяет вместить бОльшее количество обьектов в отведенную оперативную память. он разделяет общее состояние обьектов между собой.

// Разделяемое состояние (одинаково для всех деревьев одного типа)
type TreeType struct {
	Name     string
	Texture  string
	Color    string
	MeshData []byte // большие данные модели
}

// Уникальное состояние (разное для каждого дерева)
type Tree struct {
	X, Y, Z float64
	Type    *TreeType // ссылка на разделяемый тип
}

// Фабрика типов деревьев
type TreeFactory struct {
	treeTypes map[string]*TreeType
}

func NewTreeFactory() *TreeFactory {
	return &TreeFactory{
		treeTypes: make(map[string]*TreeType),
	}
}

func (f *TreeFactory) GetTreeType(name, texture, color string) *TreeType {
	key := name + "_" + texture + "_" + color

	if treeType, exists := f.treeTypes[key]; exists {
		return treeType
	}

	// Загрузка тяжелых данных (делается один раз)
	meshData := loadMeshData(name)

	treeType := &TreeType{
		Name:     name,
		Texture:  texture,
		Color:    color,
		MeshData: meshData, // тяжелые данные
	}

	f.treeTypes[key] = treeType
	fmt.Printf("Created new tree type: %s\n", key)
	return treeType
}

func loadMeshData(name string) []byte {
	// Имитация загрузки больших данных (3D модель)
	return make([]byte, 1024*1024) // 1 MB на модель
}

// Игровой мир
func main() {
	factory := NewTreeFactory()

	// Создаем лес из 10000 деревьев
	trees := make([]*Tree, 0, 10000)

	// Типы деревьев (создаются только 3 раза)
	oakType := factory.GetTreeType("oak", "oak_texture.png", "green")
	pineType := factory.GetTreeType("pine", "pine_texture.png", "dark_green")
	birchType := factory.GetTreeType("birch", "birch_texture.png", "white")

	// Создаем деревья
	rand.Seed(42)
	for i := 0; i < 10000; i++ {
		var treeType *TreeType

		// Случайно выбираем тип
		switch rand.Intn(3) {
		case 0:
			treeType = oakType
		case 1:
			treeType = pineType
		case 2:
			treeType = birchType
		}

		trees = append(trees, &Tree{
			X:    rand.Float64() * 1000,
			Y:    rand.Float64() * 1000,
			Z:    0,
			Type: treeType, // ссылка на разделяемый тип
		})
	}

	fmt.Printf("Created %d trees\n", len(trees))
	fmt.Printf("Tree types in memory: %d\n", len(factory.treeTypes))
	// Вместо 10000 * 1MB = 10GB памяти, используем 3 * 1MB = 3MB
}
