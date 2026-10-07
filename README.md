# Personal Blog

A simple web blog where you can write, publish, edit and delete articles. Built with Go and vanilla HTML/CSS — no frameworks, no database, just files and templates.

🇬 [English](#english) | 🇺 [Русский](#русский)

---

## English

### ✨ Features

**Guest section (public):**
- Home page with a list of all published articles (sorted by ID)
- Article page showing title, publication date and full content

**Admin section (protected):**
- Dashboard with the full list of articles and quick actions (Edit / Delete)
- "New Article" form with title, publication date and content
- "Edit Article" form pre-filled with the current values
- Delete via POST (safe — no accidental GET-deletes)

**Under the hood:**
- HTTP Basic Authentication for the admin section
- Filesystem storage — each article lives in `articles.json`
- Go `html/template` with automatic HTML escaping (XSS-safe)
- Post/Redirect/Get pattern — no duplicate submissions on F5
- Clean architecture: handlers, repository and auth are in separate files

### 🛠 Tech Stack

- **Backend:** Go (`net/http`, `html/template`, `encoding/json`, `crypto/subtle`)
- **Frontend:** vanilla HTML + CSS (no JavaScript)
- **Storage:** single JSON file (`articles.json`)

### 📁 Project Structure

```
personal-blog/
├── main.go            # Server entry point and routes
├── handlers.go        # HTTP handlers (pages, forms, CRUD)
├── repository.go      # Filesystem storage (load / save / find / delete)
── auth.go            # Basic Auth middleware
├── articles.json      # Articles database (created automatically)
├── ui/
│   ├── index.html     # Guest home page
│   ├── templates/     # Go templates
│   │   ├── article.html
│   │   ├── admin.html
│   │   └── form.html  # shared by New and Edit
│   └── static/
│       └── style.css
└── README.md
```

### 🚀 Getting Started

**Requirements:** Go 1.22 or newer (uses `r.PathValue`).

```bash
# Clone and run
git clone <repo-url>
cd personal-blog
go run .
```

The server starts on `http://localhost:8080`.

### 🔐 Credentials

Admin section is protected with HTTP Basic Auth. Default credentials (hardcoded in `auth.go`):

| Field    | Value   |
|----------|---------|
| Username | `admin` |
| Password | `admin` |

⚠️ **For production use, change the credentials before deploying.**

### 📖 Usage

**As a guest:**
1. Open `http://localhost:8080` — you'll see the list of articles.
2. Click on any title to read the full article.

**As an admin:**
1. Open any `/admin` route — the browser will ask for login/password.
2. Dashboard (`/admin`) — see all articles with Edit / Delete buttons.
3. Click **+ Add** or go to `/admin/new` — fill the form and press **Publish**.
4. Click **Edit** next to any article — change fields and press **Update**.
5. Click **Delete** — the article is removed immediately (POST request).

### 🛣 Routes

| Method | Path                      | Access | Description                    |
|--------|---------------------------|--------|--------------------------------|
| GET    | `/`                       | Public | Home page (article list)       |
| GET    | `/article/{id}`           | Public | Single article page            |
| GET    | `/admin`                  | Admin  | Dashboard                      |
| GET    | `/admin/new`              | Admin  | New article form               |
| POST   | `/admin/new/publish`      | Admin  | Create article + redirect      |
| GET    | `/admin/edit/{id}`        | Admin  | Edit article form              |
| POST   | `/admin/edit/{id}/update` | Admin  | Update article + redirect      |
| POST   | `/admin/delete/{id}`      | Admin  | Delete article + redirect      |
| GET    | `/static/*`               | Public | Static files (CSS)             |

### ⚙️ How It Works

- **Storage.** All articles are stored in a single `articles.json` file. Each article is a JSON object with `id`, `title`, `content` and `published_at`. New IDs are generated as `max(existing IDs) + 1`.
- **Templates.** Go's `html/template` renders pages server-side. User-supplied content is auto-escaped, preventing XSS.
- **PRG pattern.** Every mutating POST (`publish`, `update`, `delete`) ends with `http.Redirect(..., StatusSeeOther)`. This prevents duplicate submissions when the user refreshes the page.
- **Auth middleware.** `basicAuth` wraps every `/admin*` handler. It checks credentials using `subtle.ConstantTimeCompare` (timing-safe) and returns `401 Unauthorized` with `WWW-Authenticate` header on failure.
- **Safe deletes.** Delete is a POST form with a button styled as a link. This prevents search engines and link previews from accidentally deleting articles.

### 🛡 Error Handling

- `400 Bad Request` — invalid article ID in URL
- `401 Unauthorized` — missing or wrong admin credentials
- `404 Not Found` — article doesn't exist
- `405 Method Not Allowed` — wrong HTTP method for the endpoint
- `500 Internal Server Error` — template or storage failure

###  Possible Improvements

- Markdown support for article content
- Categories and tags
- Search by title / content
- Pagination on the home page
- User registration instead of hardcoded credentials
- Move storage to SQLite / PostgreSQL
- Rich text editor (e.g. TinyMCE) for the admin form
- Draft / published status

---

## Русский

### ✨ Возможности

**Гостевая часть (публичная):**
- Главная страница со списком всех опубликованных статей (по убыванию ID)
- Страница статьи с заголовком, датой публикации и полным текстом

**Админка (защищённая):**
- Дашборд со списком статей и быстрыми действиями (Edit / Delete)
- Форма «Новая статья» с полями заголовка, даты и содержания
- Форма «Редактировать» с предзаполненными значениями
- Удаление через POST (безопасно — нет случайных удалений по GET)

**Под капотом:**
- HTTP Basic Authentication для админки
- Хранение в файловой системе — все статьи в `articles.json`
- Go `html/template` с автоматическим экранированием HTML (защита от XSS)
- Паттерн Post/Redirect/Get — нет дублирующихся отправок по F5
- Чистая архитектура: handlers, repository и auth в отдельных файлах

### 🛠 Технологии

- **Бэкенд:** Go (`net/http`, `html/template`, `encoding/json`, `crypto/subtle`)
- **Фронтенд:** vanilla HTML + CSS (без JavaScript)
- **Хранилище:** один JSON-файл (`articles.json`)

### 📁 Структура проекта

```
personal-blog/
├── main.go            # Точка входа сервера и маршруты
├── handlers.go        # HTTP-обработчики (страницы, формы, CRUD)
├── repository.go      # Работа с файлами (загрузка / сохранение / поиск / удаление)
├── auth.go            # Middleware Basic Auth
├── articles.json      # База статей (создаётся автоматически)
├── ui/
│   ├── index.html     # Главная страница для гостей
│   ├── templates/     # Go-шаблоны
│   │   ├── article.html
│   │   ├── admin.html
│   │   └── form.html  # общая для New и Edit
│   └── static/
│       └── style.css
── README.md
```

### 🚀 Запуск

**Требования:** Go 1.22 или новее (используется `r.PathValue`).

```bash
# Склонировать и запустить
git clone <repo-url>
cd personal-blog
go run .
```

Сервер стартует на `http://localhost:8080`.

### 🔐 Учётные данные

Админка защищена HTTP Basic Auth. Стандартные учётные данные (захардкожены в `auth.go`):

| Поле     | Значение |
|----------|----------|
| Логин    | `admin`  |
| Пароль   | `admin`  |

⚠️ **Для продакшена обязательно смените учётные данные перед деплоем.**

### 📖 Использование

**Как гость:**
1. Откройте `http://localhost:8080` — увидите список статей.
2. Кликните по заголовку, чтобы прочитать статью целиком.

**Как админ:**
1. Откройте любой `/admin` маршрут — браузер запросит логин и пароль.
2. Дашборд (`/admin`) — все статьи с кнопками Edit / Delete.
3. Нажмите **+ Add** или перейдите на `/admin/new` — заполните форму и нажмите **Publish**.
4. Нажмите **Edit** рядом со статьёй — измените поля и нажмите **Update**.
5. Нажмите **Delete** — статья удалится сразу (POST-запрос).

### 🛣 Маршруты

| Метод  | Путь                      | Доступ  | Описание                          |
|--------|---------------------------|---------|-----------------------------------|
| GET    | `/`                       | Публичный | Главная (список статей)         |
| GET    | `/article/{id}`           | Публичный | Страница одной статьи           |
| GET    | `/admin`                  | Админ   | Дашборд                           |
| GET    | `/admin/new`              | Админ   | Форма новой статьи                |
| POST   | `/admin/new/publish`      | Админ   | Создать статью + редирект         |
| GET    | `/admin/edit/{id}`        | Админ   | Форма редактирования              |
| POST   | `/admin/edit/{id}/update` | Админ   | Обновить статью + редирект        |
| POST   | `/admin/delete/{id}`      | Админ   | Удалить статью + редирект         |
| GET    | `/static/*`               | Публичный | Статика (CSS)                    |

### ⚙️ Как это работает

- **Хранилище.** Все статьи лежат в одном файле `articles.json`. Каждая статья — JSON-объект с полями `id`, `title`, `content` и `published_at`. Новые ID генерируются как `max(существующих ID) + 1`.
- **Шаблоны.** Go `html/template` рендерит страницы на сервере. Пользовательский контент автоматически экранируется — это защита от XSS.
- **PRG-паттерн.** Каждый изменяющий POST (`publish`, `update`, `delete`) заканчивается `http.Redirect(..., StatusSeeOther)`. Это предотвращает дублирование при обновлении страницы.
- **Middleware авторизации.** `basicAuth` оборачивает все `/admin*` обработчики. Проверка учётных данных идёт через `subtle.ConstantTimeCompare` (защита от timing-атак). При ошибке возвращается `401 Unauthorized` с заголовком `WWW-Authenticate`.
- **Безопасное удаление.** Delete — это POST-форма с кнопкой, стилизованной под ссылку. Это защищает от случайного удаления поисковиками и превью ссылок.

### 🛡 Обработка ошибок

- `400 Bad Request` — некорректный ID статьи в URL
- `401 Unauthorized` — неверные или отсутствующие учётные данные админа
- `404 Not Found` — статья не найдена
- `405 Method Not Allowed` — неверный HTTP-метод
- `500 Internal Server Error` — ошибка