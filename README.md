# Практическое задание № 12 Синицын Антон Григорьевич ЭФМО-01-25

# Практическое задание: Работа с PostgreSQL

> В этом задании была реализована настройка удалённого сервера PostgreSQL, работа с базой данных `todo`, а также разработаны дополнительные функции работы с задачами.

---

## 🔹 Базовая настройка сервера и БД

### 1. Подключение к удалённому серверу и активация PostgreSQL

<img width="974" height="256" alt="connect-server" src="https://github.com/user-attachments/assets/b48e4a57-f085-462b-a6c8-8d786ae01df6" />

---

### 2. Создание базы данных `todo` и проверка тестовой записи

<img width="961" height="150" alt="create-db" src="https://github.com/user-attachments/assets/c26c1d7f-ebce-44b9-beaf-a2d82c4f1e11" />

---

### 3. Разрешение подключения через порт 5432

<img width="536" height="102" alt="port-5432" src="https://github.com/user-attachments/assets/9a68c04e-8569-4b05-9e90-ee3510c904f7" />

---

### 4. Настройка внешнего доступа к БД

<img width="974" height="79" alt="external-access" src="https://github.com/user-attachments/assets/d3821b5f-f28b-42bb-ab19-20deef1a8fdc" />

---

### 5. Тестовое подключение из localhost

<img width="974" height="252" alt="local-test" src="https://github.com/user-attachments/assets/4426b67b-ccec-4300-b0a6-136ff3c55588" />

---

### 6. Вставка тестовых данных в таблицу `todo`

<img width="827" height="316" alt="insert-test-data" src="https://github.com/user-attachments/assets/e8a254c6-62a5-4873-90e7-c0e815a6637f" />

---

### 7. Проверка данных

<img width="974" height="225" alt="check-data" src="https://github.com/user-attachments/assets/04be76d9-3577-4d0b-90ff-398787cae886" />

---

## 🔹 Дополнительные функции

### 1️⃣ Фильтрация по параметру `Done`

#### Изменение создания записей с полем `Done`

<img width="974" height="414" alt="create-with-done" src="https://github.com/user-attachments/assets/e673f01c-f4b4-4e6f-828d-b728cdfe8e45" />

#### Новый метод `create task` в репозитории

<img width="974" height="238" alt="create-task-repo" src="https://github.com/user-attachments/assets/9742a5c8-8280-44c0-a60a-fbfaf0b765dc" />

#### Функция `ListDone`

<img width="766" height="471" alt="list-done" src="https://github.com/user-attachments/assets/7fd683fc-a0b0-4a1b-9898-63500593fcc0" />

#### Тестирование функции

<img width="641" height="318" alt="test-list-done" src="https://github.com/user-attachments/assets/3b786c40-c0cf-4f01-b3d0-4cca6a06dcd7" />

#### Результат

<img width="850" height="342" alt="result-list-done" src="https://github.com/user-attachments/assets/436cf023-50d3-4579-accf-d108733c8626" />

---

### 2️⃣ Поиск записи по ID

#### Функция поиска

<img width="810" height="306" alt="get-by-id" src="https://github.com/user-attachments/assets/a7c4d7d8-c217-410f-95c9-d7bf011b9153" />

#### Тестирование

<img width="974" height="375" alt="test-get-by-id" src="https://github.com/user-attachments/assets/73015c3f-22cc-4194-8d6f-285efda6af29" />

#### Результат

<img width="445" height="150" alt="result-get-by-id" src="https://github.com/user-attachments/assets/674c6287-7e75-4fa2-b07d-9b53b0734941" />

---

### 3️⃣ Массовая вставка записей

#### Функция

<img width="858" height="695" alt="bulk-insert" src="https://github.com/user-attachments/assets/701b079d-e49e-437a-b6ef-12e3aee33fb9" />

#### Тестирование

<img width="930" height="516" alt="test-bulk-insert" src="https://github.com/user-attachments/assets/f173a9d9-748a-467a-9c2d-73b99da44e72" />

#### Результат

<img width="795" height="388" alt="result-bulk-insert" src="https://github.com/user-attachments/assets/4d2da491-f3e4-449c-ac7c-c4b48810fb21" />

---

### 4️⃣ Настройка подключений

Текущая конфигурация:

<img width="383" height="136" alt="current-config" src="https://github.com/user-attachments/assets/2d8a3b8c-e019-41c4-94a7-4d4115126c6b" />

Новые настройки (рекомендуется 4 одновременных подключения):

<img width="974" height="105" alt="new-config" src="https://github.com/user-attachments/assets/fbdcf7b7-0f4c-4599-b59a-5f1f95a17706" />

Результат:

<img width="426" height="136" alt="result-config" src="https://github.com/user-attachments/assets/a4767511-944a-4772-8843-04cc9eaf9849" />

---

## ❓ Вопросы

1. **Что такое пул соединений `sql.DB` и зачем его настраивать?**  
   Пул соединений — это набор готовых подключений к БД, которые переиспользуются. Настройка пула позволяет избежать открытия нового соединения для каждого запроса и повышает производительность.

2. **Почему используются плейсхолдеры `$1`, `$2`?**  
   Плейсхолдеры обеспечивают:
   - Защиту от SQL-инъекций  
   - Универсальность запросов (изменение столбцов)  
   - Повышение производительности (кэширование плана запроса в БД)

3. **Чем отличаются `Query`, `QueryRow` и `Exec`?**  
   - `Query` — для `SELECT`, возвращает несколько строк  
   - `QueryRow` — для `SELECT`, возвращает одну строку  
   - `Exec` — для `INSERT`, `UPDATE`, `DELETE`, не возвращает данные
