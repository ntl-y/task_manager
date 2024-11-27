![image](https://github.com/user-attachments/assets/24dc7cac-68c5-4ea7-92d8-c72e918a736f)# Task Manager Monitoring

Этот проект представляет собой бэкенд-сервис для управления задачами с настройкой мониторинга с использованием **Prometheus** и **Grafana**. Проект включает все необходимые компоненты, такие как база данных **PostgreSQL**, конфигурации мониторинга и графический дашборд.

---

### Запустить Docker Compose
```bash
docker-compose up --build
```
Это поднимет следующие сервисы:
- Бэкенд-сервис на `http://localhost:8888`
- PostgreSQL на `localhost:5555`
- Prometheus на `http://localhost:9090`
- Grafana на `http://localhost:3000` (логин/пароль: `admin/admin`)

---

## 🧪 Тестирование

1. **Проверить бэкенд API**:
   Отправьте HTTP-запрос для создания задачи:
   ```bash
   curl -X POST http://localhost:8888/tasks -d '{"title": "New Titile", "describtion":"New Desc"}' -H "Content-Type: application/json"
   ```

2. **Проверить метрики**:
   Перейдите по адресу: `http://localhost:8889/metrics`. Вы увидите экспортированные метрики.

3. **Prometheus**:
   Откройте `http://localhost:9090`, выполните запрос:
   ```promQL
   task_manager_tasks_created_total
   ```

4. **Grafana**:
   - Войдите в Grafana (`http://localhost:3000`).
   - **Общее количество созданных задач** (`task_manager_tasks_created_total`), сгруппированное по статусу.
   - **Продолжительность создания задач** (`task_manager_task_creation_duration_seconds`), с отображением 95-го перцентиля.

![Grafana Dashboard](https://github.com/user-attachments/assets/3b3e3a1d-c4bc-49fb-9cff-a0b486da89b9)
