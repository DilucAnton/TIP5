package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	// .env не обязателен; если файла нет — ошибка игнорируется
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// fallback — прямой DSN в коде (только для учебного стенда!)
		dsn = os.Getenv("DATABASE_URL")
	}

	db, err := openDB(dsn)
	if err != nil {
		log.Fatalf("openDB error: %v", err)
	}

	// fmt.Println("\n=== Текущие настройки пула ===")
	/* посмотреть изначальные настройки
	stats := db.Stats()
	fmt.Printf("Открытых соединений: %d\n", stats.OpenConnections)
	fmt.Printf("Используется сейчас: %d\n", stats.InUse)
	fmt.Printf("Простаивает: %d\n", stats.Idle)
	fmt.Printf("Максимум открытых: %d\n", stats.MaxOpenConnections)
	*/

	fmt.Println("\n=== Новые настройки пула ===")
	db.SetMaxOpenConns(4)                   // Максимум 3 одновременных подключений
	db.SetMaxIdleConns(2)                   // 2 подключение остается открытыми когда не используются
	db.SetConnMaxLifetime(30 * time.Minute) // Переподключ каждые 30 минут

	stats := db.Stats()
	fmt.Printf("Открытых соединений: %d\n", stats.OpenConnections)
	fmt.Printf("Используется сейчас: %d\n", stats.InUse)
	fmt.Printf("Простаивает: %d\n", stats.Idle)
	fmt.Printf("Максимум открытых: %d\n", stats.MaxOpenConnections)

	defer db.Close()

	repo := NewRepo(db)

	// 1) Вставим пару задач
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tasksToCreate := []struct {
		title string
		done  bool
	}{
		// {"Тест CreateMany 4", true},
		// {"Тест CreateMany 5", true},
		// {"Тест CreateMany 6", true},
	}

	err = repo.CreateMany(ctx, tasksToCreate)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	/*
		// ЗАМЕНЕН НА create many
		for _, task := range tasksToCreate {
			id, err := repo.CreateTask(ctx, task.title, task.done)
			if err != nil {
				log.Fatalf("CreateTask error: %v", err)
			}
			log.Printf("Inserted task id=%d (%s)", id, task.title, task.done)
		}
	*/

	// 2) Прочитаем список задач
	ctxList, cancelList := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelList()

	tasks, err := repo.ListTasks(ctxList)
	if err != nil {
		log.Fatalf("ListTasks error: %v", err)
	}

	// 3) Напечатаем
	fmt.Println("=== Tasks ===")
	for _, t := range tasks {
		fmt.Printf("#%d | %-24s | done=%-5v | %s\n",
			t.ID, t.Title, t.Done, t.CreatedAt.Format(time.RFC3339))
	}

	// 4) Выведем невыполненные задачи через ListDone
	fmt.Println("\n=== Testing ListDone ===")

	// получаем невыполненные задачи
	ctxDone, cancelDone := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelDone()

	undoneTasks, err := repo.ListDone(ctxDone, false)
	if err != nil {
		log.Printf("ListDone (false) error: %v", err)
	} else {
		fmt.Printf("Невыполненные задачи (%d):\n", len(undoneTasks))
		for _, t := range undoneTasks {
			fmt.Printf("  #%d | %s\n", t.ID, t.Title)
		}
	}

	// FindByID
	fmt.Println("\n=== Testing FindID ===")
	ctxFind, cancelFind := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFind()
	var dinfID = 4

	task, err := repo.FindByID(ctxFind, dinfID)
	if err != nil {
		log.Printf("FindByID error for id=%d: %v", dinfID, err)
	} else {
		fmt.Printf("Найдена задача id %d:\n", task.ID)
		fmt.Printf("  Заголовок: %s\n", task.Title)
		fmt.Printf("  Статус: %v\n", task.Done)
		fmt.Printf("  Создана: %s\n", task.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}
