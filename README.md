# Коляда Даниил
## Практическая работа №15

### Цель работы

Освоить публикацию backend-приложения на удалённом Linux-сервере, научиться подключаться к VPS по SSH, размещать исполняемый файл приложения, настраивать переменные окружения, создавать unit-файл systemd, управлять сервисом через systemctl, анализировать логи через journalctl и выполнять базовую процедуру обновления версии приложения

---

### Шаги

Подключение к VPS по SSH  
С локального компьютера выполнили подключение

```
ssh daniel@192.168.31.210
```

После успешного подключения попали в терминал удалённой машины
![Screenshot](./screenshots/Screenshot_1.png)

---

Обновили пакеты на сервере
```
sudo apt update && sudo apt upgrade -y
```
![Screenshot](./screenshots/Screenshot_2.png)

---

Создали отдельного пользователя для сервиса  
Создали системного пользователя, от имени которого будет запускаться приложение

Такой пользователь нужен только для запуска сервиса
```
sudo useradd --system --no-create-home --shell /usr/sbin/nologin tasksuser
```

Назначение параметров:
- `--system` — системный пользователь
- `--no-create-home` — без домашней директории
- `--shell /usr/sbin/nologin` — запрет интерактивного входа

---

Создали директорию приложения  
Создали каталог для бинарника и служебных файлов приложения

Директория /opt/tasks будет использоваться как место размещения приложения.
```
sudo mkdir -p /opt/tasks
sudo chown -R tasksuser:tasksuser /opt/tasks
```

---

Подготовили конфигурационный файл  
Создали отдельную директорию для конфигурации
```
sudo mkdir -p /etc/tasks
```

Создали env-файл
```
sudo nano /etc/tasks/.env
```
![Screenshot](./screenshots/Screenshot_3.png)

---

После сохранения задали безопасные права

Это означает, что читать и изменять файл сможет только root
```
sudo chown root:root /etc/tasks/.env
sudo chmod 600 /etc/tasks/.env
```

---

Собрали Linux-бинарник на локальной машине  
На локальном компьютере перешли в папку сервиса tasks и выполните сборку под Linux
```
GOOS=linux GOARCH=amd64 go build -o bin/tasks ./cmd/tasks
```

После этого появился исполняемый файл `bin/tasks`

---

Скопировали бинарник на VPS  
С локального компьютера передали файл на сервер
```
scp bin/tasks daniel@192.168.31.210:/tmp/tasks
```

После этого на VPS бинарник будет лежать во временной директории /tmp/tasks

---

Переместили бинарник в рабочую директорию  
На VPS выполнили

Теперь бинарник размещён в целевой директории и готов к запуску
```
sudo mv /tmp/tasks /opt/tasks/tasks
sudo chown tasksuser:tasksuser /opt/tasks/tasks
sudo chmod 755 /opt/tasks/tasks
```

---

Создали unit-файл systemd  
Создайте файл службы
```
sudo nano /etc/systemd/system/tasks.service
```
![Screenshot](./screenshots/Screenshot_4.png)

---

Разобрали назначение параметров unit-файла  
Основные параметры:
- `Description` — краткое описание службы;
- `After=network.target` — запускать сервис после инициализации сети;
- `User=tasksuser` — запуск не от root, а от отдельного пользователя;
- `WorkingDirectory=/opt/tasks` — рабочая директория приложения;
- `EnvironmentFile=/etc/tasks/.env` — подключение внешнего файла конфигурации;
- `ExecStart=/opt/tasks/tasks` — команда запуска приложения;
- `Restart=always` — всегда перезапускать сервис при аварийном завершении;
- `RestartSec=2` — подождать 2 секунды перед повторным запуском;
- `NoNewPrivileges=true` — ограничить получение дополнительных привилегий процессом;
- `LimitNOFILE=65535` — увеличить лимит открытых файлов;
- `WantedBy=multi-user.target` — включить службу в обычный многопользовательский режим запуска системы.

---

Перечитали конфигурацию systemd  
После создания unit-файла необходимо сообщить systemd, что появилась новая служба
```
sudo systemctl daemon-reload
```

---

Запустили сервис  
Выполнили запуск
```
sudo systemctl start tasks
```

---

Включили автозапуск  
Чтобы сервис автоматически стартовал после перезагрузки VPS, выполнили
```
sudo systemctl enable tasks
```

---

Проверили статус сервиса  
Посмотрели текущее состояние
```
sudo systemctl status tasks
```
В статусе видно, что служба активна и запущена
![Screenshot](./screenshots/Screenshot_5.png)

---

Посмотрели логи через journalctl  
Вывели последние записи журнала
```
sudo journalctl -u tasks --no-pager -n 100
```
![Screenshot](./screenshots/Screenshot_6.png)

---

Проверили доступность приложения  
Выполните проверку на локальной машине
![Screenshot](./screenshots/Screenshot_7.png)

---

Выполнили обновление версии приложения  
Минимальная процедура обновления:
1. На локальной машине собрали новый бинарник
2. Скопировали его на VPS
3. Остановили текущий сервис
4. Сохранили старую версию
5. Заменили бинарник
6. Запустили сервис снова

```
sudo systemctl stop tasks
sudo mv /opt/tasks/tasks /opt/tasks/tasks.old
sudo mv /tmp/tasks /opt/tasks/tasks
sudo chown tasksuser:tasksuser /opt/tasks/tasks
sudo chmod 755 /opt/tasks/tasks
sudo systemctl start tasks
```

![Screenshot](./screenshots/Screenshot_8.png)
![Screenshot](./screenshots/Screenshot_9.png)

---

### Выводы

Освоили публикацию backend-приложения на удалённом Linux-сервере, научились подключаться к VPS по SSH, размещать исполняемый файл приложения, настраивать переменные окружения, создавать unit-файл systemd, управлять сервисом через systemctl, анализировать логи через journalctl и выполнять базовую процедуру обновления версии приложения

---

### Дерево проекта

```
├── Dockerfile
├── auth
│   └── auth.go
├── bin
│   └── tasks
├── certs
│   ├── ca.crt
│   ├── server.crt
│   ├── server.key
│   └── server.key.ex
├── cmd
│   └── main.go
├── db
│   └── db.go
├── dtos
│   ├── requests.go
│   └── responses.go
├── go.mod
├── go.sum
├── handlers
│   └── handlers.go
├── logger
│   └── logger.go
├── middleware
│   └── middleware.go
└── utils
    ├── utils.go
    └── utils_test.go

11 directories, 18 files
```