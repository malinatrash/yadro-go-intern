package task

import (
	"testing"
)

func TestCanSortBalls(t *testing.T) {

	t.Run("Максимальное количество шаров одного цвета в одном контейнере", func(t *testing.T) {
		containers := [][]int{{1000000000}}
		got := CanSortBalls(containers)
		want := "yes"
		assertCorrectResponse(t, got, want)
	})

	t.Run("Каждый контейнер содержит только один цвет, но количество цветов превышает количество контейнеров", func(t *testing.T) {
		containers := [][]int{{0, 0, 0, 1}, {0, 1}, {0, 0, 1}}
		got := CanSortBalls(containers)
		want := "no"
		assertCorrectResponse(t, got, want)
	})

	t.Run("Базовый тестовый случай, когда сортировка возможна", func(t *testing.T) {
		got := CanSortBalls([][]int{{1, 2}, {2, 1}})
		want := "yes"
		assertCorrectResponse(t, got, want)
	})

	t.Run("Тестовый случай, когда сортировка невозможна", func(t *testing.T) {
		got := CanSortBalls([][]int{{10, 20, 30}, {1, 1, 1}, {0, 0, 1}})
		want := "no"
		assertCorrectResponse(t, got, want)
	})

	t.Run("Минимальное количество контейнеров и шаров", func(t *testing.T) {
		got := CanSortBalls([][]int{{1}})
		want := "yes"
		assertCorrectResponse(t, got, want)
	})

	t.Run("Все шары одного цвета в одном контейнере, остальные пустые", func(t *testing.T) {
		got := CanSortBalls([][]int{{1, 1, 1}, {}, {}})
		want := "no"
		assertCorrectResponse(t, got, want)
	})

	t.Run("Невозможность сортировки из-за неравного количества шаров одного цвета", func(t *testing.T) {
		got := CanSortBalls([][]int{{1, 1, 2}, {2, 2, 3}, {3, 3, 3}})
		want := "no"
		assertCorrectResponse(t, got, want)
	})

	t.Run("Все шары уже отсортированы", func(t *testing.T) {
		got := CanSortBalls([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}})
		want := "yes"
		assertCorrectResponse(t, got, want)
	})

}

func assertCorrectResponse(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
