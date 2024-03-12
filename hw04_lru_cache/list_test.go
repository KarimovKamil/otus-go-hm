package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}

func TestList2(t *testing.T) {
	t.Run("complex2", func(t *testing.T) {
		l := NewList()

		// проверяем правильно ли проставляется первый и последний элемент при вставке в пустой список
		l.PushBack(10) // [10]
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 10, l.Back().Value)
		require.Equal(t, 1, l.Len())

		// проверяем что список корректно очищается при удалении единственного элемента
		l.Remove(l.Front()) // []
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
		require.Equal(t, 0, l.Len())

		// проверяем что вставка нового элемента в только что очищенный список происходит корректно
		l.PushFront(20) // [20]
		require.Equal(t, 20, l.Front().Value)
		require.Equal(t, 20, l.Back().Value)
		require.Equal(t, 1, l.Len())

		// проверяем вставку элементов в начало и конец
		l.PushFront(10) // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 30, l.Back().Value)
		require.Equal(t, 3, l.Len())

		// проверяем, что удаление первого элемента в списке происходит корректно
		l.Remove(l.Front()) // [20, 30]
		require.Equal(t, 20, l.Front().Value)
		require.Equal(t, 2, l.Len())

		// проверяем, что удаление последнего элемента в списке происходит корректно
		l.Remove(l.Back()) // [20]
		require.Equal(t, 20, l.Back().Value)
		require.Equal(t, 1, l.Len())

		// проверяем, что перенос единственного элемента в списке происходит корректно
		l.MoveToFront(l.Back()) // [20]
		require.Equal(t, 20, l.Front().Value)
		require.Equal(t, 20, l.Back().Value)
		require.Equal(t, 1, l.Len())

		// проверяем, что перенос последнего элемента в списке в его начало происходит корректно
		l.PushBack(50)          // [20, 50]
		l.MoveToFront(l.Back()) // [50, 20]
		require.Equal(t, 50, l.Front().Value)
		require.Equal(t, 20, l.Back().Value)

		// проверяем, что перенос первого элемента в списке в его начало происходит корректно
		l.MoveToFront(l.Front()) // [50, 20]
		require.Equal(t, 50, l.Front().Value)
		require.Equal(t, 20, l.Back().Value)
	})
}
