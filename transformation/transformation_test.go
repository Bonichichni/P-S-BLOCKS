package transformation

import "testing"

func TestTransformSBlock(t *testing.T) {
	input := [8]int{1, 0, 0, 1, 1, 1, 0, 1}
	expectedOutput := [8]int{0, 1, 0, 1, 1, 1, 1, 0}
	//Я не створював окрему функцію для зворотнього перетворення, тому що алгоритм такий самий тільки потрібна таблиця зворотніх констант
	//Для зворотного перетворення підійде і ця функція, яка проводить стандартне перетворення
	inputInverse := [8]int{0, 1, 0, 1, 1, 1, 1, 0}
	expectedOutputInverse := [8]int{1, 0, 0, 1, 1, 1, 0, 1}
	resultInverse := TransformSBlock(inputInverse, InverseTable(SBox))

	result := TransformSBlock(input, SBox)
	if result != expectedOutput {
		t.Errorf("Функція TransformSBlock повернула неправильний результат. Очікувано %v, отримано %v", expectedOutput, result)
	}

	if resultInverse != expectedOutputInverse {
		t.Errorf("Функція TransformSBlock повернула неправильний результат. Очікувано %v, отримано %v", expectedOutput, result)
	}

}

func TestTransformPBlock(t *testing.T) {
	input := [8]int{1, 1, 1, 0, 0, 0, 0, 1}
	expectedOutput := [8]int{1, 1, 0, 1, 1, 0, 0, 0}
	result := TransformPBlock(input)
	if result != expectedOutput {
		t.Errorf("Функція TransformPBlock повернула неправильний результат. Очікувано %v, отримано %v", expectedOutput, result)
	}
}

func TestInverseTransformPBlock(t *testing.T) {
	input := [8]int{1, 1, 0, 1, 1, 0, 0, 0}
	expectedOutput := [8]int{1, 1, 1, 0, 0, 0, 0, 1}
	result := InverseTransformPBlock(input)
	if result != expectedOutput {
		t.Errorf("Функція InverseTransformPBlock повернула неправильний результат. Очікувано %v, отримано %v", expectedOutput, result)
	}
}
