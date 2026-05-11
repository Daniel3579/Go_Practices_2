# Коляда Даниил
## Практическая работа №16

### Цель работы

Освоить базовую публикацию контейнеризированного backend-приложения в Kubernetes, научиться описывать Deployment и Service, передавать конфигурацию через ConfigMap, настраивать readiness и liveness probes, применять манифесты через kubectl и проверять состояние Pod и Service

---

### Шаги
![Screenshot](./screenshots/Screenshot_9.png)

Шаг 1. Проверить доступ к кластеру
Убедитесь, что kubectl подключён к нужному кластеру:
kubectl cluster-info
kubectl get nodes
Если команды отрабатывают успешно и показывают информацию о кластере, доступ настроен корректно.

Шаг 6. Применить ConfigMap
kubectl apply -f deploy/k8s/configmap.yaml
Шаг 7. Применить Deployment
kubectl apply -f deploy/k8s/deployment.yaml
Шаг 8. Применить Service
kubectl apply -f deploy/k8s/service.yaml
Именно такая последовательность применения манифестов указана в исходном материале.
Шаг 9. Проверить Pod
Проверьте, что Pod создан и находится в рабочем состоянии:
kubectl get pods
Если нужно более подробное описание:
kubectl describe pod <pod-name>
Эти команды обязательны для проверки состояния после применения манифестов.
Шаг 10. Проверить Deployment
kubectl get deployment
kubectl describe deployment tasks
Здесь можно убедиться, что Deployment действительно поддерживает нужное число реплик и не фиксирует ошибок запуска.
Шаг 11. Проверить Service
kubectl get svc
kubectl describe svc tasks
Это позволяет убедиться, что Service создан и связан с нужными Pod.
Шаг 12. Посмотреть логи контейнера
Чтобы убедиться, что приложение стартовало без ошибок, выведите логи Pod:
kubectl logs <pod-name>
Если Pod был перезапущен, это тоже полезно видно по логам и по описанию Pod.
Шаг 13. Проверить доступ через port-forward
Для демонстрации работы сервиса извне выполните:
kubectl port-forward svc/tasks 8082:8082
После этого в другом терминале проверьте endpoint:
curl -i http://localhost:8082/health
Это рекомендуемый способ демонстрации доступа к опубликованному сервису в рамках ПЗ 16.
Шаг 14. Проверить реакцию readiness и liveness
После запуска сервиса важно убедиться, что probes не приводят к аварийному перезапуску контейнера и что Pod переходит в состояние готовности.
Проверьте:
kubectl get pods
kubectl describe pod <pod-name>
В описании Pod можно увидеть информацию о probes, перезапусках и событиях.
Шаг 15. Выполнить минимальное масштабирование
В качестве дополнительной демонстрации можно увеличить число экземпляров приложения:
kubectl scale deployment tasks --replicas=2
kubectl get pods
После выполнения команды должно быть видно уже два Pod для одного Deployment. Возможность показать минимальное масштабирование прямо предусмотрена материалом ПЗ 16.
Шаг 16. Вернуть одну реплику
После проверки масштабирования можно вернуть исходное состояние:
kubectl scale deployment tasks --replicas=1

Шаг 17. Удалить ресурсы после завершения работы
Если необходимо освободить стенд после демонстрации:
kubectl delete -f deploy/k8s/service.yaml
kubectl delete -f deploy/k8s/deployment.yaml
kubectl delete -f deploy/k8s/configmap.yaml

---

### Выводы

Освоили базовую публикацию контейнеризированного backend-приложения в Kubernetes, научились описывать Deployment и Service, передавать конфигурацию через ConfigMap, настраивать readiness и liveness probes, применять манифесты через kubectl и проверять состояние Pod и Service

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