Что нужно сделать чтобы всё заработало:
1. Убедись что запущено:
- PostgreSQL 
- NATS Streaming - выполни в терминале:
go install github.com/nats-io/nats-streaming-server@latest //для установки
nats-streaming-server -p 4223 -m 8223
А если установлен:
cd C:\nats-streaming-server
nats-streaming-server.exe -p 4223 -m 8223
2. Запусти сервис через терминал в папке проекта order_service:
go run cmd/server/main.go 
3. Отправь тестовые данные в папке проекта order_service  (в другом окне терминала):
go run publish.go
4. Открой в браузере:
http://localhost:8080/
5. Найди заказ:
Введи в поиске: b563feb7b2b84b6test
Что должно получиться:
- Сервис подключается к NATS и ждёт данные
- При отправке тестовых данных сохраняет их в БД и кэш
- В веб-интерфейсе можно найти заказ по ID
- Показывает информацию о заказе, доставке, оплате и товарах
# Kim-RWB-L0
