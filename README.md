# Коляда Даниил
## Практическая работа №16

### Цель работы

Освоить базовую публикацию контейнеризированного backend-приложения в Kubernetes, научиться описывать Deployment и Service, передавать конфигурацию через ConfigMap, настраивать readiness и liveness probes, применять манифесты через kubectl и проверять состояние Pod и Service

---

### Шаги

Проверили доступ к кластеру  
Убедились, что kubectl подключён к нужному кластеру
```
kubectl cluster-info
kubectl get nodes
```

Команды отрабатывают успешно и показывают информацию о кластере, доступ настроен корректно
![Screenshot](./screenshots/Screenshot_1.png)

---

Применили ConfigMap
```
kubectl apply -f deploy/k8s/configmap.yaml
```

Применили Deployment
```
kubectl apply -f deploy/k8s/deployment.yaml
```

Применили Service
```
kubectl apply -f deploy/k8s/service.yaml
```

![Screenshot](./screenshots/Screenshot_2.png)

---

Проверили Pod  
Проверили, что Pod создан и находится в рабочем состоянии
```
kubectl get pods
kubectl describe pod tasks
```

![Screenshot](./screenshots/Screenshot_3.png)

---

Проверили Deployment
```
kubectl get deployment
kubectl describe deployment tasks
```

Убедились, что Deployment действительно поддерживает нужное число реплик и не фиксирует ошибок запуска
![Screenshot](./screenshots/Screenshot_4.png)

---

Проверили Service
```
kubectl get svc
kubectl describe svc tasks
```

Убедились, что Service создан и связан с нужными Pod
![Screenshot](./screenshots/Screenshot_5.png)

---

Посмотрели логи контейнера  
Убедились, что приложение стартовало без ошибок
```
kubectl logs tasks
```

![Screenshot](./screenshots/Screenshot_6.png)

---

Проверили доступ через port-forward
```
kubectl port-forward svc/tasks 8082:8082
```

После этого проверили endpoint
![Screenshot](./screenshots/Screenshot_7.png)
![Screenshot](./screenshots/Screenshot_8.png)

---

Выполнили минимальное масштабирование  
Увеличили число экземпляров приложения
```
kubectl scale deployment tasks --replicas=2
kubectl get pods
```

После выполнения команды видно уже два Pod для одного Deployment
![Screenshot](./screenshots/Screenshot_9.png)

---

Вернули одну реплику  
После проверки масштабирования вернули в исходное состояние
```
kubectl scale deployment tasks --replicas=1
```

![Screenshot](./screenshots/Screenshot_10.png)

---

Удалили ресурсы после завершения работы
```
kubectl delete -f deploy/k8s/service.yaml
kubectl delete -f deploy/k8s/deployment.yaml
kubectl delete -f deploy/k8s/configmap.yaml
```

![Screenshot](./screenshots/Screenshot_11.png)

---

### Выводы

Освоили базовую публикацию контейнеризированного backend-приложения в Kubernetes, научились описывать Deployment и Service, передавать конфигурацию через ConfigMap, настраивать readiness и liveness probes, применять манифесты через kubectl и проверять состояние Pod и Service

---

### Дерево проекта

```
├── k8s
│   ├── configmap.yaml
│   ├── deployment.yaml
│   └── service.yaml
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

13 directories, 21 files
```