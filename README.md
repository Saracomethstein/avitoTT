# Avito Test Task

Этот проект представляет собой реализацию на основе Go, разработанную для работы в среде Docker. Он следует рекомендациям тестового задания Avito, демонстрируя сервис для управления тендерами и ставками.

## Инструкция по запуску проекта в тестовой среде:

1. Склонируйте репозиторий:
```bash
git clone https://github.com/Saracomethstein/avitoTT.git
```

2. Перейдите в дирикторию проекта и запустите докер образ:
```bash
make docker-up
```

3. После сборки проект доступен по ссылке:
```bash
http://localhost:8080/api/ping
```

4. На данный момент проект работает по ссылке:
```bash
https://cnrprod1725723419-team-78602-32501.avito2024.codenrock.com/api/ping
```

---

## Инструкия по сборке проекта, который залит в деплой:
1. Перейдите в папку с проектом:
```bash
cd deploy/zadanie-6105/
```
ls -l
2. Соберите проект:
```bash
docker build . -t avito-tender-service
```

3. Запустите проект:
```bash
docker run -p 8080:8080 avito-tender-service
```

## Благодарность

Спасибо Avito за возможность и вызов.